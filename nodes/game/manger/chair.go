package manger

import (
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"time"
)

// IChair 座椅接口
type IChair interface {
	ToSimpleInfo() *protoMsg.PlayerSimpleInfo
	GetPlayerInfo() *protoMsg.PlayerInfo // 玩家信息
	GetID() int32                        // ID
	HandCard() []byte                    // 手牌
	SetHandCard(card []byte)             // 设置手牌
	Score() int64                        // 分数
	SetScore(score int64)                // 设置分数
	Timer() *time.Timer                  // 定时器
	SetTimer(t *time.Timer)              // 设置定时器
	Total() int64                        // 消费总额
	SetTotal(total int64)                // 设置消费总额
	Gain() int64                         // 获利
	SetGain(gain int64)                  // 设置获利
}

// Chair 椅子
type Chair struct {
	*protoMsg.PlayerInfo
	ID       int32       // 椅子号(仅表示序号) 0:空
	handCard []byte      // 手牌
	score    int64       // 当前一局的消费额(斗地主叫分阶段为叫分分值)
	total    int64       // 消费总额
	gain     int64       // 获利
	timer    *time.Timer // 座位上的定时器
}

// ToSimpleInfo 转成简要信息
func (self *Chair) ToSimpleInfo() *protoMsg.PlayerSimpleInfo {
	return &protoMsg.PlayerSimpleInfo{
		Uid:     self.UserID,
		HeadId:  self.FaceID,
		ChairId: self.GetID(),
		Score:   self.Score(),
		RankNo:  self.Ranking,
		Name:    self.Name,
	}
}

// GetPlayerInfo 玩家信息
func (self *Chair) GetPlayerInfo() *protoMsg.PlayerInfo {
	return self.PlayerInfo
}

// GetID 座椅ID
func (self *Chair) GetID() int32 {
	return self.ID
}

// HandCard 手牌
func (self *Chair) HandCard() []byte {
	return self.handCard
}

// Timer 定时器
func (self *Chair) Timer() *time.Timer {
	return self.timer
}

// Score 分数
func (self *Chair) Score() int64 {
	return self.score
}

// Total 消费总额
func (self *Chair) Total() int64 {
	return self.total
}

// Gain 获利
func (self *Chair) Gain() int64 {
	return self.gain
}

// SetHandCard 设置手牌
func (self *Chair) SetHandCard(card []byte) {
	self.handCard = card
}

// SetScore 设置分数
func (self *Chair) SetScore(sc int64) {
	self.score = sc
}

// SetTimer 设置定时器
func (self *Chair) SetTimer(t *time.Timer) {
	self.timer = t
}

// SetTotal 设置消费总额
func (self *Chair) SetTotal(total int64) {
	self.total = total
}

// SetGain 设置获利
func (self *Chair) SetGain(gain int64) {
	self.gain = gain
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

func NewChair(id int32, gid int64, player *Player) IChair {
	chair := &Chair{
		PlayerInfo: player.PlayerInfo,
		ID:         id,
		handCard:   make([]byte, 0),
		score:      player.Gold,
		total:      0,
		gain:       0,
		timer:      &time.Timer{},
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
