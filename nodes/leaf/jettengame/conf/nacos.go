package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/po2656233/goleaf/log"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	RegisterPath = "register"
	RemovePath   = "remove"
	MethodWeb    = "websocket"
	MethodTcp    = "tcp"
	MethodKcp    = "kcp"
)

type Nacos struct {
	Enable       bool   `json:"enable"`
	IP           string `json:"ip"`
	Port         uint64 `json:"port"`
	NamespaceID  string `json:"namespaceId"`
	ServerDataID string `json:"serverDataId"`
	GameDataID   string `json:"gameDataId"`
	Group        string `json:"group"`
	ClusterName  string `json:"clusterName"`
}

var nacos Nacos

func init() {
	// 1、读取nacos信息
	if data, err := os.ReadFile(NacosJsonPath); err != nil {
		log.Fatal("%v", err)
	} else if err = json.Unmarshal(data, &nacos); err != nil {
		log.Fatal("%v", err)
	}
	// 2.1、如果没有配置nacos,则读取本地配置
	if !nacos.Enable {
		InitJson()
		InitYml()
		return
	}

	// 2.2、加载nacos服上的配置信息
	if err := loadNacosConfig(); err != nil {
		os.Exit(1)
	}
}

func getNacosClientParam() (sc []constant.ServerConfig, cc constant.ClientConfig) {
	//create ServerConfig
	sc = []constant.ServerConfig{
		*constant.NewServerConfig(nacos.IP, nacos.Port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc = *constant.NewClientConfig(
		constant.WithNamespaceId(nacos.NamespaceID),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"), //在nacos服关闭后,使用之前缓存的内容
		constant.WithLogLevel("debug"),
	)

	return
}

func loadNacosConfig() error {
	// 获取客户端参数信息
	sc, cc := getNacosClientParam()

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Error("[nacos] NewConfigClient err:%v", err)
		return err
	}

	// get config 获取nacos上的配置信息
	// GameBase
	{
		content, err := client.GetConfig(vo.ConfigParam{
			DataId: nacos.ServerDataID,
			Group:  nacos.Group,
		})
		if err != nil || content == "" {
			log.Error("[nacos] dataId:%v GetConfig err:%v", nacos.ServerDataID, err)
			return err
		}
		//log.Debug("[nacos] dataId:%v GetConfig:%v", dataId, content)
		log.Release("[nacos] dataId:%v success!!", nacos.ServerDataID)

		// parse server config 解析配置信息
		if err = updateJsonConf(content); err != nil {
			log.Error("[nacos] GetConfig updateJsonConf err:%v", err)
		}
	}

	// GameSetting
	{
		content, err := client.GetConfig(vo.ConfigParam{
			DataId: nacos.GameDataID,
			Group:  nacos.Group,
		})
		if err != nil || content == "" {
			log.Error("[nacos] GetConfig err:%v", err)
			return err
		}
		//log.Debug("[nacos] dataId:%v GetConfig:%v", dataId01, content)
		log.Release("[nacos] dataId:%v success!!", nacos.GameDataID)

		// parse server config 解析配置信息
		if err = updateYamlConf(content); err != nil {
			log.Error("[nacos] GetConfig updateYamlConf  err:%v", err)
		}
	}

	// 不监听基础配置,因为多服务之间会引起配置错乱，而各自游戏配置则不会。亦或各服务之间使用唯一的服务名。 Listen config change,key=dataId+group+namespaceId.
	//err = client.ListenConfig(vo.ConfigParam{
	//	DataId: nacosServerDataId,
	//	Group:  nacosGroup,
	//	OnChange: func(namespace, group, dataId, data string) { // 监听配置信息
	//		fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
	//		if err = updateJsonConf(data); err != nil {
	//			log.Error("[nacos] ListenConfig updateJsonConf err:%v", err)
	//		}
	//	},
	//})
	//if err != nil {
	//	return err
	//}
	//// 查找配置
	//searchPage, _ := client.SearchConfig(vo.SearchConfigParam{
	//	Search:   "blur",
	//	DataId:   "",
	//	Group:    "",
	//	PageNo:   1,
	//	PageSize: 10,
	//})
	//fmt.Printf("Search config:%+v \n", searchPage)

	err = client.ListenConfig(vo.ConfigParam{
		DataId: nacos.GameDataID,
		Group:  nacos.Group,
		OnChange: func(namespace, group, dataId, data string) { // 监听配置信息
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
			if err = updateYamlConf(data); err != nil {
				log.Error("[nacos] ListenConfig updateYamlConf  err:%v", err)
			}
		},
	})

	// web服
	if Server.WSAddr != "" {
		if address, ok := registerToNacosServer(ServerName, nacos.ClusterName, nacos.Group); ok {
			Server.WSAddr = address
			log.Release("[nacos] Web服 地址:%v", address)
		}
	}

	// tcp服
	if Server.TCPAddr != "" {
		if address, ok := registerToNacosServer(ServerName, nacos.ClusterName, nacos.Group); ok {
			if address == Server.WSAddr {
				log.Fatal("[nacos] 获取TCP端口失败 与web地址一致:%v", address)
			}
			Server.TCPAddr = address
			log.Release("[nacos] Tcp服 地址:%v", address)
		}
	}
	// kcp服
	if Server.KCPAddr != "" {
		if address, ok := registerToNacosServer(ServerName, nacos.ClusterName, nacos.Group); ok {
			if address == Server.TCPAddr || address == Server.WSAddr {
				log.Fatal("[nacos] 获取KCP端口失败 与web地址一致:%v", address)
			}
			Server.KCPAddr = address
			log.Release("[nacos] Kcp服 地址:%v", address)
		}
	}

	// 更新nacos配置
	data, _ := json.MarshalIndent(Server, "", "\t")
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  nacos.ServerDataID,
		Group:   nacos.Group,
		Content: string(data),
		Type:    vo.JSON,
	})
	return err
}

