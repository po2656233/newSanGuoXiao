// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: baccarat.proto

package pb

import (
	
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


// kindID 2001
//场景
type BaccaratSceneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp  int64       `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`      //时间戳
	Inning     string      `protobuf:"bytes,2,opt,name=inning,proto3" json:"inning,omitempty"`             //牌局号
	Chips      []int32     `protobuf:"varint,3,rep,packed,name=chips,proto3" json:"chips,omitempty"`       //筹 码
	AwardAreas [][]byte    `protobuf:"bytes,4,rep,name=awardAreas,proto3" json:"awardAreas,omitempty"`     //开奖记录(路单)
	AreaBets   []int64     `protobuf:"varint,5,rep,packed,name=areaBets,proto3" json:"areaBets,omitempty"` //各下注区当前总下注额
	MyBets     []int64     `protobuf:"varint,6,rep,packed,name=myBets,proto3" json:"myBets,omitempty"`     //个人在各下注区的总下注额
	AllPlayers *PlayerList `protobuf:"bytes,7,opt,name=allPlayers,proto3" json:"allPlayers,omitempty"`     //玩家列表
}

func (x *BaccaratSceneResp) Reset() {
	*x = BaccaratSceneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratSceneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratSceneResp) ProtoMessage() {}

func (x *BaccaratSceneResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratSceneResp.ProtoReflect.Descriptor instead.
func (*BaccaratSceneResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{0}
}

func (x *BaccaratSceneResp) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *BaccaratSceneResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

func (x *BaccaratSceneResp) GetChips() []int32 {
	if x != nil {
		return x.Chips
	}
	return nil
}

func (x *BaccaratSceneResp) GetAwardAreas() [][]byte {
	if x != nil {
		return x.AwardAreas
	}
	return nil
}

func (x *BaccaratSceneResp) GetAreaBets() []int64 {
	if x != nil {
		return x.AreaBets
	}
	return nil
}

func (x *BaccaratSceneResp) GetMyBets() []int64 {
	if x != nil {
		return x.MyBets
	}
	return nil
}

func (x *BaccaratSceneResp) GetAllPlayers() *PlayerList {
	if x != nil {
		return x.AllPlayers
	}
	return nil
}

//状态
// 服务端推送
//(准备)
type BaccaratStateStartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times  *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
	Inning string    `protobuf:"bytes,2,opt,name=inning,proto3" json:"inning,omitempty"` // 牌局号
}

func (x *BaccaratStateStartResp) Reset() {
	*x = BaccaratStateStartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratStateStartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratStateStartResp) ProtoMessage() {}

func (x *BaccaratStateStartResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratStateStartResp.ProtoReflect.Descriptor instead.
func (*BaccaratStateStartResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{1}
}

func (x *BaccaratStateStartResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *BaccaratStateStartResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

//(下注)
type BaccaratStatePlayingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *BaccaratStatePlayingResp) Reset() {
	*x = BaccaratStatePlayingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratStatePlayingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratStatePlayingResp) ProtoMessage() {}

func (x *BaccaratStatePlayingResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratStatePlayingResp.ProtoReflect.Descriptor instead.
func (*BaccaratStatePlayingResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{2}
}

func (x *BaccaratStatePlayingResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//(开奖)
type BaccaratStateOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times    *TimeInfo         `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
	OpenInfo *BaccaratOpenResp `protobuf:"bytes,2,opt,name=openInfo,proto3" json:"openInfo,omitempty"`
}

func (x *BaccaratStateOpenResp) Reset() {
	*x = BaccaratStateOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratStateOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratStateOpenResp) ProtoMessage() {}

func (x *BaccaratStateOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratStateOpenResp.ProtoReflect.Descriptor instead.
func (*BaccaratStateOpenResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{3}
}

func (x *BaccaratStateOpenResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *BaccaratStateOpenResp) GetOpenInfo() *BaccaratOpenResp {
	if x != nil {
		return x.OpenInfo
	}
	return nil
}

//(结束)
type BaccaratStateOverResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *BaccaratStateOverResp) Reset() {
	*x = BaccaratStateOverResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratStateOverResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratStateOverResp) ProtoMessage() {}

func (x *BaccaratStateOverResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratStateOverResp.ProtoReflect.Descriptor instead.
func (*BaccaratStateOverResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{4}
}

func (x *BaccaratStateOverResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//游戏消息
//抢庄
type BaccaratHostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWant bool `protobuf:"varint,1,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消上庄
}

func (x *BaccaratHostReq) Reset() {
	*x = BaccaratHostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratHostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratHostReq) ProtoMessage() {}

func (x *BaccaratHostReq) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratHostReq.ProtoReflect.Descriptor instead.
func (*BaccaratHostReq) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{5}
}

func (x *BaccaratHostReq) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

type BaccaratHostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsWant bool  `protobuf:"varint,2,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消上庄
}

func (x *BaccaratHostResp) Reset() {
	*x = BaccaratHostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratHostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratHostResp) ProtoMessage() {}

func (x *BaccaratHostResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratHostResp.ProtoReflect.Descriptor instead.
func (*BaccaratHostResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{6}
}

func (x *BaccaratHostResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BaccaratHostResp) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

//超级抢庄
type BaccaratSuperHostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWant bool `protobuf:"varint,1,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消
}

