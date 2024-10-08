package manger

import (
	"encoding/binary"
	"github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"net"
	"strconv"
	. "superman/internal/constant"
	gateMsg "superman/internal/protocol/go_file/gate"
	"superman/internal/rpc"
	"superman/internal/utils"
	"sync"
	"time"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}

// ClientManger agent
type ClientManger struct {
	sync.Map
}

var clientManger *ClientManger = nil
var clientOnce sync.Once

// GetClientMgr 玩家管理对象(单例模式)//manger.persons = make(map[int64]*Person)
func GetClientMgr() *ClientManger {
	clientOnce.Do(func() {
		clientManger = &ClientManger{
			Map: sync.Map{},
		}
	})
	return clientManger
}

func (self *ClientManger) SetApp(app facade.IApplication) {
	val, ok := self.Load("app")
	if !ok || val == nil {
		self.Store("app", app)
	}
}
func (self *ClientManger) GetApp() facade.IApplication {
	val, ok := self.Load("app")
	if !ok {
		return nil
	}
	return val.(facade.IApplication)
}

// Append 添加玩家
func (self *ClientManger) Append(userID int64, agent Agent) bool {
	if _, ok := self.Load(userID); ok {
		//log.Debugf("客户端IP:%v 已經存在", v.(Agent).RemoteAddr())
		return false
	}
	log.Debugf("新增一个客户端IP:%v 玩家:%v", agent.RemoteAddr(), userID)
	self.Store(userID, agent)
	return true
}

// Get 客户端连接是否存在
func (self *ClientManger) Get(userID int64) (Agent, bool) {
	if v, ok := self.Load(userID); ok {
		agent := v.(Agent)
		//		log.Debugf("玩家:%v 其客户端IP:%v GET", userID, agent.RemoteAddr())
		return agent, true
	}
	log.Debugf("玩家:%v 待绑定客户端", userID)
	return nil, false
}

// DeleteClient 删除客户端
func (self *ClientManger) DeleteClient(userID int64) {
	self.Delete(userID)
	rpc.SendDataToAcc(self.GetApp(), &gateMsg.LogoutReq{
		Uid: userID,
	})
}

// /------------------------广播---------------------------------------///

// NotifyAll 全网广播
func (self *ClientManger) NotifyAll(msg proto.Message) {
	self.Range(func(key, value interface{}) bool {
		if _, ok := key.(int64); !ok {
			return true
		}

		agent := value.(Agent)
		if nil == agent {
			log.Debugf("无效客户端:%v", value)
			return true
		}

		//校验用户信息
		userData := agent.UserData()
		if nil != userData {
			user := userData.(*Player)
			//if user.Sex == 0x0F { //过滤机器人
			//	return true
			//}
			log.Debugf("通知玩家：%v %v", user.UserID, user.Account)
		} else {
			log.Debugf("err:客户端IP：%v 无效玩家信息", agent.RemoteAddr())
		}

		//广播给客户端
		agent.WriteMsg(msg)
		return true
	})
}

func (self *ClientManger) SendData(agent Agent, msg proto.Message) {
	if userData := agent.UserData(); userData != nil {
		//user := userData.(*Player)
		//if user.Sex == 0x0F { //过滤机器人
		//	return
		//}
		agent.WriteMsg(msg)
	} else {
		agent.WriteMsg(&gateMsg.ResultResp{
			State: FAILED,
			Hints: string("无效请求,请重新登录!"),
		})
	}

}
func (self *ClientManger) SendResult(agent Agent, state int32, hints string) {
	self.SendData(agent, &gateMsg.ResultResp{
		State: state,
		Hints: hints,
	})
}
func (self *ClientManger) SendResultX(uid int64, state int32, hints string) {
	self.SendTo(uid, &gateMsg.ResultResp{
		State: state,
		Hints: hints,
	})
}

// 彈窗提示
func (self *ClientManger) SendPopResult(agent Agent, state int32, title, hints string) {
	self.SendData(agent, &gateMsg.ResultPopResp{
		Flag:  state,
		Title: title,
		Hints: hints,
	})
}

// 彈窗提示
func (self *ClientManger) SendPopResultX(uid int64, state int32, title, hints string) {
	self.SendTo(uid, &gateMsg.ResultPopResp{
		Flag:  state,
		Title: title,
		Hints: hints,
	})
}

