package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	clog "github.com/po2656233/superplace/logger"
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

// GetProtoResult 获取协议结构数据 msgID + data.len + data
func GetProtoResult(code int) ([]byte, error) {
	return GetProtoData(&gateMsg.ResultResp{
		State: int32(code),
		Hints: StatusText[code],
	})
}

// GetProtoResultPop 获取协议结构数据 msgID + data.len + data
func GetProtoResultPop(code int) ([]byte, error) {
	return GetProtoData(&gateMsg.ResultPopResp{
		Title: StatusText[Title001],
		Flag:  int32(code),
		Hints: StatusText[code],
	})
}

// GetProtoResultPopWarn 获取协议结构数据 msgID + data.len + data
func GetProtoResultPopWarn(code int) ([]byte, error) {
	return GetProtoData(&gateMsg.ResultPopResp{
		Title: StatusText[Title002],
		Flag:  int32(code),
		Hints: StatusText[code],
	})
}

// GetProtoResultPopFatal 获取协议结构数据 msgID + data.len + data
func GetProtoResultPopFatal(code int) ([]byte, error) {
	return GetProtoData(&gateMsg.ResultPopResp{
		Title: StatusText[Title005],
		Flag:  int32(code),
		Hints: StatusText[code],
	})
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

func ParseResult(code int) (uint32, []byte, error) {
	return ParseProto(&gateMsg.ResultResp{
		State: int32(code),
		Hints: StatusText[code],
	})
}

func ParseResultPop(code int) (uint32, []byte, error) {
	return ParseProto(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title001],
		Hints: StatusText[code],
	})
}

func ParseResultPopWarn(code int) (uint32, []byte, error) {
	return ParseProto(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title002],
		Hints: StatusText[code],
	})
}
func ParseResultPopFatal(code int) (uint32, []byte, error) {
	return ParseProto(&gateMsg.ResultPopResp{
		Flag:  int32(code),
		Title: StatusText[Title005],
		Hints: StatusText[code],
	})
}