func (x *BaccaratSuperHostReq) Reset() {
	*x = BaccaratSuperHostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratSuperHostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratSuperHostReq) ProtoMessage() {}

func (x *BaccaratSuperHostReq) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratSuperHostReq.ProtoReflect.Descriptor instead.
func (*BaccaratSuperHostReq) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{7}
}

func (x *BaccaratSuperHostReq) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

type BaccaratSuperHostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsWant bool  `protobuf:"varint,2,opt,name=isWant,proto3" json:"isWant,omitempty"` //true上庄 false取消
}

func (x *BaccaratSuperHostResp) Reset() {
	*x = BaccaratSuperHostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratSuperHostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratSuperHostResp) ProtoMessage() {}

func (x *BaccaratSuperHostResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratSuperHostResp.ProtoReflect.Descriptor instead.
func (*BaccaratSuperHostResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{8}
}

func (x *BaccaratSuperHostResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BaccaratSuperHostResp) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

//下注
type BaccaratBetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BetArea  int32 `protobuf:"varint,1,opt,name=betArea,proto3" json:"betArea,omitempty"`   //下注区域
	BetScore int64 `protobuf:"varint,2,opt,name=betScore,proto3" json:"betScore,omitempty"` //下注金额
}

func (x *BaccaratBetReq) Reset() {
	*x = BaccaratBetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratBetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratBetReq) ProtoMessage() {}

func (x *BaccaratBetReq) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratBetReq.ProtoReflect.Descriptor instead.
func (*BaccaratBetReq) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{9}
}

func (x *BaccaratBetReq) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *BaccaratBetReq) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

type BaccaratBetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	BetArea  int32 `protobuf:"varint,2,opt,name=betArea,proto3" json:"betArea,omitempty"`   //下注区域
	BetScore int64 `protobuf:"varint,3,opt,name=betScore,proto3" json:"betScore,omitempty"` //下注金额
}

func (x *BaccaratBetResp) Reset() {
	*x = BaccaratBetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratBetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratBetResp) ProtoMessage() {}

func (x *BaccaratBetResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratBetResp.ProtoReflect.Descriptor instead.
func (*BaccaratBetResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{10}
}

func (x *BaccaratBetResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BaccaratBetResp) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *BaccaratBetResp) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

//结束
type BaccaratOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwardArea  []byte    `protobuf:"bytes,1,opt,name=awardArea,proto3" json:"awardArea,omitempty"`   //开奖区域
	PlayerCard *CardInfo `protobuf:"bytes,2,opt,name=playerCard,proto3" json:"playerCard,omitempty"` //闲家的牌
	BankerCard *CardInfo `protobuf:"bytes,3,opt,name=bankerCard,proto3" json:"bankerCard,omitempty"` //庄家的牌
}

func (x *BaccaratOpenResp) Reset() {
	*x = BaccaratOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratOpenResp) ProtoMessage() {}

func (x *BaccaratOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratOpenResp.ProtoReflect.Descriptor instead.
func (*BaccaratOpenResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{11}
}

func (x *BaccaratOpenResp) GetAwardArea() []byte {
	if x != nil {
		return x.AwardArea
	}
	return nil
}

func (x *BaccaratOpenResp) GetPlayerCard() *CardInfo {
	if x != nil {
		return x.PlayerCard
	}
	return nil
}

func (x *BaccaratOpenResp) GetBankerCard() *CardInfo {
	if x != nil {
		return x.BankerCard
	}
	return nil
}

//结算
type BaccaratCheckoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyAcquire int64   `protobuf:"varint,1,opt,name=myAcquire,proto3" json:"myAcquire,omitempty"`      //获得金币(结算)
	Acquires  []int64 `protobuf:"varint,2,rep,packed,name=acquires,proto3" json:"acquires,omitempty"` //
}

func (x *BaccaratCheckoutResp) Reset() {
	*x = BaccaratCheckoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baccarat_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaccaratCheckoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaccaratCheckoutResp) ProtoMessage() {}

