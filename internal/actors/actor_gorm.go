package actors

import (
	"errors"
	superGorm "github.com/po2656233/superplace/components/gorm"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	. "sanguoxiao/internal/constant"
	pb "sanguoxiao/internal/protocol/gofile"
	"sanguoxiao/internal/sqlmodel"
	. "sanguoxiao/internal/utils"
	"strconv"
	"sync"
	"time"
)

// UserBindTable uid绑定第三方平台表
type UserBindTable struct {
	UID      int64  `gorm:"column:uid;primary_key;comment:'用户唯一id'" json:"uid"`
	SdkId    int32  `gorm:"column:sdk_id;comment:'sdk配置id'" json:"sdkId"`
	PID      int32  `gorm:"column:pid;comment:'平台id'" json:"pid"`
	OpenId   string `gorm:"column:open_id;comment:'平台帐号open_id'" json:"openId"`
	BindTime int64  `gorm:"column:bind_time;comment:'绑定时间'" json:"bindTime"`
}

func (*UserBindTable) TableName() string {
	return "user_bind"
}

func (p *UserBindTable) PrimaryKey() interface{} {
	return p.UID
}

type ActorDB struct {
	cactor.Base
	db *gorm.DB
	sync.RWMutex
}

func (self *ActorDB) AliasID() string {
	return "db"
}

// OnInit Actor初始化前触发该函数
func (self *ActorDB) OnInit() {
	// db配置的注解
	// 打开profile-dev.json，找到"game-1"和"db"配置
	// 当前示例启动的节点id为 game-1
	// db_id_list参数配置了center_db_1，表示当前节点可以连接该数据库
	// 当前节点启时注册了gorm组件  app.Register(cherryGORM.NewComponent())
	// 通过gorm组件可以获取对应的gorm.DB对象
	// 后续操作请参考gorm的用法
	component := superGorm.NewComponent()
	name := component.Name()
	// 获取gorm组件
	gorm := self.App().Find(name).(*superGorm.Component)
	if gorm == nil {
		clog.DPanicf("[component = %s] not found.", name)
	}

	// 获取 db_id = "center_db_1" 的配置
	centerDbID := self.App().Settings().GetConfig(DbList).GetString(CenterDb)
	self.db = gorm.GetDb(centerDbID)
	if self.db == nil {
		clog.Panic("center_db_1 not found")
	}
	//// 每秒查询一次db
	//p.Timer().Add(5*time.Second, p.selectDB)
	//// 1秒后进行一次分页查询
	//p.Timer().AddOnce(1*time.Second, p.selectPagination)
	self.Remote().Register(self.Register)
}

func (self *ActorDB) selectDB() {
	userBindTable := &UserBindTable{}
	tx := self.db.First(userBindTable)
	if tx.Error != nil {
		clog.Warn(tx.Error)
	}

	clog.Infof("%+v", userBindTable)
}

func (self *ActorDB) selectPagination() {
	list, count := self.pagination(1, 10)
	clog.Infof("count = %d", count)

	for _, table := range list {
		clog.Infof("%+v", table)
	}
}

// pagination 分页查询
func (self *ActorDB) pagination(page, pageSize int) ([]*UserBindTable, int64) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	var list []*UserBindTable
	var count int64

	self.db.Model(&UserBindTable{}).Count(&count)

	if count > 0 {
		list = make([]*UserBindTable, pageSize)
		s := self.db.Limit(pageSize).Offset((page - 1) * pageSize)
		if err := s.Find(&list).Error; err != nil {
			clog.Warn(err)
		}
	}

	return list, count
}

// //////////////////////////////以下有关写操作 强烈要求加锁////////////////////////////////////////////////////////////////////

// AddUser 完成注册 【新增一个玩家】
func (self *ActorDB) AddUser(user sqlmodel.User) (int64, error) {
	self.Lock()
	defer self.Unlock()
	if 0 < self.CheckUser(user.Account) {
		return 0, errors.New("用户已经存在")
	}
	// SignInTime
	now := time.Now()
	user.Signintime = now.Unix()
	user.CreatedAt = now
	err := self.db.Table(user.TableName()).Create(&user).Error
	if !CheckError(err) {
		return 0, err
	}
	return user.ID, nil
}

// CheckUser 获取玩家ID
func (self *ActorDB) CheckUser(name string) (uid int64) {
	user := sqlmodel.User{}
	err := self.db.Table(user.TableName()).Select("id").Where("name = ?", name).Find(&uid).Error
	if !CheckError(err) {
		return 0
	}
	return uid
}

// /////////////////////////////////////////////////////////
func (self *ActorDB) Register(req *pb.RegisterReq) {
	m := req
	identity := uuid.NewV4().String()
	user := sqlmodel.User{
		Name:     m.Name,
		Account:  m.Name,
		Password: Md5Sum(m.Password + strconv.FormatInt(int64(len(m.Password)), 10)),
		Passport: m.PassPortID,
		Realname: m.RealName,
		Phone:    m.PhoneNum,
		Address:  m.Address,
		Email:    m.Email,
		Identity: identity,
		//Clientaddr:   clientAddr,
		Machinecode:  m.MachineCode,
		Referralcode: m.InvitationCode,
		//Serveraddr:   serverAddr,
		Face:   int32(m.FaceID),
		Gender: int32(m.Gender),
		Age:    int32(m.Age),
		//Vip:          LessVIP,
		//Level:        LessLevel,
		//Platformid:   int32(m.PlatformID),
		//Agentid:      agentID,
		//Withdraw:     INVALID,
		//Deposit:      LessMoney,
		//Money:        LessMoney,
	}
	clog.Warnf("ooook req:%v", req)
	self.AddUser(user)
}
