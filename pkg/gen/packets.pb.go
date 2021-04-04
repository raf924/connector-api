// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: packets.proto

package gen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UserEvent int32

const (
	UserEvent_JOINED UserEvent = 0
	UserEvent_LEFT   UserEvent = 1
)

// Enum value maps for UserEvent.
var (
	UserEvent_name = map[int32]string{
		0: "JOINED",
		1: "LEFT",
	}
	UserEvent_value = map[string]int32{
		"JOINED": 0,
		"LEFT":   1,
	}
)

func (x UserEvent) Enum() *UserEvent {
	p := new(UserEvent)
	*p = x
	return p
}

func (x UserEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_packets_proto_enumTypes[0].Descriptor()
}

func (UserEvent) Type() protoreflect.EnumType {
	return &file_packets_proto_enumTypes[0]
}

func (x UserEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserEvent.Descriptor instead.
func (UserEvent) EnumDescriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{0}
}

type MessagePacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User      *User                `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Private   bool                 `protobuf:"varint,4,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *MessagePacket) Reset() {
	*x = MessagePacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePacket) ProtoMessage() {}

func (x *MessagePacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePacket.ProtoReflect.Descriptor instead.
func (*MessagePacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{0}
}

func (x *MessagePacket) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MessagePacket) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessagePacket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MessagePacket) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type UserPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	User      *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Event     UserEvent            `protobuf:"varint,3,opt,name=event,proto3,enum=packets.UserEvent" json:"event,omitempty"`
}

func (x *UserPacket) Reset() {
	*x = UserPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPacket) ProtoMessage() {}

func (x *UserPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPacket.ProtoReflect.Descriptor instead.
func (*UserPacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{1}
}

func (x *UserPacket) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *UserPacket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserPacket) GetEvent() UserEvent {
	if x != nil {
		return x.Event
	}
	return UserEvent_JOINED
}

type CommandPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Command   string               `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Args      []string             `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	User      *User                `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Private   bool                 `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
	ArgString string               `protobuf:"bytes,6,opt,name=argString,proto3" json:"argString,omitempty"`
}

func (x *CommandPacket) Reset() {
	*x = CommandPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandPacket) ProtoMessage() {}

func (x *CommandPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandPacket.ProtoReflect.Descriptor instead.
func (*CommandPacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{2}
}

func (x *CommandPacket) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CommandPacket) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandPacket) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *CommandPacket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CommandPacket) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *CommandPacket) GetArgString() string {
	if x != nil {
		return x.ArgString
	}
	return ""
}

type BotPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Recipient *User                `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Private   bool                 `protobuf:"varint,4,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *BotPacket) Reset() {
	*x = BotPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotPacket) ProtoMessage() {}

func (x *BotPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotPacket.ProtoReflect.Descriptor instead.
func (*BotPacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{3}
}

func (x *BotPacket) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *BotPacket) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BotPacket) GetRecipient() *User {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *BotPacket) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type ConfirmationPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotUser *User   `protobuf:"bytes,1,opt,name=botUser,proto3" json:"botUser,omitempty"`
	Users   []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ConfirmationPacket) Reset() {
	*x = ConfirmationPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationPacket) ProtoMessage() {}

func (x *ConfirmationPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationPacket.ProtoReflect.Descriptor instead.
func (*ConfirmationPacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmationPacket) GetBotUser() *User {
	if x != nil {
		return x.BotUser
	}
	return nil
}

func (x *ConfirmationPacket) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type RegistrationPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trigger  string     `protobuf:"bytes,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Commands []*Command `protobuf:"bytes,2,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *RegistrationPacket) Reset() {
	*x = RegistrationPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationPacket) ProtoMessage() {}

func (x *RegistrationPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationPacket.ProtoReflect.Descriptor instead.
func (*RegistrationPacket) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{5}
}

func (x *RegistrationPacket) GetTrigger() string {
	if x != nil {
		return x.Trigger
	}
	return ""
}

func (x *RegistrationPacket) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_packets_proto protoreflect.FileDescriptor

var file_packets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x72, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x72, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x42, 0x6f,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22,
	0x5c, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x62, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x5c, 0x0a,
	0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2a, 0x21, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x4f, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x01, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x66,
	0x39, 0x32, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_packets_proto_rawDescOnce sync.Once
	file_packets_proto_rawDescData = file_packets_proto_rawDesc
)

func file_packets_proto_rawDescGZIP() []byte {
	file_packets_proto_rawDescOnce.Do(func() {
		file_packets_proto_rawDescData = protoimpl.X.CompressGZIP(file_packets_proto_rawDescData)
	})
	return file_packets_proto_rawDescData
}

var file_packets_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_packets_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_packets_proto_goTypes = []interface{}{
	(UserEvent)(0),              // 0: packets.UserEvent
	(*MessagePacket)(nil),       // 1: packets.MessagePacket
	(*UserPacket)(nil),          // 2: packets.UserPacket
	(*CommandPacket)(nil),       // 3: packets.CommandPacket
	(*BotPacket)(nil),           // 4: packets.BotPacket
	(*ConfirmationPacket)(nil),  // 5: packets.ConfirmationPacket
	(*RegistrationPacket)(nil),  // 6: packets.RegistrationPacket
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*User)(nil),                // 8: user.User
	(*Command)(nil),             // 9: command.Command
}
var file_packets_proto_depIdxs = []int32{
	7,  // 0: packets.MessagePacket.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: packets.MessagePacket.user:type_name -> user.User
	7,  // 2: packets.UserPacket.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 3: packets.UserPacket.user:type_name -> user.User
	0,  // 4: packets.UserPacket.event:type_name -> packets.UserEvent
	7,  // 5: packets.CommandPacket.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 6: packets.CommandPacket.user:type_name -> user.User
	7,  // 7: packets.BotPacket.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 8: packets.BotPacket.recipient:type_name -> user.User
	8,  // 9: packets.ConfirmationPacket.botUser:type_name -> user.User
	8,  // 10: packets.ConfirmationPacket.users:type_name -> user.User
	9,  // 11: packets.RegistrationPacket.commands:type_name -> command.Command
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_packets_proto_init() }
func file_packets_proto_init() {
	if File_packets_proto != nil {
		return
	}
	file_user_proto_init()
	file_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_packets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePacket); i {
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
		file_packets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPacket); i {
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
		file_packets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandPacket); i {
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
		file_packets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotPacket); i {
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
		file_packets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationPacket); i {
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
		file_packets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationPacket); i {
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
			RawDescriptor: file_packets_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packets_proto_goTypes,
		DependencyIndexes: file_packets_proto_depIdxs,
		EnumInfos:         file_packets_proto_enumTypes,
		MessageInfos:      file_packets_proto_msgTypes,
	}.Build()
	File_packets_proto = out.File
	file_packets_proto_rawDesc = nil
	file_packets_proto_goTypes = nil
	file_packets_proto_depIdxs = nil
}