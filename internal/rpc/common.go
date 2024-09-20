package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/po2656233/superplace/const/code"
	exReflect "github.com/po2656233/superplace/extend/reflect"
	"github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"os"
	"reflect"
	"strconv"
	"strings"
	. "superman/internal/constant"
	pb "superman/internal/protocol/gofile"
)

const (
	Request1  = "Req"
	Request2  = "Request"
	Response1 = "Resp"
	Response2 = "Response"
)

var (
	MapIdMsg = make(map[string]string)
	Endian   = binary.BigEndian
)

// MessageSender 是一个封装类，用于减少SendData的传参
type MessageSender struct {
	app      facade.IApplication
	source   string
	actorID  string
	nodeType string
}

// NewMessageSender 创建一个新的MessageSender实例
func NewMessageSender(app facade.IApplication, source, actorID, nodeType string) *MessageSender {
	return &MessageSender{
		app:      app,
		source:   source,
		actorID:  actorID,
		nodeType: nodeType,
	}
}

// SendData 返送请求
func (ms *MessageSender) SendData(req proto.Message) interface{} {
	resp, errCode := SendData(ms.app, ms.source, ms.actorID, ms.nodeType, req)
	if code.IsFail(errCode) {
		clog.Warnf("[SendData] source:%s actor:%v node:%s req:%+v errCode:%v", ms.source, ms.actorID, ms.nodeType, req, errCode)
	}
	return resp
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// LoadMsgInfos 获取消息映射文件message_id.json
func LoadMsgInfos() map[string]string {
	data, _ := os.ReadFile(MSGFile)
	//创建map，用于接收解码好的数据
	//创建文件的解码器
	_ = json.Unmarshal(data, &MapIdMsg)
	//解码文件中的数据，丢入dataMap所在的内存
	//fmt.Println("解码成功", mapFuncs)
	return MapIdMsg
}

// GetProtoData 获取协议结构数据 msgID + data.len + data
func GetProtoData(msg proto.Message) ([]byte, error) {
	strId := ""
	name := reflect.Indirect(reflect.ValueOf(msg)).Type().Name()
	for id, protoName := range MapIdMsg {
		if protoName == name {
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

// Ping 实现了Ping功能
func Ping(app facade.IApplication) bool {
	_, errCode := SendDataToOps(app, &pb.PingReq{})
	return errCode == code.OK
}

// SendData 发送数据
func SendData(app facade.IApplication, source, actorID, nodeType string, req proto.Message) (resp interface{}, errCode int32) {
	descriptor := req.ProtoReflect().Descriptor()
	if descriptor == nil {
		clog.Errorf("[SendData] descriptor err:descriptor is nil")
		return resp, NetworkErr02
	}

	// 根据请求类型 创建回包数据类型便于数据序列化于结构体中
	respTypeName := ""
	funcName := exReflect.GetStructName(req)
	fullName := string(descriptor.FullName())
	fSize := len(fullName)
	r1Size := len(Request1)
	r2Size := len(Request2)
	if funcName == "PingReq" {
		funcName = "Ping"
		respTypeName = strings.Replace(fullName, "PingReq", "PongResp", -1)
	} else if r1Size < fSize && fullName[fSize-r1Size:] == Request1 {
		respTypeName = fullName[:fSize-r1Size] + Response1
		funcName = strings.TrimSuffix(funcName, Request1)
	} else if r2Size < fSize && fullName[fSize-r2Size:] == Request2 {
		respTypeName = fullName[:fSize-r2Size] + Response2
		//funcName = strings.TrimSuffix(funcName, Request2)
	}

	target := GetTargetPath(app, actorID, nodeType)
	if respTypeName != "" {
		msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(respTypeName))
		if err != nil {
			clog.Warnf("[SendData] sourcePath:%s targetPath:%v reqData:%+v no resp! err:%v", source, target, req, err)
		} else {
			resp = msgType.New().Interface()
		}
	}

	errCode = app.ActorSystem().CallWait(source, target, funcName, req, resp)
	if code.IsFail(errCode) {
		clog.Warnf("[SendData] sourcePath:%s targetPath:%v funcName:%s reqData:%+v errCode:%v", source, target, funcName, req, errCode)
	}
	return resp, errCode
}

// SendDataToDB 发送给数据库节点
func SendDataToDB(app facade.IApplication, req proto.Message) (interface{}, int32) {
	return SendData(app, SourcePath, DBActor, NodeTypeCenter, req)
}

// SendDataToAcc 发送给账号节点
func SendDataToAcc(app facade.IApplication, req proto.Message) (interface{}, int32) {
	return SendData(app, SourcePath, AccActor, NodeTypeCenter, req)
}
func SendDataToOps(app facade.IApplication, req proto.Message) (interface{}, int32) {
	return SendData(app, SourcePath, OpsActor, NodeTypeCenter, req)
}
func SendDataToGame(app facade.IApplication, req proto.Message) (interface{}, int32) {
	return SendData(app, SourcePath, GameActor, NodeTypeCenter, req)
}
func GetTargetPath(app facade.IApplication, actorID, nodeType string) string {
	list := app.Discovery().ListByType(nodeType)
	if len(list) == 0 {
		return ""
	}
	return list[0].GetNodeId() + actorID
}

// [废弃]resp是回包内容,该交由发起请求方处理
//var chServer = NewServer(10000)
//var mapResp = make(map[reflect.Type]interface{})

//func init() {
//	registerMsg(&pb.RegisterReq{}, account.RegisterResp)
//	registerMsg(&pb.LoginReq{}, account.LoginResp)
//}
//
//func registerMsg(m, h interface{}) {
//	descriptor := m.(proto.Message).ProtoReflect().Descriptor()
//	if descriptor == nil {
//		clog.Errorf("[rpc]registerMsg descriptor err: %v", m)
//		return
//	}
//	respTypeName := utils.ReplaceLast(string(descriptor.FullName()), "Req", "Resp")
//	msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(respTypeName))
//	if err != nil {
//		clog.Errorf("[rpc]registerMsg FindMessageByName err: %v", err)
//		return
//	}
//	mapResp[reflect.TypeOf(m)] = msgType.New().Interface()
//	//chServer.Register(reflect.TypeOf(m), h)
//}
//
//
//
//// SendData 发送数据
//func SendData(app facade.IApplication, source, target string, req proto.Message) int32 {
//	reqType := reflect.TypeOf(req)
//	funcName := exReflect.GetStructName(req)
//	resp, ok := mapResp[reqType]
//	if !ok {
//		clog.Warnf("[SendData] sourcePath:%s targetPath:%v funcName:%s reqData:%+v no resp!", source, target, funcName, req)
//	}
//	errCode := app.ActorSystem().CallWait(source, target, funcName, req, resp)
//	if code.IsFail(errCode) {
//		clog.Warnf("[SendData] sourcePath:%s targetPath:%v funcName:%s reqData:%+v errCode:%v", source, target, funcName, req, errCode)
//	} else if resp != nil {
//		// 分派处理回包内容
//		//chServer.Go(reqType, resp)
//	}
//	return errCode
//}
