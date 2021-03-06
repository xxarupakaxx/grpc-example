// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: pb/player.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Result int32

const (
	Result_DRAW Result = 0
	Result_LOSE Result = 1
	Result_WIN  Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "DRAW",
		1: "LOSE",
		2: "WIN",
	}
	Result_value = map[string]int32{
		"DRAW": 0,
		"LOSE": 1,
		"WIN":  2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_player_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_pb_player_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{0}
}

type Piece_XO int32

const (
	Piece_X       Piece_XO = 0
	Piece_O       Piece_XO = 1
	Piece_UNKNOWN Piece_XO = 2
)

// Enum value maps for Piece_XO.
var (
	Piece_XO_name = map[int32]string{
		0: "X",
		1: "O",
		2: "UNKNOWN",
	}
	Piece_XO_value = map[string]int32{
		"X":       0,
		"O":       1,
		"UNKNOWN": 2,
	}
)

func (x Piece_XO) Enum() *Piece_XO {
	p := new(Piece_XO)
	*p = x
	return p
}

func (x Piece_XO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Piece_XO) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_player_proto_enumTypes[1].Descriptor()
}

func (Piece_XO) Type() protoreflect.EnumType {
	return &file_pb_player_proto_enumTypes[1]
}

func (x Piece_XO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Piece_XO.Descriptor instead.
func (Piece_XO) EnumDescriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{3, 0}
}

type PlayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID int32   `protobuf:"varint,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	// Types that are assignable to Action:
	//	*PlayRequest_Start
	//	*PlayRequest_Playing
	Action isPlayRequest_Action `protobuf_oneof:"Action"`
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{0}
}

func (x *PlayRequest) GetRoomID() int32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

func (x *PlayRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (m *PlayRequest) GetAction() isPlayRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *PlayRequest) GetStart() *PlayRequest_StartAction {
	if x, ok := x.GetAction().(*PlayRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *PlayRequest) GetPlaying() *PlayRequest_PlayAction {
	if x, ok := x.GetAction().(*PlayRequest_Playing); ok {
		return x.Playing
	}
	return nil
}

type isPlayRequest_Action interface {
	isPlayRequest_Action()
}

type PlayRequest_Start struct {
	Start *PlayRequest_StartAction `protobuf:"bytes,3,opt,name=start,proto3,oneof"`
}

type PlayRequest_Playing struct {
	Playing *PlayRequest_PlayAction `protobuf:"bytes,4,opt,name=playing,proto3,oneof"`
}

func (*PlayRequest_Start) isPlayRequest_Action() {}

func (*PlayRequest_Playing) isPlayRequest_Action() {}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*PlayResponse_Ready
	//	*PlayResponse_Play
	//	*PlayResponse_Finish
	Event isPlayResponse_Event `protobuf_oneof:"Event"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{1}
}

func (m *PlayResponse) GetEvent() isPlayResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *PlayResponse) GetReady() *PlayResponse_ReadyEvent {
	if x, ok := x.GetEvent().(*PlayResponse_Ready); ok {
		return x.Ready
	}
	return nil
}

func (x *PlayResponse) GetPlay() *PlayResponse_PlayEvent {
	if x, ok := x.GetEvent().(*PlayResponse_Play); ok {
		return x.Play
	}
	return nil
}

func (x *PlayResponse) GetFinish() *PlayResponse_FinishEvent {
	if x, ok := x.GetEvent().(*PlayResponse_Finish); ok {
		return x.Finish
	}
	return nil
}

type isPlayResponse_Event interface {
	isPlayResponse_Event()
}

type PlayResponse_Ready struct {
	Ready *PlayResponse_ReadyEvent `protobuf:"bytes,1,opt,name=ready,proto3,oneof"`
}

type PlayResponse_Play struct {
	Play *PlayResponse_PlayEvent `protobuf:"bytes,2,opt,name=play,proto3,oneof"`
}

type PlayResponse_Finish struct {
	Finish *PlayResponse_FinishEvent `protobuf:"bytes,3,opt,name=finish,proto3,oneof"`
}

func (*PlayResponse_Ready) isPlayResponse_Event() {}

func (*PlayResponse_Play) isPlayResponse_Event() {}

func (*PlayResponse_Finish) isPlayResponse_Event() {}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID int32  `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Piece    *Piece `protobuf:"bytes,2,opt,name=piece,proto3" json:"piece,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetPlayerID() int32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *Player) GetPiece() *Piece {
	if x != nil {
		return x.Piece
	}
	return nil
}

