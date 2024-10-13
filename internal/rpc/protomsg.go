package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"google.golang.org/protobuf/proto"
	"os"
	"reflect"
	"strconv"
	. "superman/internal/constant"
	gateMsg "superman/internal/protocol/go_file/gate"
)

// Message 定义结构体
type Message struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Node string `json:"node"`
}

type Messages []Message

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var (
	MapIdMsg = make(map[string]Message)
	Endian   = binary.BigEndian
)

// LoadMsgInfos 获取消息映射文件message_id.json
func LoadMsgInfos() map[string]Message {
	data, _ := os.ReadFile(MSGFile)
	// 解析JSON数据
	err := json.Unmarshal(data, &MapIdMsg)
	if err != nil {
		clog.Errorf("Error:%v", err)
	}
	for s, message := range MapIdMsg {
		id, _ := strconv.ParseInt(s, 10, 64)
		message.ID = uint32(id)
		MapIdMsg[s] = message
	}
	return MapIdMsg
}

// GetProtoData 获取协议结构数据 msgID + data.len + data
func GetProtoData(msg proto.Message) ([]byte, error) {
	strId := ""
	name := reflect.Indirect(reflect.ValueOf(msg)).Type().Name()
	for id, message := range MapIdMsg {
		if message.Name == name {
			strId = id
			break
		}
	}
	if strId == "" {
		return nil, fmt.Errorf("no data")
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, err
	}
	pkg := bytes.NewBuffer([]byte{})
	err = binary.Write(pkg, Endian, uint32(id))
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(msg)
	err = binary.Write(pkg, Endian, uint32(len(data)))
	if err != nil {
		return nil, err
	}
	data = append(pkg.Bytes(), data...)
	return data, nil
}

func ParseProto(msg proto.Message) (uint32, []byte, error) {
	strId := ""
	name := reflect.Indirect(reflect.ValueOf(msg)).Type().Name()
	for id, message := range MapIdMsg {
		if message.Name == name {
			strId = id
			break
		}
	}
	if strId == "" {
		return 0, nil, fmt.Errorf("no data")
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, nil, err
	}
	data, err := proto.Marshal(msg)
	return uint32(id), data, err
}

///////////////////////////simple协议/////////////////////////////////////

// SendResult 反馈结果
func SendResult(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultResp{
		State: int32(code),
		Hints: StatusText[code],
	})
}

func SendHint(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title001],
		Hints: StatusText[code],
	})
}
func SendWarn(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title002],
		Hints: StatusText[code],
	})
}

func SendError(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title003],
		Hints: StatusText[code],
	})
}

// SendSerious 严重错误
func SendSerious(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title004],
		Hints: StatusText[code],
	})
}

func SendFatal(p *simple.ActorBase, code int) {
	p.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title005],
		Hints: StatusText[code],
	})
}

///////////////////////simple agent//////////////////////////////////////////////

// ASendResult 反馈结果
func ASendResult(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultResp{
		State: int32(code),
		Hints: StatusText[code],
	})
}

func ASendHint(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title001],
		Hints: StatusText[code],
	})
}
func ASendWarn(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title002],
		Hints: StatusText[code],
	})
}

func ASendError(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title003],
		Hints: StatusText[code],
	})
}

// ASendSerious 严重错误
func ASendSerious(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title004],
		Hints: StatusText[code],
	})
}

func ASendFatal(agent *simple.Agent, code int) {
	agent.SendMsg(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title005],
		Hints: StatusText[code],
	})
}