func (self *ClientManger) SendError(agent Agent) {
	agent.WriteMsg(&gateMsg.ResultResp{
		State: FAILED,
		Hints: StatusText[Login06],
	})
}
func (self *ClientManger) SendErrorX(uid int64) {
	self.SendTo(uid, &gateMsg.ResultResp{
		State: FAILED,
		Hints: StatusText[Login06],
	})
}

func (self *ClientManger) SendTo(userID int64, msg proto.Message) {
	//defer wait.Done()
	//wait.Add(1)
	self.Range(func(key, value interface{}) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		agent := value.(Agent)
		if nil == agent {
			log.Debugf("无效客户端:%v", value)
			return true
		}

		//获取用户ID
		//userData := agent.UserData()
		if uid == userID {
			//广播给客户端
			agent.WriteMsg(msg)
			//log.Debugf("发送[指定]玩家：%v", userID)
			return false
		}
		return true
	})
}

// 通知除指定玩家外的玩家们
func (self *ClientManger) NotifyButOthers(userIDs []int64, msg proto.Message) {
	//defer wait.Done()
	//wait.Add(1)
	self.Range(func(key, value interface{}) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		agent := value.(Agent)
		if nil == agent {
			log.Debugf("无效客户端:%v", value)
			return true
		}

		//获取用户ID
		for _, item := range userIDs {
			if item == uid {
				return true
			}
		}

		log.Debugf("通知[部分]玩家：%v", uid)
		//广播给客户端
		agent.WriteMsg(msg)
		return true
	})
}

// 发送这批玩家,但除某一个外
func (self *ClientManger) NotifyButOne(userIDs []int64, noNeedToSend int64, msg proto.Message) {
	self.Range(func(key, value interface{}) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		agent := value.(Agent)
		if nil == agent {
			log.Debugf("无效客户端:%v", value)
			return true
		}

		for _, item := range userIDs {
			if item == uid && uid != noNeedToSend {
				//广播给客户端
				agent.WriteMsg(msg)
				break
				//log.Debugf("通知[指定]玩家：%v", uid)
			}
		}
		return true
	})
}

// 发给这部分玩家
func (self *ClientManger) NotifyOthers(userIDs []int64, msg proto.Message) {
	self.Range(func(key, value interface{}) bool {
		uid, ok := key.(int64)
		if !ok {
			return true
		}
		agent := value.(Agent)
		if nil == agent {
			log.Debugf("无效客户端:%v", value)
			return true
		}

		//获取用户ID
		for _, item := range userIDs {
			if uid == item {
				//广播给客户端
				agent.WriteMsg(msg)
				break
				//				l
				// og.Debugf("通知[指定]玩家：%v", uid)
			}
		}
		return true
	})
}

// ////////////////////////////////////////////震惊自己的动态加密/////////////////////////////////////////////////////////////////
// 随机取位,两两字节按位相拼,时间戳做引参,末尾字节255
// 取位2~7 首位随机数，第二为长度
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func disassemble(data, byteTime byte, bit int) (byte, byte) {
	if bit <= 0 || 8 < bit {
		return 0, 0
	}
	return data<<byte(bit) ^ (byteTime >> byte(8-bit)), byteTime<<byte(bit) ^ (data >> byte(8-bit))
}

