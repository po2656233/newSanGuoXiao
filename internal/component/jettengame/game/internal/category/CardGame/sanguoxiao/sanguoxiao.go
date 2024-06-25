package sanguoxiao

import (
	"sanguoxiao/internal/component/jettengame/conf"
	. "sanguoxiao/internal/component/jettengame/manger"
	_ "sanguoxiao/internal/component/jettengame/sql/mysql" //仅仅希望导入 包内的init函数
	"sync"
)

// BrTuitongziGame 继承于GameItem
type SanguoxiaoGame struct {
	*Game

	personBetInfo *sync.Map //map[int64][]*protoMsg.BrTuitongziBetResp //下注信息

}

// NewGame 创建实例
func NewGame(game *Game) *SanguoxiaoGame {
	p := &SanguoxiaoGame{
		Game: game,
	}

	p.Config = conf.YamlObj
	p.Init()
	return p
}

// Init 初始化信息
func (self *SanguoxiaoGame) Init() {

}

// Scene 场 景
func (self *SanguoxiaoGame) Scene(args []interface{}) {

}

// Start 开 始
func (self *SanguoxiaoGame) Start(args []interface{}) {

}

// Playing 游 戏
func (self *SanguoxiaoGame) Playing(args []interface{}) {

}

// Over 结 算
func (self *SanguoxiaoGame) Over(args []interface{}) {

}

// UpdateInfo 更新信息
func (self *SanguoxiaoGame) UpdateInfo(args []interface{}) {

}

// SuperControl 超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭
func (self *SanguoxiaoGame) SuperControl(args []interface{}) {

}
