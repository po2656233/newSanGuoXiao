package manger

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	"sync"
	"sync/atomic"
)

// Table 桌牌
type Table struct {
	*protoMsg.TableInfo
	GameHandle IGameOperate
	sitters    []IChair  // 座位上的玩家
	lookers    []*Player // 旁观者
	maxChairID int32     // 便于新增椅子号
	sitCount   int32
	sync.RWMutex
}

// ///////////////////桌牌功能//////////////////////////////////////////////
//var settingFixTime int64

// Init 加载配置 [废弃:不加载公用配置,拆分至子游戏内部实现]
func (tb *Table) Init() {
	//解析游戏
	////获取文件的
	//fInfo, _ := os.Stat(settingFile)
	//lastFixTime := fInfo.ModTime().Unix()
	//// 在文件被修改或配置内容为空时，则读取配置信息
	//if lastFixTime != settingFixTime  {
	//	//读取配置信息
	//	yamlFile, err := os.ReadFile(settingFile)
	//	//log.Debug("yamlFile:%v", yamlFile)
	//	if err != nil {
	//		log.Fatal("yamlFile.Get err #%v ", err)
	//	}
	//
	//	err = yaml.Unmarshal(yamlFile, self.Config)
	//	// err = yaml.Unmarshal(yamlFile, &resultMap)
	//	if err != nil {
	//		log.Error("Unmarshal: %v", err)
	//	}
	//}
	//settingFixTime = lastFixTime

	//log.Debug("conf:%v", self.Config)
}

// AddChair 添加椅子
func (tb *Table) AddChair(player *Player) error {
	tb.Lock()
	defer tb.Unlock()
	// 是否满员了
	if tb.MaxSitter != Unlimited && tb.MaxSitter < int32(len(tb.sitters)+1) {
		return fmt.Errorf("已经满员")
	}

	playerList := make([]int64, 0)
	player.PlayerInfo.State = protoMsg.PlayerState_PlayerSitDown
	resp := &protoMsg.JoinGameReadyQueueResp{
		RoomID:     tb.Rid,
		TableID:    tb.Id,
		PlayerList: make([]*protoMsg.PlayerSimpleInfo, 0),
	}
	// 检测是否存在了
	var chair IChair
	for _, sitter := range tb.sitters {
		info := sitter.GetPlayerInfo()
		if info.UserID == player.UserID {
			chair = sitter
			break
		}
	}
	isNew := false
	if chair == nil {
		isNew = true
		chair = NewChair(atomic.AddInt32(&tb.maxChairID, 1), tb.Gid, player)
		tb.sitters = append(tb.sitters, chair)
		atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
	}

	// 获取玩家简要信息
	for _, sitter := range tb.sitters {
		info := sitter.GetPlayerInfo()
		resp.PlayerList = append(resp.PlayerList, sitter.ToSimpleInfo())
		playerList = append(playerList, info.UserID)
	}
	// 发送给刚坐下的玩家，关于其他玩家的信息
	player.InRooId = tb.Rid
	player.InTableId = tb.Id
	player.InChairId = chair.GetID()
	player.GameHandle = tb.GameHandle
	GetClientMgr().SendTo(player.UserID, resp)

	// 发给其他玩家，当前刚坐下的玩家
	resp.PlayerList = []*protoMsg.PlayerSimpleInfo{chair.ToSimpleInfo()}
	GetClientMgr().NotifyButOne(playerList, player.UserID, resp)

	// 游戏准备 添加刚坐下的玩家
	tb.GameHandle.UpdateInfo([]interface{}{player.State, player.UserID})
	// 满员后,开始游戏
	if isNew && ((tb.MaxSitter == Unlimited && tb.sitCount == Default) || tb.MaxSitter == tb.sitCount) {
		tb.GameHandle.Start(nil)
	}
	return nil
}

// GetChair 获取牌桌信息
func (tb *Table) GetChair(uid int64) IChair {
	tb.RLock()
	defer tb.RUnlock()
	for _, sitter := range tb.sitters {
		if sitter.GetPlayerInfo().UserID == uid {
			return sitter
		}
	}
	return nil
}

func (tb *Table) GetChairInfos() []*protoMsg.PlayerInfo {
	tb.RLock()
	defer tb.RUnlock()
	list := make([]*protoMsg.PlayerInfo, 0)
	for _, sitter := range tb.sitters {
		list = append(list, sitter.GetPlayerInfo())
	}
	return list
}

func (tb *Table) RemoveChair(uid int64) {
	tb.Lock()
	defer tb.Unlock()
	for _, sitter := range tb.sitters {
		if sitter.GetPlayerInfo().UserID == uid {
			tb.sitters = utils.RemoveValue(tb.sitters, sitter)
			atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
			return
		}
	}
}

// NextChair 轮换规则
func (tb *Table) NextChair(curUid int64) IChair {
	tb.RLock()
	defer tb.RUnlock()
	size := len(tb.sitters)
	if 0 == size {
		return nil
	}
	index := 0
	for i, sitter := range tb.sitters {
		if sitter.GetPlayerInfo().UserID == curUid {
			index = i
			break
		}
	}
	index++
	if index <= size {
		index = 0
	}
	for i := 0; i < size; i++ {
		idx := index + i
		if idx <= size {
			idx -= size
		}
		playInfo := tb.sitters[idx].GetPlayerInfo()
		if playInfo.State < protoMsg.PlayerState_PlayerTrustee && playInfo.State > protoMsg.PlayerState_PlayerAgree {
			return tb.sitters[idx]
		}
	}
	return nil
}
func (tb *Table) SitCount() int32 {
	return atomic.LoadInt32(&tb.sitCount)
}

// Reset 重置座位上的玩家
func (tb *Table) Reset() int64 {
	log.Debugf("[%v:%v]   \t扫地僧出来干活了...List %v", tb.Name, tb.Id, tb.sitters)
	// 遍历过程中存在slice的数据删除
	tb.Lock()
	for _, sitter := range tb.sitters {
		sit, ok := sitter.(*Chair)
		if ok {
			continue
		}
		sit.State = protoMsg.PlayerState_PlayerSitDown
		sit.SetGain(INVALID)
		sit.SetScore(INVALID)
		sit.SetTotal(INVALID)
		t := sit.Timer()
		if t != nil {
			t.Stop()
			sit.SetTimer(nil)
		}
	}
	tb.Unlock()

	tb.CorrectCommission()
	tb.Init()
	return tb.Id
}

// CorrectCommission 修正税收
func (tb *Table) CorrectCommission() {

}
