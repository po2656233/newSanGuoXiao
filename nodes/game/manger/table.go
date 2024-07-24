package manger

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	"sync"
	"sync/atomic"
	"time"
)

// Table 桌牌
type Table struct {
	*protoMsg.TableInfo
	GameHandle IGameOperate
	sitters    []*Chair  // 座位上的玩家
	lookers    []*Player // 旁观者
	MaxChairID int32
	sitCount   int32
	sync.RWMutex
}

// Chair 椅子
type Chair struct {
	*protoMsg.PlayerInfo
	ID       int32       // 椅子号(仅表示序号) 0:空
	HandCard []byte      // 手牌
	Score    int64       // 当前一局的消费额(斗地主叫分阶段为叫分分值)
	Total    int64       // 消费总额
	Gain     int64       // 获利
	Timer    *time.Timer // 座位上的定时器
}

// ZjhChair zjh桌位
type ZjhChair struct {
	*Chair
	Multiple  int     //倍数|番数
	Challenge []int64 //挑战过的玩家
}

// MahjongChair 麻将桌位
type MahjongChair struct {
	*Chair
	Multiple int //倍数|番数
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
	chair := &Chair{
		PlayerInfo: player.PlayerInfo,
		ID:         atomic.AddInt32(&tb.MaxChairID, 1),
		HandCard:   make([]byte, 0),
		Score:      0,
		Total:      0,
		Gain:       0,
		Timer:      &time.Timer{},
	}
	tb.sitters = append(tb.sitters, chair)
	tb.sitters = utils.SliceRemoveDuplicate(tb.sitters).([]*Chair)
	atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
	for _, sitter := range tb.sitters {
		resp.PlayerList = append(resp.PlayerList, &protoMsg.PlayerSimpleInfo{
			Uid:     sitter.UserID,
			HeadId:  sitter.FaceID,
			ChairId: sitter.InChairId,
			Score:   sitter.Score,
			RankNo:  sitter.Ranking,
			Name:    sitter.Name,
		})
		playerList = append(playerList, sitter.UserID)
	}
	// 发送给刚坐下的玩家，关于其他玩家的信息
	GetClientMgr().SendTo(player.UserID, resp)

	resp.PlayerList = []*protoMsg.PlayerSimpleInfo{{
		Uid:     chair.UserID,
		HeadId:  chair.FaceID,
		ChairId: chair.InChairId,
		Score:   chair.Score,
		RankNo:  chair.Ranking,
		Name:    chair.Name,
	}}
	// 发给其他玩家，当前刚坐下的玩家
	GetClientMgr().NotifyButOne(playerList, player.UserID, resp)
	// 游戏准备 添加刚坐下的玩家
	tb.GameHandle.Ready([]interface{}{player})
	// 满员后,开始游戏
	if (tb.MaxSitter == Unlimited && tb.sitCount == Default) || tb.MaxSitter == tb.sitCount {
		tb.GameHandle.Start(nil)
	}

	return nil
}

// GetChair 获取牌桌信息
func (tb *Table) GetChair(uid int64) *Chair {
	tb.RLock()
	defer tb.RUnlock()
	for _, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == uid {
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
		list = append(list, sitter.PlayerInfo)
	}
	return list
}

func (tb *Table) RemoveChair(player *Player) {
	tb.Lock()
	defer tb.Unlock()
	player.PlayerInfo.State = protoMsg.PlayerState_PlayerStandUp
	for i, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == player.UserID {
			tb.sitters = append(tb.sitters[:i], tb.sitters[i+1:]...)
			break
		}
	}
	tb.sitters = utils.SliceRemoveDuplicate(tb.sitters).([]*Chair)
	atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
}

// NextChair 轮换规则
func (tb *Table) NextChair(curUid int64) *Chair {
	tb.RLock()
	defer tb.RUnlock()
	size := len(tb.sitters)
	if 0 == size {
		return nil
	}
	index := 0
	for i, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == curUid {
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
		if tb.sitters[idx].State < protoMsg.PlayerState_PlayerTrustee && tb.sitters[idx].State > protoMsg.PlayerState_PlayerAgree {
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
	for _, sit := range tb.sitters {
		sit.State = protoMsg.PlayerState_PlayerSitDown
		sit.Gain = INVALID
		sit.Score = INVALID
		sit.Total = INVALID
		sit.HandCard = make([]byte, 0)
		if sit.Timer != nil {
			sit.Timer.Stop()
			sit.Timer = nil
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