func (x *BaccaratCheckoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_baccarat_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaccaratCheckoutResp.ProtoReflect.Descriptor instead.
func (*BaccaratCheckoutResp) Descriptor() ([]byte, []int) {
	return file_baccarat_proto_rawDescGZIP(), []int{12}
}

func (x *BaccaratCheckoutResp) GetMyAcquire() int64 {
	if x != nil {
		return x.MyAcquire
	}
	return 0
}

func (x *BaccaratCheckoutResp) GetAcquires() []int64 {
	if x != nil {
		return x.Acquires
	}
	return nil
}

var File_baccarat_proto protoreflect.FileDescriptor

var file_baccarat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61,
	0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x77, 0x61, 0x72, 0x64, 0x41,
	0x72, 0x65, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x77, 0x61, 0x72,
	0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x42, 0x65,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x42, 0x65,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x79, 0x42, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x06, 0x6d, 0x79, 0x42, 0x65, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x61, 0x6c,
	0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a,
	0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x54, 0x0a, 0x16, 0x42, 0x61,
	0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x22, 0x3e, 0x0a, 0x18, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x22, 0x6d, 0x0a, 0x15, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x4f, 0x70, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3b, 0x0a, 0x15, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0f,
	0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x63, 0x61,
	0x72, 0x61, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x42,
	0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x15, 0x42,
	0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x57, 0x61, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74,
	0x42, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x5f, 0x0a, 0x0f,
	0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72,
	0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x65, 0x74, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x10, 0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61,
	0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2c,
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x22, 0x50, 0x0a, 0x14,
	0x42, 0x61, 0x63, 0x63, 0x61, 0x72, 0x61, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x42, 0x05,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baccarat_proto_rawDescOnce sync.Once
	file_baccarat_proto_rawDescData = file_baccarat_proto_rawDesc
)

func file_baccarat_proto_rawDescGZIP() []byte {
	file_baccarat_proto_rawDescOnce.Do(func() {
		file_baccarat_proto_rawDescData = protoimpl.X.CompressGZIP(file_baccarat_proto_rawDescData)
	})
	return file_baccarat_proto_rawDescData
}

var file_baccarat_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_baccarat_proto_goTypes = []interface{}{
	(*BaccaratSceneResp)(nil),        // 0: pb.BaccaratSceneResp
	(*BaccaratStateStartResp)(nil),   // 1: pb.BaccaratStateStartResp
	(*BaccaratStatePlayingResp)(nil), // 2: pb.BaccaratStatePlayingResp
	(*BaccaratStateOpenResp)(nil),    // 3: pb.BaccaratStateOpenResp
	(*BaccaratStateOverResp)(nil),    // 4: pb.BaccaratStateOverResp
	(*BaccaratHostReq)(nil),          // 5: pb.BaccaratHostReq
	(*BaccaratHostResp)(nil),         // 6: pb.BaccaratHostResp
	(*BaccaratSuperHostReq)(nil),     // 7: pb.BaccaratSuperHostReq
	(*BaccaratSuperHostResp)(nil),    // 8: pb.BaccaratSuperHostResp
	(*BaccaratBetReq)(nil),           // 9: pb.BaccaratBetReq
	(*BaccaratBetResp)(nil),          // 10: pb.BaccaratBetResp
	(*BaccaratOpenResp)(nil),         // 11: pb.BaccaratOpenResp
	(*BaccaratCheckoutResp)(nil),     // 12: pb.BaccaratCheckoutResp
	(*PlayerList)(nil),               // 13: pb.PlayerList
	(*TimeInfo)(nil),                 // 14: pb.TimeInfo
	(*CardInfo)(nil),                 // 15: pb.CardInfo
}
var file_baccarat_proto_depIdxs = []int32{
	13, // 0: pb.BaccaratSceneResp.allPlayers:type_name -> pb.PlayerList
	14, // 1: pb.BaccaratStateStartResp.times:type_name -> pb.TimeInfo
	14, // 2: pb.BaccaratStatePlayingResp.times:type_name -> pb.TimeInfo
	14, // 3: pb.BaccaratStateOpenResp.times:type_name -> pb.TimeInfo
	11, // 4: pb.BaccaratStateOpenResp.openInfo:type_name -> pb.BaccaratOpenResp
	14, // 5: pb.BaccaratStateOverResp.times:type_name -> pb.TimeInfo
	15, // 6: pb.BaccaratOpenResp.playerCard:type_name -> pb.CardInfo
	15, // 7: pb.BaccaratOpenResp.bankerCard:type_name -> pb.CardInfo
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_baccarat_proto_init() }
func file_baccarat_proto_init() {
	if File_baccarat_proto != nil {
		return
	}
	file_baseinfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_baccarat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratSceneResp); i {
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
		file_baccarat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratStateStartResp); i {
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
		file_baccarat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratStatePlayingResp); i {
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
		file_baccarat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratStateOpenResp); i {
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
		file_baccarat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratStateOverResp); i {
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
		file_baccarat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratHostReq); i {
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
		file_baccarat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratHostResp); i {
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
		file_baccarat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratSuperHostReq); i {
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
		file_baccarat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratSuperHostResp); i {
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
		file_baccarat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratBetReq); i {
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
		file_baccarat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratBetResp); i {
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
		file_baccarat_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratOpenResp); i {
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
		file_baccarat_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaccaratCheckoutResp); i {
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
			RawDescriptor: file_baccarat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baccarat_proto_goTypes,
		DependencyIndexes: file_baccarat_proto_depIdxs,
		MessageInfos:      file_baccarat_proto_msgTypes,
	}.Build()
	File_baccarat_proto = out.File
	file_baccarat_proto_rawDesc = nil
	file_baccarat_proto_goTypes = nil
	file_baccarat_proto_depIdxs = nil
}