// NacosCancelListen 取消监听
func nacosCancelListen() {
	if !nacos.Enable {
		return
	}

	// 获取客户端参数信息
	sc, cc := getNacosClientParam()
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Error("[nacos] NewConfigClient err:%v", err)
		return
	}
	//client.CancelListenConfig(vo.ConfigParam{
	//	DataId: nacosServerDataId,
	//	Group:  nacosGroup,
	//})
	_ = client.CancelListenConfig(vo.ConfigParam{
		DataId: nacos.GameDataID,
		Group:  nacos.Group,
	})
}

// registerToNacosServer 注册到nacos
func registerToNacosServer(serverName, clusterName, groupName string) (address string, success bool) {
	// 创建serverConfig
	ip, _ := GetInternetIp() //getNetIP() //
	port := GetFreePort()

	// 获取客户端参数信息
	sc, cc := getNacosClientParam()

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal("[nacos] 初始化nacos失败: %s", err.Error())
		return "", false
	}

	// 注册实例
	data := map[string]string{"name": ServerName, "idc": "shanghai", "timestamp": time.Now().String()}
	success, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serverName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    data,
		ClusterName: clusterName, // 默认值DEFAULT "DEFAULT",       //
		GroupName:   groupName,   // 默认值DEFAULT_GROUP "DEFAULT_GROUP", //
	})
	if err != nil {
		log.Fatal("[nacos] ip:%v port:%v 注册服务失败: %s ", ip, port, err.Error())
	}
	address = ip + fmt.Sprintf(":%v", port)
	log.Release("[nacos] 注册服务地址(%v) success:%v!", address, success)
	return
}

// ToCenterServer 向中心服注册/删除
func ToCenterServer(route, method string, addr *string) error {
	// 移除时 取消监听
	if route == RemovePath {
		nacosCancelListen()
	}

	// 提取IP
	toMethod := method
	if method == MethodKcp {
		toMethod = MethodTcp
	}
	ipData := *addr
	if ipData == "" {
		log.Error("[Center] [%v]-->[%v] no address!", route, method)
		return errors.New(fmt.Sprintf("%v->%v no address!", route, method))
	}

	address := strings.Split(ipData, ":")
	ip := address[0]
	if ip == "" {
		ip, _ = GetInternetIp()
	}
	port := ""
	if 1 < len(address) {
		port = address[1]
	}
	if port == "" {
		port = fmt.Sprintf("%d", GetFreePort())
	}

	// 设置表单信息 即 本服务信息
	data := url.Values{
		"ip":      []string{ip},
		"port":    []string{port},
		"type":    []string{toMethod},
		"name":    []string{ServerName},
		"maxload": []string{fmt.Sprintf("%d", Server.MaxConnNum)}, // 最大负载
	}
	log.Release("[Center] [%v]-->[%v]\tREQ参数:%v", route, method, data)

	// 创建post请求
	req, err := http.NewRequest("POST", Server.CenterWeb+"/"+route, strings.NewReader(data.Encode()))
	if err != nil {
		log.Error("[Center] 创建请求失败:%v", err)
		return err
	}

	// 设置账号
	req.SetBasicAuth(Server.CenterUser, Server.CenterPSW)
	// 设置请求头
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// 创建一个 HTTP 客户端   发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("[Center] 发送请求失败:%v", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应信息
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("[Center] 读取响应失败:%v", err)
		return err
	}
	if !strings.Contains(string(body), "success") {
		return errors.New(string(body))
	}
	*addr = ip + ":" + port
	return nil
}

// GetFreePort 获取一个空闲的端口;端口避免写死,因为要启动多个实例,测试负载均衡
func GetFreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}

// GetInternetIp 获取公网IP
func GetInternetIp() (string, error) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		log.Error("GetInternetIp err:%v", err.Error())
		return "", err
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip, nil
}

// GetNetIP 获取本机网卡IP 不建议使用
func GetNetIP() (ipv4 string, err error) {
	// 获取所有网卡
	address, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr := range address {
		// 这个网络地址是IP地址: ipv4, ipv6
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}
