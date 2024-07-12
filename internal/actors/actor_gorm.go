package actors

import (
	"errors"
	superGorm "github.com/po2656233/superplace/components/gorm"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	"gorm.io/gorm"
	. "superman/internal/constant"
	sqlmodel "superman/internal/sql_model"
	. "superman/internal/utils"
	"sync"
	"time"
)

type ActorDB struct {
	cactor.Base
	db    *gorm.DB
	curDB string
	sync.RWMutex
}

func (self *ActorDB) AliasID() string {
	return "db"
}

// OnInit Actor初始化前触发该函数
func (self *ActorDB) OnInit() {
	self.Remote().Register(self.Register)
	self.Remote().Register(self.Login)
	self.Remote().Register(self.ClassList)
	self.changeDB(CenterDb)
	//// 每秒查询一次db
	//p.Timer().Add(5*time.Second, p.selectDB)
	//// 1秒后进行一次分页查询
	//p.Timer().AddOnce(1*time.Second, p.selectPagination)

}
func (self *ActorDB) changeDB(dbNode string) {
	if self.curDB == dbNode {
		return
	}
	self.curDB = dbNode
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
	gormCpt := self.App().Find(name).(*superGorm.Component)
	if gormCpt == nil {
		clog.DPanicf("[component = %s] not found.", name)
		return
	}
	// 获取 db_id = "center_db_1" 的配置
	dbID := self.App().Settings().GetConfig(DbList).GetString(dbNode)
	if self.db != nil {
		dbObj, _ := self.db.DB()
		dbObj.Close()
		self.db = nil
	}
	self.db = gormCpt.GetDb(dbID)
	if self.db == nil {
		clog.Panic(dbID, " not found")
	}
}

/////////////////////////////////////////////////////////////////////////////////

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
func (self *ActorDB) CheckUser(account string) (uid int64) {
	user := sqlmodel.User{}
	err := self.db.Table(user.TableName()).Select("id").Where("account = ?", account).Find(&uid).Error
	if !CheckError(err) {
		return 0
	}
	return uid
}

// CheckRoom 获取玩家ID
func (self *ActorDB) CheckRoom(hostid int64, name string) (count int64, err error) {
	room := sqlmodel.Room{}
	err = self.db.Table(room.TableName()).Where("hostid = ? AND name = ?", hostid, name).Count(&count).Error
	CheckError(err)
	return
}

// GetUserInfo 获取玩家信息
func (self *ActorDB) GetUserInfo(account, password string) (*sqlmodel.User, error) {
	user := &sqlmodel.User{}
	query := "`account`= ? AND `password` = ?"
	err := self.db.Table(user.TableName()).Select("*").Where(query, account, password).Find(user).Error
	CheckError(err)
	return user, err
}

// GetUserID 获取玩家ID
func (self *ActorDB) GetUserID(account, password string) (uid int64, err error) {
	user := &sqlmodel.User{}
	query := "`account`= ? AND `password` = ?"
	err = self.db.Table(user.TableName()).Select("id").Where(query, account, password).Find(&uid).Error
	CheckError(err)
	return
}

// GetClassify 获取游戏种类
func (self *ActorDB) GetClassify() (user []*sqlmodel.Kindinfo, err error) {
	user = make([]*sqlmodel.Kindinfo, 0)
	kind := &sqlmodel.Kindinfo{}
	err = self.db.Table(kind.TableName()).Select("*").Find(&user).Error
	CheckError(err)
	return
}

// GetRooms 获取房间列表
func (self *ActorDB) GetRooms(hostid int64) (rooms []*sqlmodel.Room, err error) {
	rooms = make([]*sqlmodel.Room, 0)
	room := &sqlmodel.Room{}
	err = self.db.Table(room.TableName()).Select("*").Where("hostid=?", hostid).Find(&rooms).Error
	CheckError(err)
	return
}

// AddRoom 新增房间
func (self *ActorDB) AddRoom(user sqlmodel.Room) (int64, error) {
	self.Lock()
	defer self.Unlock()
	if 0 < self.CheckUser(user.Name) {
		return 0, errors.New("房间已经存在")
	}
	user.CreatedAt = time.Now()
	err := self.db.Table(user.TableName()).Create(&user).Error
	if !CheckError(err) {
		return 0, err
	}
	return user.ID, nil
}

/////////////////////////////////////////////////////////////////////////////////////////

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

// BindUser 获取玩家ID
func (self *ActorDB) BindUser(user sqlmodel.UserBind) error {
	err := self.db.Table(user.TableName()).Create(&user).Error
	CheckError(err)
	return err
}
