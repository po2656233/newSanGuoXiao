package process

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/po2656233/superplace/logger"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
	"strings"
	. "superman/internal/constant"
	"superman/internal/redis_cluster"
)

const (
	protoPackageName = "pb"
)

// -------------------------
// | id | protobuf message |
// -------------------------

type Processor struct {
	littleEndian bool
	msgInfo      map[uint16]*MsgInfo
	msgID        map[reflect.Type]uint16
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}

type MsgHandler func([]interface{})

type MsgRaw struct {
	msgID      uint16
	msgRawData []byte
}

var redisHandle = redis_cluster.SingleRedis()

func NewProcessor(isLittleEndian bool) *Processor {
	p := new(Processor)
	p.littleEndian = isLittleEndian
	p.msgID = make(map[reflect.Type]uint16)
	return p
}

// SetByteOrder It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// Register It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(msg proto.Message) (uint16, string) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("protobuf message pointer required")
		return 0, ""
	}
	if _, ok := p.msgID[msgType]; ok {
		log.Fatal("message %s is already registered", msgType)
		return 0, ""
	}

	// 获取最大消息ID
	zVal, err := redisHandle.GetTopCount(context.Background(), KeyMsgProto, 1)
	if err != nil && err.Error() != StatusText[Redis01] {
		log.Fatal("max msgId can't set in redis_cluster err:%v", err)
	}
	fId := float64(0)
	if 0 < len(zVal) {
		fId = zVal[0].Score
	}
	maxId := uint16(fId)

	// 结构体最大判断
	size := len(p.msgInfo)
	if 0 == size {
		p.msgInfo = make(map[uint16]*MsgInfo)
	}
	if size >= math.MaxUint16 || fId >= math.MaxUint16 {
		log.Fatal("too many protobuf messages maxId:%f (max = %v) ", fId, math.MaxUint16)
	}

	//根据消息结构体名称 从redis里获取消息ID,若获取不到则从最大消息ID中
	// 消息体名称
	packageName := "*" + protoPackageName + "."
	strMsg := msgType.String()
	if strings.Contains(strMsg, packageName) {
		strMsg = strMsg[len(packageName):]
	}

	// 获取消息ID
	id := maxId + 1
	fVal := redisHandle.DB.ZScore(context.Background(), KeyMsgProto, strMsg)
	fId = fVal.Val()
	if fVal.Err() != nil || uint16(fId) == 0 {
		err = redisHandle.DB.ZAdd(context.Background(), KeyMsgProto, &redis.Z{
			Score:  float64(id),
			Member: strMsg,
		}).Err()
		log.Warnf("[warn]%v was id in redis_cluster. ID(%v) err:%v", strMsg, id, err)
	} else {
		id = uint16(fId)
	}
	i := new(MsgInfo)
	i.msgType = msgType
	p.msgInfo[id] = i
	p.msgID[msgType] = id
	return id, strMsg
}

// SetRouter It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRouter(msg proto.Message, msgRouter *Server) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	p.msgInfo[id].msgRouter = msgRouter
}

// SetHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetHandler(msg proto.Message, msgHandler MsgHandler) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	p.msgInfo[id].msgHandler = msgHandler
}

// SetRawHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRawHandler(id uint16, msgRawHandler MsgHandler) {
	if id >= uint16(len(p.msgInfo)) {
		log.Fatal("message id %v not registered", id)
	}

	p.msgInfo[id].msgRawHandler = msgRawHandler
}

// Route goroutine safe
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	// raw
	if msgRaw, ok := msg.(MsgRaw); ok {
		if msgRaw.msgID >= uint16(len(p.msgInfo)) {
			return fmt.Errorf("message id %v not registered", msgRaw.msgID)
		}
		i := p.msgInfo[msgRaw.msgID]
		if i.msgRawHandler != nil {
			i.msgRawHandler([]interface{}{msgRaw.msgID, msgRaw.msgRawData, userData})
		}
		return nil
	}

	// protobuf
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		return fmt.Errorf("message %s not registered", msgType)
	}
	i := p.msgInfo[id]
	if i.msgHandler != nil {
		i.msgHandler([]interface{}{msg, userData})
	}
	if i.msgRouter != nil {
		i.msgRouter.Go(msgType, msg, userData)
	}
	return nil
}

// Unmarshal goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}

	// id
	var id uint16
	if p.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}

	// msg
	i, ok := p.msgInfo[id]
	if !ok {
		return nil, fmt.Errorf("message id %v not registered", id)
	}
	if i.msgRawHandler != nil {
		return MsgRaw{id, data[2:]}, nil
	} else {
		msg := reflect.New(i.msgType.Elem()).Interface()
		return msg, proto.Unmarshal(data[2:], msg.(proto.Message))
	}
}

// Marshal goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	msgType := reflect.TypeOf(msg)

	// id
	_id, ok := p.msgID[msgType]
	if !ok {
		err := fmt.Errorf("message %s not registered", msgType)
		return nil, err
	}

	id := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(id, _id)
	} else {
		binary.BigEndian.PutUint16(id, _id)
	}

	// data
	data, err := proto.Marshal(msg.(proto.Message))
	return [][]byte{id, data}, err
}

// Range goroutine safe
func (p *Processor) Range(f func(id uint16, t reflect.Type)) {
	for id, i := range p.msgInfo {
		f(id, i.msgType)
	}
}
