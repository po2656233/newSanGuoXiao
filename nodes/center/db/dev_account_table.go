package db

import (
	"github.com/po2656233/superplace/const/code"
	sgxTime "github.com/po2656233/superplace/extend/time"
	sgxError "github.com/po2656233/superplace/logger/error"
	"superman/internal/guid"
	"superman/internal/hints"
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

func DevAccountRegister(accountName, password, ip string) int32 {
	devAccount, _ := DevAccountWithName(accountName)
	if devAccount != nil {
		return hints.Register05
	}

	devAccountTable := &DevAccountTable{
		AccountId:   guid.Next(),
		AccountName: accountName,
		Password:    password,
		CreateIP:    ip,
		CreateTime:  sgxTime.Now().Unix(),
	}

	devAccountCache.Put(accountName, devAccountTable)
	// TODO 保存db

	return code.OK
}

func DevAccountWithName(accountName string) (*DevAccountTable, error) {
	val, found := devAccountCache.GetIfPresent(accountName)
	if found == false {
		return nil, sgxError.Error("account not found")
	}

	return val.(*DevAccountTable), nil
}

//// loadDevAccount 节点启动时预加载帐号数据
//func loadDevAccount() {
//	// 演示用，直接手工构建几个帐号
//	for i := 1; i <= 10; i++ {
//		index := sgxString.ToString(i)
//
//		devAccount := &DevAccountTable{
//			AccountId:   guid.Next(),
//			AccountName: "test" + index,
//			Password:    "test" + index,
//			CreateIP:    "127.0.0.1",
//			CreateTime:  sgxTime.Now().ToMillisecond(),
//		}
//
//		devAccountCache.Put(devAccount.AccountName, devAccount)
//	}
//
//	sgxLogger.Info("preload DevAccountTable")
//}
