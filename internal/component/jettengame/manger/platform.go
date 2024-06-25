package manger

import (
	"sync"

	"github.com/po2656233/goleaf/log"
)

// ----------平台----------------------

// PlatformInfo 平台信息
type PlatformInfo struct {
	ID         int64    //平台ID 0表示无效
	Name       string   //平台名称
	ServerList []string //服务器地址
	ClassIDs   []int64  //游戏分类列表
}

// PlatformManger 管理房间
type PlatformManger struct {
	sync.Map
	//rooms map[int32]*RoomInfo
}

// //////////////////平台管理////////////////////////////////////
var platformManger *PlatformManger = nil
var platformOnce sync.Once

// GetPlatformManger 平台管理对象(单例模式)//manger.persons = make(map[int32]*Player)
func GetPlatformManger() *PlatformManger {
	platformOnce.Do(func() {
		platformManger = &PlatformManger{
			sync.Map{},
		}
	})
	return platformManger
}

// Append 添加玩家
func (itself *PlatformManger) Append(plat *PlatformInfo) bool {
	if _, ok := itself.Load(plat.ID); !ok {
		log.Release("新增平台ID:%v NAME:%v 房间号列表:%v", plat.ID, plat.Name, plat.ClassIDs)
		itself.Store(plat.ID, plat)
		return true
	} else {
		log.Release("平台ID:%v 已經存在 NAME:%v", plat.ID, plat.Name, plat.ClassIDs)
		return false
	}
}

// Get 获取指定平台[根据索引,即userID]
func (itself *PlatformManger) Get(platformID int64) *PlatformInfo {
	value, ok := itself.Load(platformID)
	if ok {
		return value.(*PlatformInfo)
	}
	return nil
}

func (itself *PlatformManger) GetSelf(platformName string) *PlatformInfo {
	var plat *PlatformInfo = nil
	itself.Range(func(key, value interface{}) bool {
		plat = value.(*PlatformInfo)
		if plat.Name == platformName {
			return false
		}
		return true
	})
	return plat
}

// Exist 平台是否存在
func (itself *PlatformManger) Exist(platformID int32) bool {
	isHas := false
	itself.Range(func(key, value interface{}) bool {

		if key.(int32) == platformID {
			isHas = true
			return false
		}
		return true
	})
	return isHas
}
