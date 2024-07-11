package manger

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/sql/redis"
	"sync"
	"time"

	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
)

// Player 玩家属性
type Player struct {
	//基本信息
	*protoMsg.PlayerInfo
	Token string
	//用户包裹
	Knapsack *protoMsg.KnapsackInfo
	// 用于追踪用户行为
	PtrRoom  *Room  // 所在房间 时刻更新最新房间 状态
	PtrTable *Table // 所在桌子 更新桌椅状态
}

// CalculateInfo 结算信息
type CalculateInfo struct {
	UserID   int64  // 结算对象(接受充值的ID)
	ByUID    int64  // 该用户发起
	PreMoney int64  // 结算前的钱
	Payment  int64  // 付款金额
	Code     int32  // 操作码(即为什么付款)
	Order    string // 订单号
	Remark   string // 备注
	/////////////////操作记录////////////////////////
	Gid    int64 // 具体游戏
	HostID int64 // 房主
	Kid    int32 // 类型ID
	Level  int32 // 级别
}

// Rule 规则信息
type Rule struct {
	GameID     int64 //游戏ID
	EnterScore int32 //进 场 分
	LessScore  int32 //坐 下 分
	MaxOnline  int32 //在线人数
	State      int32 //状   态
}

// IUserBehavior 玩家行为
type IUserBehavior interface {
	Enter(args []interface{})     //入场
	Out(args []interface{})       //出场
	Offline(args []interface{})   //离线
	Reconnect(args []interface{}) //重入
	Ready(args []interface{})     //准备
	CallScore(args []interface{}) //叫分
	Host(args []interface{})      //抢庄/地主叫分
	SuperHost(args []interface{}) //超级抢庄
	Look(args []interface{})      //看牌
	Roll(args []interface{})      //摇骰子
	Trustee(args []interface{})   //玩家托管
	Disbanded(args []interface{}) //解散游戏(房主权限)
}

// PlayerManger 管理玩家
type PlayerManger struct {
	sync.Map
} // == persons map[int64]*Player

// ------------------------管理接口--------------------------------------------------//
var manger *PlayerManger = nil
var once sync.Once

// GetPlayerManger 玩家管理对象(单例模式)//manger.persons = make(map[int64]*Player)
func GetPlayerManger() *PlayerManger {
	once.Do(func() {
		manger = &PlayerManger{
			sync.Map{},
		}
	})
	return manger
}

func ToPlayer(info *protoMsg.PlayerInfo) *Player {
	return &Player{
		PlayerInfo: info,
	}
}

func (itself *Player) GetToken() (string, bool) {
	// 生成token 仅为了取用户ID
	token, err := CreateTokenHs256(CustomClaims{
		ID:      itself.UserID,
		Account: itself.Account,
		PlatId:  itself.PlatformID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatInt(itself.UserID, 10),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})
	//cc, err := ParseTokenHs256(token)
	//_ = cc
	if err != nil {
		return token, false
	}

	data, _ := PBToBytes(itself.ToMsg())
	ret, err1 := redis.RedisHandle().Set(GetToken(itself.UserID), data, Date*time.Second).Result()
	if err1 != nil {
		log.Debug("[玩家] ID:%v 账号:%v 获取TOKEN 失败 ERR:%v %v", itself.UserID, itself.Account, err1, ret)
		return "", false
	}
	return token, true
}

func (itself *Player) ToMsg() *protoMsg.PlayerInfo {
	tid := int32(INVALID)
	if nil != itself.PtrTable {
		tid = itself.PtrTable.Num
	}
	info := &protoMsg.PlayerInfo{
		UserID:     itself.UserID,
		Name:       itself.Name,
		FaceID:     itself.FaceID,
		Age:        itself.Age,
		Sex:        itself.Sex,
		Gold:       itself.Gold,
		Money:      itself.Money,
		Level:      itself.Level,
		Account:    itself.Account,
		State:      itself.State,
		PlatformID: itself.PlatformID,
		RoomNum:    itself.RoomNum,
		GameID:     itself.GameID,
		TableID:    itself.TableID,
		ChairID:    itself.ChairID,
	}
	info.TableID = tid
	info.Sex = info.Sex & 0x0F
	return info
}

// Append 添加玩家
func (playerSelf *PlayerManger) Append(play *Player) (*Player, bool) {
	v, ok := playerSelf.Load(play.UserID)
	if !ok || v == nil {
		//log.Debug("[新增玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
		playerSelf.Store(play.UserID, play)
		return play, true
	}
	//log.Debug("[已經存在玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
	return v.(*Player), false
}

// Get 获取指定玩家[根据索引,即userID]
func (playerSelf *PlayerManger) Get(userID int64) *Player {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player)
	}
	return nil
}

// CheckName 获取指定玩家[根据索引,即userID]
func (playerSelf *PlayerManger) CheckName(userID int64) string {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player).Name
	}
	return ""
}

func (playerSelf *PlayerManger) CheckRobot(userID int64) bool {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player).Sex == 0x0F
	}
	return false
}

func (playerSelf *PlayerManger) CheckTable(gid int64) (play *Player, tb *Table, isOk bool) {
	playerSelf.Range(func(key, value interface{}) bool {
		if man, ok := value.(*Player); ok && man.GameID == gid {
			tb = man.PtrTable
			play = man
			isOk = play != nil && tb != nil
		}
		if isOk {
			return false
		}
		return true
	})
	return
}