type Piece struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xo Piece_XO `protobuf:"varint,1,opt,name=xo,proto3,enum=example.Piece_XO" json:"xo,omitempty"`
}

func (x *Piece) Reset() {
	*x = Piece{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piece) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piece) ProtoMessage() {}

func (x *Piece) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piece.ProtoReflect.Descriptor instead.
func (*Piece) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{3}
}

func (x *Piece) GetXo() Piece_XO {
	if x != nil {
		return x.Xo
	}
	return Piece_X
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pieces []*Piece `protobuf:"bytes,1,rep,name=pieces,proto3" json:"pieces,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{4}
}

func (x *Board) GetPieces() []*Piece {
	if x != nil {
		return x.Pieces
	}
	return nil
}

type PlayRequest_StartAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayRequest_StartAction) Reset() {
	*x = PlayRequest_StartAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest_StartAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest_StartAction) ProtoMessage() {}

func (x *PlayRequest_StartAction) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest_StartAction.ProtoReflect.Descriptor instead.
func (*PlayRequest_StartAction) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{0, 0}
}

type PlayRequest_PlayAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceNumber int32 `protobuf:"varint,1,opt,name=placeNumber,proto3" json:"placeNumber,omitempty"`
}

func (x *PlayRequest_PlayAction) Reset() {
	*x = PlayRequest_PlayAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest_PlayAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest_PlayAction) ProtoMessage() {}

func (x *PlayRequest_PlayAction) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest_PlayAction.ProtoReflect.Descriptor instead.
func (*PlayRequest_PlayAction) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PlayRequest_PlayAction) GetPlaceNumber() int32 {
	if x != nil {
		return x.PlaceNumber
	}
	return 0
}

type PlayResponse_ReadyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayResponse_ReadyEvent) Reset() {
	*x = PlayResponse_ReadyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_ReadyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_ReadyEvent) ProtoMessage() {}

func (x *PlayResponse_ReadyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_ReadyEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_ReadyEvent) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{1, 0}
}

type PlayResponse_PlayEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Board  *Board  `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *PlayResponse_PlayEvent) Reset() {
	*x = PlayResponse_PlayEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_PlayEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_PlayEvent) ProtoMessage() {}

