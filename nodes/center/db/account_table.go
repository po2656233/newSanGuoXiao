package db

import (
	"errors"
	"fmt"
	"github.com/po2656233/superplace/const/code"
	exString "github.com/po2656233/superplace/extend/string"
	exTime "github.com/po2656233/superplace/extend/time"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	"superman/internal/guid"
)

// DevAccountTable 开发模式的帐号信息表(platform.TypeDevMode)
type DevAccountTable struct {
	AccountId   int64  `gorm:"column:account_id;primary_key;comment:'帐号id'" json:"accountId"`
	AccountName string `gorm:"column:account_name;comment:'帐号名'" json:"accountName"`
	Password    string `gorm:"column:password;comment:'密码'" json:"-"`
	CreateIP    string `gorm:"column:create_ip;comment:'创建ip'" json:"createIP"`
	CreateTime  int64  `gorm:"column:create_time;comment:'创建时间'" json:"createTime"`
}

func (*DevAccountTable) TableName() string {
	return "dev_account"
}

func DevAccountRegister(uid int64, accountName, password, ip string) int32 {
	devAccount, _ := DevAccountWithName(accountName)
	if devAccount != nil {
		return Register05
	}
	devAccountTable := &DevAccountTable{
		AccountId:   uid, //guid.Next(),
		AccountName: accountName,
		Password:    password,
		CreateIP:    ip,
		CreateTime:  exTime.Now().Unix(),
	}

	devAccountCache.Put(accountName, devAccountTable)
	// TODO 保存db

	return code.OK
}

func DevAccountWithName(accountName string) (*DevAccountTable, error) {
	val, found := devAccountCache.GetIfPresent(accountName)
	if found == false {
		return nil, errors.New("account not found")
	}

	return val.(*DevAccountTable), nil
}

// loadDevAccount 节点启动时预加载帐号数据
func loadDevAccount() {
	// 演示用，直接手工构建几个帐 号
	for i := 1; i <= 10; i++ {
		index := exString.ToString(i)

		devAccount := &DevAccountTable{
			AccountId:   guid.Next(),
			AccountName: "test" + index,
			Password:    "test" + index,
			CreateIP:    "127.0.0.1",
			CreateTime:  exTime.Now().ToMillisecond(),
		}

		devAccountCache.Put(devAccount.AccountName, devAccount)
	}

	log.Info("preload DevAccountTable")
}

// UserBindTable uid绑定第三方平台表
type UserBindTable struct {
	UID      int64  `gorm:"column:uid;primary_key;comment:'用户唯一id'" json:"uid"`
	SdkId    int32  `gorm:"column:sdk_id;comment:'sdk id'" json:"sdkId"`
	PID      int32  `gorm:"column:pid;comment:'平台id'" json:"pid"`
	OpenId   string `gorm:"column:open_id;comment:'平台帐号open_id'" json:"openId"`
	BindTime int64  `gorm:"column:bind_time;comment:'绑定时间'" json:"bindTime"`
	UpTime   int64  `gorm:"column:up_time;comment:'最后一次更新时间'" json:"upTime"`
}

func (*UserBindTable) TableName() string {
	return "user_bind"
}

func GetUID(pid int32, openId string) (int64, bool) {
	cacheKey := fmt.Sprintf(uidKey, pid, openId)
	val, found := uidCache.GetIfPresent(cacheKey)
	if found == false {
		return 0, false
	}

	return val.(int64), true
}

// BindUID 绑定UID
func BindUID(sdkId, pid int32, openId string) (int64, bool) {
	// TODO 根据 platformType的配置要求，决定查询UID的方式：
	// 条件1: platformType + openId查询，是否存在uid
	// 条件2: pid + openId查询，是否存在uid

	uid, ok := GetUID(pid, openId)
	if ok {
		return uid, true
	}

	userBind := &UserBindTable{
		UID:      guid.Next(),
		SdkId:    sdkId,
		PID:      pid,
		OpenId:   openId,
		BindTime: exTime.Now().ToMillisecond(),
	}

	cacheKey := fmt.Sprintf(uidKey, pid, openId)
	uidCache.Put(cacheKey, userBind.UID)

	return userBind.UID, true
}
