package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	exViper "github.com/po2656233/superplace/extend/viper"
	"sync"
	"time"
)

type RdbClient struct {
	redisOption
	DB     *redis.ClusterClient //本地使用
	Client *redis.Client
}

type redisOption struct {
	// A seed list of host:port addresses of cluster nodes.
	Addrs           []string
	MaxRedirects    int
	ReadOnly        bool
	RouteByLatency  bool
	RouteRandomly   bool
	Username        string
	Password        string
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
	PoolFIFO bool
	// PoolSize applies per cluster node and not for the whole cluster.
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}
type Redis struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Addrs    []string `json:"addrs"`
}

type CnfRedis struct {
	Redis Redis `json:"redis"`
}

var once sync.Once
var redisObj *RdbClient

func SingleRedis() *RdbClient {
	once.Do(func() {
		vp := exViper.NewViper("internal/redis/redis.toml")
		cfr := CnfRedis{}
		vp.Unmarshal(&cfr)
		redisObj = &RdbClient{
			redisOption: redisOption{
				Addrs:    cfr.Redis.Addrs,
				Username: cfr.Redis.Username,
				Password: cfr.Redis.Password,
			},
		}
		initClients()
	})

	return redisObj
}

func initClients() {
	redisObj.DB = redis.NewClusterClient(&redis.ClusterOptions{
		//-------------------------------------------------------------------------------------------
		//集群相关的参数
		//集群节点地址，理论上只要填一个可用的节点客户端就可以自动获取到集群的所有节点信息。但是最好多填一些节点以增加容灾能力，因为只填一个节点的话，如果这个节点出现了异常情况，则Go应用程序在启动过程中无法获取到集群信息。
		Addrs:    redisObj.Addrs,
		Username: redisObj.Username,
		Password: redisObj.Password,
		//MaxRedirects: -1, // 当遇到网络错误或者MOVED/ASK重定向命令时，最多重试几次，默认3
		//只含读操作的命令的"节点选择策略"。默认都是false，即只能在主节点上执行。
		ReadOnly: false, // 置为true则允许在从节点上执行只含读操作的命令
		// 默认false。 置为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中选取Ping()的响应时长最短的一个节点来读数据
		RouteByLatency: false,
		// 默认false。置为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中随机挑选一个节点来读数据
		RouteRandomly: false,

		//用户可定制读取节点信息的函数，比如在非集群模式下可以从zookeeper读取。
		//但如果面向的是redis cluster集群，则客户端自动通过cluster slots命令从集群获取节点信息，不会用到这个函数。
		ClusterSlots: func(_ context.Context) ([]redis.ClusterSlot, error) {
			return []redis.ClusterSlot{
				{
					Start: 0,
					End:   16383,
					Nodes: []redis.ClusterNode{{Addr: ":6379"}},
				},
			}, nil
		},

		//传入的参数是新建的redis.Client
		NewClient: func(opt *redis.Options) *redis.Client {
			return redis.NewClient(opt)
		},

		//------------------------------------------------------------------------------------------------------
		//ClusterClient管理着一组redis.Client,下面的参数和非集群模式下的redis.Options参数一致，但默认值有差别。
		//初始化时，ClusterClient会把下列参数传递给每一个redis.Client

		//钩子函数
		//仅当客户端执行命令需要从连接池获取连接时，如果连接池需要新建连接则会调用此钩子函数
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {
			fmt.Printf("redis conn=%v\n", conn)
			return nil
		},

		//每一个redis.Client的连接池容量及闲置连接数量，而不是cluterClient总体的连接池大小。实际上没有总的连接池
		//而是由各个redis.Client自行去实现和维护各自的连接池。
		PoolSize:     15, // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

	})
	err := redisObj.DB.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
}
func (rdb *RdbClient) Get(ctx context.Context, key string) *redis.StringCmd {
	if rdb.DB != nil {
		rdb.DB.Get(ctx, key)
	}
	return nil
}

func (rdb *RdbClient) Close() {
	if rdb.DB != nil {
		rdb.DB.Close()
		rdb.DB = nil
	}
}
