package manger

import (
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"time"
)

//// IChair 座椅接口
//type IChair interface {
//	ToSimpleInfo() *protoMsg.PlayerSimpleInfo
//	GetPlayerInfo() *protoMsg.PlayerInfo // 玩家信息
//	GetID() int32                        // ID
//	HandCard() []byte                    // 手牌
//	SetHandCard(card []byte)             // 设置手牌
//	Score() int64                        // 分数
//	SetScore(score int64)                // 设置分数
//	Timer() *time.Timer                  // 定时器
//	SetTimer(t *time.Timer)              // 设置定时器
//	Total() int64                        // 消费总额
//	SetTotal(total int64)                // 设置消费总额
//	Gain() int64                         // 获利
//	SetGain(gain int64)                  // 设置获利
//}

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

// ToSimpleInfo 转成简要信息
func (self *Chair) ToSimpleInfo() *protoMsg.PlayerSimpleInfo {
	return &protoMsg.PlayerSimpleInfo{
		Uid:     self.UserID,
		HeadId:  self.FaceID,
		ChairId: self.ID,
		Score:   self.Score,
		RankNo:  self.Ranking,
		Name:    self.Name,
	}
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

func NewChair(id int32, gid int64, player *Player) any {
	chair := &Chair{
		PlayerInfo: player.PlayerInfo,
		ID:         id,
		HandCard:   make([]byte, 0),
		Score:      INVALID,
		Total:      INVALID,
		Gain:       INVALID,
		Timer:      time.NewTimer(0),
	}
	switch gid {
	case Mahjong:
		return &MahjongChair{
			Chair:    chair,
			Multiple: INVALID,
		}
	default:
		return chair
	}
}