// DynamicEncode 加密 将一个字节拆成两个或多个
func DynamicEncode(data []byte) []byte {
	tempData := data
	endatas := make([]byte, 0)
	dataLen := len(tempData)
	//fmt.Println("---------------加密-----------加密---------加密-----------------")
	if dataLen <= 0 {
		return endatas
	}
	////当前时间戳
	timestamp := time.Now().Unix()
	strTime := strconv.FormatInt(timestamp, 10)
	btTime := []byte(strTime)
	timeLen := len(btTime)
	//一个字节八位 随机位数，表示前几位和后几位合并 随机数在1~7之间
	//第一位是随机位数 第二位是数据长度 第三位时间戳的起始位置
	ranNum := rand.Int()%6 + 2
	//fmt.Println("加密 随机数:", ranNum)
	endatas = append(endatas, byte(ranNum))

	byteSize := utils.Int16ToBytes(int16(dataLen))
	endatas = append(endatas, byteSize...)
	//fmt.Println("已用长度:", len(endatas), " 数据长度:", dataLen, " 时间戳:", timestamp, " 转换后的字节:", btTime)

	zhihuanSize := timeLen
	if dataLen < timeLen {
		zhihuanSize = dataLen
	}

	//和时间戳进行置换
	for i := 0; i < zhihuanSize; i++ {
		tempData[i], btTime[i] = disassemble(tempData[i], btTime[i], ranNum)
	}

	//超出时间戳长度,则两两置换
	doEncResidue := func(oData []byte, end byte, untilLen, index, oRan int) ([]byte, byte) {
		oLen := len(oData)
		if 2 < untilLen {
			twice := untilLen / 2
			for i := 0; i < twice; i++ {
				oData[index+2*i], oData[index+2*i+1] = disassemble(oData[index+2*i], oData[index+2*i+1], oRan)
			}
		}
		//末尾字节,与255置换
		if 1 == untilLen%2 { //直接与255置换
			oData[oLen-1], end = disassemble(oData[oLen-1], end, oRan)
		}
		return oData, end
	}

	endData := byte(255)
	if timeLen < dataLen {
		tempData, endData = doEncResidue(tempData, endData, dataLen-timeLen, timeLen, ranNum)
	}
	if dataLen < timeLen {
		btTime, endData = doEncResidue(btTime, endData, timeLen-dataLen, dataLen, ranNum)
	}

	endatas = append(endatas, tempData...)
	endatas = append(endatas, btTime...)
	endatas = append(endatas, endData)
	for i := 0; i < len(endatas); i++ {
		endatas[i] ^= byte(i)
	}
	return endatas
}

// DynamicDecode 解密
func DynamicDecode(endata []byte) (data, timestamp []byte) {
	//fmt.Println("--------------------------解密--------------------------")
	data = make([]byte, 0)
	timestamp = make([]byte, 0)

	size := len(endata)
	//时间戳8字节 数据长度2字节 随机数1字节  ---剩下部分 数据字节
	if size < 11 {
		return
	}
	for i := 0; i < len(endata); i++ {
		endata[i] ^= byte(i)
	}

	ranNum := endata[0]

	dataSize := binary.BigEndian.Uint16(endata[1:3])
	dataLen := int(dataSize)
	timeLen := len(strconv.FormatInt(time.Now().Unix(), 10)) //可以固定 为10字节,因为普遍取10字节
	zhihuanSize := timeLen
	if dataLen < zhihuanSize {
		zhihuanSize = dataLen
	}
	data = endata[3 : 3+dataLen]
	//fmt.Println("用于置换的数据", data, len(data))
	timestamp = endata[3+dataLen : 3+dataLen+timeLen]
	//fmt.Println("用于置换的时间", timestamp, len(timestamp))
	for i := 0; i < zhihuanSize; i++ {
		data[i], timestamp[i] = disassemble(timestamp[i], data[i], 8-int(ranNum))
	}

	doDecResidue := func(oData []byte, end byte, untilLen, index, oRan int) ([]byte, byte) {
		oLen := len(oData)
		if 2 < untilLen {
			twice := untilLen / 2
			for i := 0; i < twice; i++ {
				oData[index+2*i], oData[index+2*i+1] = disassemble(oData[index+2*i+1], oData[index+2*i], oRan)
			}
		}

		//末尾字节,与255置换
		if 1 == untilLen%2 { //直接与255置换
			end, oData[oLen-1] = disassemble(oData[oLen-1], end, oRan)
		}
		return oData, end
	}
	if zhihuanSize < dataLen {
		data, _ = doDecResidue(data, endata[size-1], dataLen-zhihuanSize, zhihuanSize, 8-int(ranNum))
	}
	if zhihuanSize < timeLen {
		timestamp, _ = doDecResidue(timestamp, endata[size-1], timeLen-zhihuanSize, zhihuanSize, 8-int(ranNum))
	}

	//fmt.Println(ranNum,"参与数据:", endata[3:3+8], endata[size-8:],dataLen)
	//fmt.Println("数据:", data, "时间戳:", timestamp, "长度:", dataLen, "随机数:", ranNum, "真实置换大小:", zhihuanSize)
	return
}