func (x *PlayResponse_PlayEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_PlayEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_PlayEvent) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PlayResponse_PlayEvent) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *PlayResponse_PlayEvent) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type PlayResponse_FinishEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=example.Result" json:"result,omitempty"`
	Board  *Board `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *PlayResponse_FinishEvent) Reset() {
	*x = PlayResponse_FinishEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_FinishEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_FinishEvent) ProtoMessage() {}

func (x *PlayResponse_FinishEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_FinishEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_FinishEvent) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{1, 2}
}

func (x *PlayResponse_FinishEvent) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_DRAW
}

func (x *PlayResponse_FinishEvent) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

var File_pb_player_proto protoreflect.FileDescriptor

var file_pb_player_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x8e, 0x02, 0x0a, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x1a, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x2e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x08, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x03, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x35, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x3b, 0x0a,
	0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x1a, 0x0c, 0x0a, 0x0a, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x5a, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x1a, 0x5c, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65,
	0x52, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x05, 0x50, 0x69, 0x65, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x02, 0x78, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x2e, 0x58, 0x4f, 0x52,
	0x02, 0x78, 0x6f, 0x22, 0x1f, 0x0a, 0x02, 0x58, 0x4f, 0x12, 0x05, 0x0a, 0x01, 0x58, 0x10, 0x00,
	0x12, 0x05, 0x0a, 0x01, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x22, 0x2f, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x26, 0x0a,
	0x06, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x06, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x73, 0x2a, 0x25, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x53,
	0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x49, 0x4e, 0x10, 0x02, 0x32, 0x46, 0x0a, 0x0b,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x79, 0x12, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x78, 0x61, 0x72, 0x75, 0x70, 0x61, 0x6b, 0x61, 0x78, 0x78, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_player_proto_rawDescOnce sync.Once
	file_pb_player_proto_rawDescData = file_pb_player_proto_rawDesc
)

func file_pb_player_proto_rawDescGZIP() []byte {
	file_pb_player_proto_rawDescOnce.Do(func() {
		file_pb_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_player_proto_rawDescData)
	})
	return file_pb_player_proto_rawDescData
}

var file_pb_player_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_player_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_player_proto_goTypes = []interface{}{
	(Result)(0),                      // 0: example.Result
	(Piece_XO)(0),                    // 1: example.Piece.XO
	(*PlayRequest)(nil),              // 2: example.PlayRequest
	(*PlayResponse)(nil),             // 3: example.PlayResponse
	(*Player)(nil),                   // 4: example.Player
	(*Piece)(nil),                    // 5: example.Piece
	(*Board)(nil),                    // 6: example.Board
	(*PlayRequest_StartAction)(nil),  // 7: example.PlayRequest.StartAction
	(*PlayRequest_PlayAction)(nil),   // 8: example.PlayRequest.PlayAction
	(*PlayResponse_ReadyEvent)(nil),  // 9: example.PlayResponse.ReadyEvent
	(*PlayResponse_PlayEvent)(nil),   // 10: example.PlayResponse.PlayEvent
	(*PlayResponse_FinishEvent)(nil), // 11: example.PlayResponse.FinishEvent
}
var file_pb_player_proto_depIdxs = []int32{
	4,  // 0: example.PlayRequest.player:type_name -> example.Player
	7,  // 1: example.PlayRequest.start:type_name -> example.PlayRequest.StartAction
	8,  // 2: example.PlayRequest.playing:type_name -> example.PlayRequest.PlayAction
	9,  // 3: example.PlayResponse.ready:type_name -> example.PlayResponse.ReadyEvent
	10, // 4: example.PlayResponse.play:type_name -> example.PlayResponse.PlayEvent
	11, // 5: example.PlayResponse.finish:type_name -> example.PlayResponse.FinishEvent
	5,  // 6: example.Player.piece:type_name -> example.Piece
	1,  // 7: example.Piece.xo:type_name -> example.Piece.XO
	5,  // 8: example.Board.pieces:type_name -> example.Piece
	4,  // 9: example.PlayResponse.PlayEvent.player:type_name -> example.Player
	6,  // 10: example.PlayResponse.PlayEvent.board:type_name -> example.Board
	0,  // 11: example.PlayResponse.FinishEvent.result:type_name -> example.Result
	6,  // 12: example.PlayResponse.FinishEvent.board:type_name -> example.Board
	2,  // 13: example.GameService.Play:input_type -> example.PlayRequest
	3,  // 14: example.GameService.Play:output_type -> example.PlayResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_pb_player_proto_init() }
func file_pb_player_proto_init() {
	if File_pb_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest); i {
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
		file_pb_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
		file_pb_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_pb_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piece); i {
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
		file_pb_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_pb_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest_StartAction); i {
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
		file_pb_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest_PlayAction); i {
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
		file_pb_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_ReadyEvent); i {
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
		file_pb_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_PlayEvent); i {
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
		file_pb_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_FinishEvent); i {
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
	file_pb_player_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PlayRequest_Start)(nil),
		(*PlayRequest_Playing)(nil),
	}
	file_pb_player_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PlayResponse_Ready)(nil),
		(*PlayResponse_Play)(nil),
		(*PlayResponse_Finish)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_player_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_player_proto_goTypes,
		DependencyIndexes: file_pb_player_proto_depIdxs,
		EnumInfos:         file_pb_player_proto_enumTypes,
		MessageInfos:      file_pb_player_proto_msgTypes,
	}.Build()
	File_pb_player_proto = out.File
	file_pb_player_proto_rawDesc = nil
	file_pb_player_proto_goTypes = nil
	file_pb_player_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (GameService_PlayClient, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Play(ctx context.Context, opts ...grpc.CallOption) (GameService_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GameService_serviceDesc.Streams[0], "/example.GameService/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServicePlayClient{stream}
	return x, nil
}

type GameService_PlayClient interface {
	Send(*PlayRequest) error
	Recv() (*PlayResponse, error)
	grpc.ClientStream
}

type gameServicePlayClient struct {
	grpc.ClientStream
}

func (x *gameServicePlayClient) Send(m *PlayRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServicePlayClient) Recv() (*PlayResponse, error) {
	m := new(PlayResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	Play(GameService_PlayServer) error
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) Play(GameService_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).Play(&gameServicePlayServer{stream})
}

type GameService_PlayServer interface {
	Send(*PlayResponse) error
	Recv() (*PlayRequest, error)
	grpc.ServerStream
}

type gameServicePlayServer struct {
	grpc.ServerStream
}

func (x *gameServicePlayServer) Send(m *PlayResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServicePlayServer) Recv() (*PlayRequest, error) {
	m := new(PlayRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _GameService_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/player.proto",
}