// 废弃 异步数据
//func (self *PlayerManger) CheckRoomInfo(kid int64, level int32) []int64 {
//    playerList:=make([]int64,0)
//    self.Range(func(key, value interface{}) bool {
//        uid := key.(int64)
//        if man, ok := value.(*Player); ok && man.ChooseKind == kid && man.ChooseLevel == level {
//            playerList = append(playerList,uid)
//        }
//        return true
//    })
//    return playerList
//}

// Exist 玩家是否存在
func (playerSelf *PlayerManger) Exist(userID int64) bool {
	if userID == 0 {
		return false
	}
	isHas := false
	playerSelf.Range(func(key, value interface{}) bool {
		if key.(int64) == userID {
			isHas = true
			return false
		}
		return true
	})
	return isHas
}

// DeleteIndex 按索引删除玩家
func (playerSelf *PlayerManger) DeleteIndex(i int64) {
	playerSelf.Delete(i)
}

// DeletePlayer 删除玩家
func (playerSelf *PlayerManger) DeletePlayer(play *Player) {
	index := int64(0)
	playerSelf.Range(func(key, value interface{}) bool {
		if key.(int64) == play.UserID {
			log.Debug("[删除]\t找到要删除的玩家:%v ", play.UserID)
			index = key.(int64)
			//isOK = true
			value = nil
			return false
		}
		return true
	})
	playerSelf.Delete(index)
	play = nil
}

// /////////////////////////行为接口////////////////////////////////////////////

// Enter 进入场景(限制条件由外部转入)
func (itself *Player) Enter(args []interface{}) { //入场
	game := args[0].(*Game)
	agent := args[1].(gate.Agent)

	userData := agent.UserData()
	if userData == nil {
		GetClientManger().SendResult(agent, FAILED, StatusText[User03])
		return
	}

	//玩家信息
	person := userData.(*Player)
	if person == nil || person.PtrTable == nil {
		GetClientManger().SendResult(agent, FAILED, StatusText[Room07])
		return
	}

	//游戏句柄
	userID := person.UserID
	gameHandle := person.PtrTable.Instance
	if gameHandle == nil {
		log.Debug("[%v:%v] 玩家%v 进入出错:%v", game.G.Name, game.ID, userID, StatusText[Game18])
		GetClientManger().SendResult(agent, FAILED, StatusText[Game18])
		return
	}

	person.GameID = game.ID
	var sceneArgs []interface{}
	sceneArgs = append(sceneArgs, game.G.Level, agent)
	gameHandle.Scene(sceneArgs) // 【进入-> 游戏场景】
}

func (itself *Player) Out(args []interface{}) { //出场
	itself.updateGameInfo(int32(GameUpdateOut), args)
}
func (itself *Player) Offline(args []interface{}) { //离线
	itself.updateGameInfo(int32(GameUpdateOffline), args)
}
func (itself *Player) Reconnect(args []interface{}) { //重入
	itself.updateGameInfo(int32(GameUpdateReconnect), args)
}
func (itself *Player) Ready(args []interface{}) { //准备
	itself.updateGameInfo(int32(GameUpdateReady), args)
}
func (itself *Player) CallScore(args []interface{}) { //叫分
	itself.updateGameInfo(int32(GameUpdateCall), args)
}
func (itself *Player) Host(args []interface{}) { //抢庄
	itself.updateGameInfo(int32(GameUpdateHost), args)
}

func (itself *Player) SuperHost(args []interface{}) { //超级抢庄
	itself.updateGameInfo(int32(GameUpdateHost), args)
}
func (itself *Player) HostList(args []interface{}) { //超级抢庄
	itself.updateGameInfo(int32(GameUpdateHostList), args)
}
func (itself *Player) Follow(args []interface{}) { //跟牌
	itself.updateGameInfo(int32(GameUpdateFollow), args)
}

func (itself *Player) Compare(args []interface{}) { //比牌
	itself.updateGameInfo(int32(GameUpdateCompare), args)
}
func (itself *Player) Look(args []interface{}) { //看牌
	itself.updateGameInfo(int32(GameUpdateLook), args)
}

func (itself *Player) Roll(args []interface{}) { //掷骰子
	itself.updateGameInfo(int32(GameUpdateRollDice), args)
}

func (itself *Player) Trustee(args []interface{}) { //托管
	itself.updateGameInfo(int32(GameUpdateTrustee), args)
}
func (itself *Player) Disbanded(args []interface{}) { //解散
	itself.updateGameInfo(int32(GameUpdateDisbanded), args)
}

func (itself *Player) Exit() { //退出
	itself.State = protoMsg.PlayerState_PlayerStandUp
	itself.GameID = INVALID
	itself.TableID = INVALID
	itself.ChairID = INVALID
	itself.PtrTable = nil //游戏正常结束退出时

}

func (itself *Player) updateGameInfo(flag int32, args []interface{}) bool {
	// 平台维度
	platform := GetPlatformManger().Get(itself.PlatformID)
	// 房间维度
	if nil != platform {
		//查找平台是否包含该房间
		var isHaveRoom = false
		for _, v := range platform.ClassIDs {
			if v == itself.RoomNum {
				isHaveRoom = true
				break
			}
		}
		if isHaveRoom {
			// 游戏维度
			if game := itself.PtrTable; game != nil && nil != game.Instance { //[1-0
				var updateArgs []interface{}
				updateArgs = append(updateArgs, flag, itself.UserID, args)
				//log.Release("\t[%v:%v]玩家:%v-> 操作:%v\n", game.Info.Name, game.GameID, itself.UserID, RecvMessage[int(flag)])
				game.Instance.UpdateInfo(updateArgs)
				return true
			}
		}

	}

	log.Debug("[error:GameUpdate]\t ->:%v userID:%v", flag, itself.UserID)
	return false
}
