// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: rpc.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 节点状态
type NodeState int32

const (
	NodeState_Open     NodeState = 0 // 开放状态,所有角色可以进入
	NodeState_Maintain NodeState = 1 // 维护状态,白名单角色可以进入
	NodeState_Closed   NodeState = 2 // 关闭状态,所有角色不可进入(同时踢除所有角色)
)

// Enum value maps for NodeState.
var (
	NodeState_name = map[int32]string{
		0: "Open",
		1: "Maintain",
		2: "Closed",
	}
	NodeState_value = map[string]int32{
		"Open":     0,
		"Maintain": 1,
		"Closed":   2,
	}
)

func (x NodeState) Enum() *NodeState {
	p := new(NodeState)
	*p = x
	return p
}

func (x NodeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeState) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}

func (NodeState) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}

func (x NodeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeState.Descriptor instead.
func (NodeState) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

// 注册开发帐号
type DevRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"` // 帐号名
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`       // 密码
	Ip          string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`                   // ip地址
}

func (x *DevRegister) Reset() {
	*x = DevRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevRegister) ProtoMessage() {}

func (x *DevRegister) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevRegister.ProtoReflect.Descriptor instead.
func (*DevRegister) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DevRegister) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *DevRegister) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DevRegister) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// 用户信息
type GetUserIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SdkId    int32  `protobuf:"varint,1,opt,name=sdkId,proto3" json:"sdkId,omitempty"`       // sdk id
	Pid      int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`           // 包id
	OpenId   string `protobuf:"bytes,3,opt,name=openId,proto3" json:"openId,omitempty"`      // sdk的openid
	ServerId int32  `protobuf:"varint,4,opt,name=serverId,proto3" json:"serverId,omitempty"` // 所在游戏服id
}

func (x *GetUserIDReq) Reset() {
	*x = GetUserIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIDReq) ProtoMessage() {}

func (x *GetUserIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIDReq.ProtoReflect.Descriptor instead.
func (*GetUserIDReq) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserIDReq) GetSdkId() int32 {
	if x != nil {
		return x.SdkId
	}
	return 0
}

func (x *GetUserIDReq) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *GetUserIDReq) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *GetUserIDReq) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type GetUserIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 用户id
}

func (x *GetUserIDResp) Reset() {
	*x = GetUserIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIDResp) ProtoMessage() {}

func (x *GetUserIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIDResp.ProtoReflect.Descriptor instead.
func (*GetUserIDResp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserIDResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x5b, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x6a, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x64, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x64, 0x6b,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x2a, 0x2f, 0x0a, 0x09, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0x02, 0x42, 0x05, 0x5a, 0x03,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_proto_goTypes = []interface{}{
	(NodeState)(0),        // 0: pb.NodeState
	(*DevRegister)(nil),   // 1: pb.DevRegister
	(*GetUserIDReq)(nil),  // 2: pb.GetUserIDReq
	(*GetUserIDResp)(nil), // 3: pb.GetUserIDResp
}
var file_rpc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevRegister); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIDReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIDResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
