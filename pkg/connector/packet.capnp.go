// Code generated by capnpc-go. DO NOT EDIT.

package connector

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Timestamp struct{ capnp.Struct }

// Timestamp_TypeID is the unique identifier for the type Timestamp.
const Timestamp_TypeID = 0xb179ae9d8904b61a

func NewTimestamp(s *capnp.Segment) (Timestamp, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Timestamp{st}, err
}

func NewRootTimestamp(s *capnp.Segment) (Timestamp, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Timestamp{st}, err
}

func ReadRootTimestamp(msg *capnp.Message) (Timestamp, error) {
	root, err := msg.Root()
	return Timestamp{root.Struct()}, err
}

func (s Timestamp) String() string {
	str, _ := text.Marshal(0xb179ae9d8904b61a, s.Struct)
	return str
}

func (s Timestamp) Milliseconds() uint64 {
	return s.Struct.Uint64(0)
}

func (s Timestamp) SetMilliseconds(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Timestamp_List is a list of Timestamp.
type Timestamp_List struct{ capnp.List }

// NewTimestamp creates a new list of Timestamp.
func NewTimestamp_List(s *capnp.Segment, sz int32) (Timestamp_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Timestamp_List{l}, err
}

func (s Timestamp_List) At(i int) Timestamp { return Timestamp{s.List.Struct(i)} }

func (s Timestamp_List) Set(i int, v Timestamp) error { return s.List.SetStruct(i, v.Struct) }

func (s Timestamp_List) String() string {
	str, _ := text.MarshalList(0xb179ae9d8904b61a, s.List)
	return str
}

// Timestamp_Future is a wrapper for a Timestamp promised by a client call.
type Timestamp_Future struct{ *capnp.Future }

func (p Timestamp_Future) Struct() (Timestamp, error) {
	s, err := p.Future.Struct()
	return Timestamp{s}, err
}

type Command struct{ capnp.Struct }

// Command_TypeID is the unique identifier for the type Command.
const Command_TypeID = 0xf42c7dd7018cbb03

func NewCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Command{st}, err
}

func NewRootCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Command{st}, err
}

func ReadRootCommand(msg *capnp.Message) (Command, error) {
	root, err := msg.Root()
	return Command{root.Struct()}, err
}

func (s Command) String() string {
	str, _ := text.Marshal(0xf42c7dd7018cbb03, s.Struct)
	return str
}

func (s Command) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Command) HasName() bool {
	return s.Struct.HasPtr(0)
}

func (s Command) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Command) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Command) Aliases() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Command) HasAliases() bool {
	return s.Struct.HasPtr(1)
}

func (s Command) SetAliases(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAliases sets the aliases field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Command) NewAliases(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Command) Usage() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Command) HasUsage() bool {
	return s.Struct.HasPtr(2)
}

func (s Command) UsageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Command) SetUsage(v string) error {
	return s.Struct.SetText(2, v)
}

// Command_List is a list of Command.
type Command_List struct{ capnp.List }

// NewCommand creates a new list of Command.
func NewCommand_List(s *capnp.Segment, sz int32) (Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Command_List{l}, err
}

func (s Command_List) At(i int) Command { return Command{s.List.Struct(i)} }

func (s Command_List) Set(i int, v Command) error { return s.List.SetStruct(i, v.Struct) }

func (s Command_List) String() string {
	str, _ := text.MarshalList(0xf42c7dd7018cbb03, s.List)
	return str
}

// Command_Future is a wrapper for a Command promised by a client call.
type Command_Future struct{ *capnp.Future }

func (p Command_Future) Struct() (Command, error) {
	s, err := p.Future.Struct()
	return Command{s}, err
}

type UserRole uint16

// UserRole_TypeID is the unique identifier for the type UserRole.
const UserRole_TypeID = 0x971dd3472a97b1b5

// Values of UserRole.
const (
	UserRole_regular UserRole = 0
	UserRole_mod     UserRole = 1
	UserRole_admin   UserRole = 2
)

// String returns the enum's constant name.
func (c UserRole) String() string {
	switch c {
	case UserRole_regular:
		return "regular"
	case UserRole_mod:
		return "mod"
	case UserRole_admin:
		return "admin"

	default:
		return ""
	}
}

// UserRoleFromString returns the enum value with a name,
// or the zero value if there's no such value.
func UserRoleFromString(c string) UserRole {
	switch c {
	case "regular":
		return UserRole_regular
	case "mod":
		return UserRole_mod
	case "admin":
		return UserRole_admin

	default:
		return 0
	}
}

type UserRole_List struct{ capnp.List }

func NewUserRole_List(s *capnp.Segment, sz int32) (UserRole_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return UserRole_List{l.List}, err
}

func (l UserRole_List) At(i int) UserRole {
	ul := capnp.UInt16List{List: l.List}
	return UserRole(ul.At(i))
}

func (l UserRole_List) Set(i int, v UserRole) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type User struct{ capnp.Struct }

// User_TypeID is the unique identifier for the type User.
const User_TypeID = 0xa1d57c8c488d1cc0

func NewUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return User{st}, err
}

func NewRootUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return User{st}, err
}

func ReadRootUser(msg *capnp.Message) (User, error) {
	root, err := msg.Root()
	return User{root.Struct()}, err
}

func (s User) String() string {
	str, _ := text.Marshal(0xa1d57c8c488d1cc0, s.Struct)
	return str
}

func (s User) Nick() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s User) HasNick() bool {
	return s.Struct.HasPtr(0)
}

func (s User) NickBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s User) SetNick(v string) error {
	return s.Struct.SetText(0, v)
}

func (s User) Id() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s User) HasId() bool {
	return s.Struct.HasPtr(1)
}

func (s User) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s User) SetId(v string) error {
	return s.Struct.SetText(1, v)
}

func (s User) Role() UserRole {
	return UserRole(s.Struct.Uint16(0))
}

func (s User) SetRole(v UserRole) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s User) JoinedAt() (Timestamp, error) {
	p, err := s.Struct.Ptr(2)
	return Timestamp{Struct: p.Struct()}, err
}

func (s User) HasJoinedAt() bool {
	return s.Struct.HasPtr(2)
}

func (s User) SetJoinedAt(v Timestamp) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewJoinedAt sets the joinedAt field to a newly
// allocated Timestamp struct, preferring placement in s's segment.
func (s User) NewJoinedAt() (Timestamp, error) {
	ss, err := NewTimestamp(s.Struct.Segment())
	if err != nil {
		return Timestamp{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// User_List is a list of User.
type User_List struct{ capnp.List }

// NewUser creates a new list of User.
func NewUser_List(s *capnp.Segment, sz int32) (User_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return User_List{l}, err
}

func (s User_List) At(i int) User { return User{s.List.Struct(i)} }

func (s User_List) Set(i int, v User) error { return s.List.SetStruct(i, v.Struct) }

func (s User_List) String() string {
	str, _ := text.MarshalList(0xa1d57c8c488d1cc0, s.List)
	return str
}

// User_Future is a wrapper for a User promised by a client call.
type User_Future struct{ *capnp.Future }

func (p User_Future) Struct() (User, error) {
	s, err := p.Future.Struct()
	return User{s}, err
}

func (p User_Future) JoinedAt() Timestamp_Future {
	return Timestamp_Future{Future: p.Future.Field(2, nil)}
}

type IncomingMessagePacket struct{ capnp.Struct }

// IncomingMessagePacket_TypeID is the unique identifier for the type IncomingMessagePacket.
const IncomingMessagePacket_TypeID = 0xb0342843020dafb0

func NewIncomingMessagePacket(s *capnp.Segment) (IncomingMessagePacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return IncomingMessagePacket{st}, err
}

func NewRootIncomingMessagePacket(s *capnp.Segment) (IncomingMessagePacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return IncomingMessagePacket{st}, err
}

func ReadRootIncomingMessagePacket(msg *capnp.Message) (IncomingMessagePacket, error) {
	root, err := msg.Root()
	return IncomingMessagePacket{root.Struct()}, err
}

func (s IncomingMessagePacket) String() string {
	str, _ := text.Marshal(0xb0342843020dafb0, s.Struct)
	return str
}

func (s IncomingMessagePacket) Timestamp() (Timestamp, error) {
	p, err := s.Struct.Ptr(0)
	return Timestamp{Struct: p.Struct()}, err
}

func (s IncomingMessagePacket) HasTimestamp() bool {
	return s.Struct.HasPtr(0)
}

func (s IncomingMessagePacket) SetTimestamp(v Timestamp) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTimestamp sets the timestamp field to a newly
// allocated Timestamp struct, preferring placement in s's segment.
func (s IncomingMessagePacket) NewTimestamp() (Timestamp, error) {
	ss, err := NewTimestamp(s.Struct.Segment())
	if err != nil {
		return Timestamp{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s IncomingMessagePacket) Message() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s IncomingMessagePacket) HasMessage() bool {
	return s.Struct.HasPtr(1)
}

func (s IncomingMessagePacket) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s IncomingMessagePacket) SetMessage(v string) error {
	return s.Struct.SetText(1, v)
}

func (s IncomingMessagePacket) Sender() (User, error) {
	p, err := s.Struct.Ptr(2)
	return User{Struct: p.Struct()}, err
}

func (s IncomingMessagePacket) HasSender() bool {
	return s.Struct.HasPtr(2)
}

func (s IncomingMessagePacket) SetSender(v User) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewSender sets the sender field to a newly
// allocated User struct, preferring placement in s's segment.
func (s IncomingMessagePacket) NewSender() (User, error) {
	ss, err := NewUser(s.Struct.Segment())
	if err != nil {
		return User{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s IncomingMessagePacket) Private() bool {
	return s.Struct.Bit(0)
}

func (s IncomingMessagePacket) SetPrivate(v bool) {
	s.Struct.SetBit(0, v)
}

func (s IncomingMessagePacket) Recipients() (User_List, error) {
	p, err := s.Struct.Ptr(3)
	return User_List{List: p.List()}, err
}

func (s IncomingMessagePacket) HasRecipients() bool {
	return s.Struct.HasPtr(3)
}

func (s IncomingMessagePacket) SetRecipients(v User_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewRecipients sets the recipients field to a newly
// allocated User_List, preferring placement in s's segment.
func (s IncomingMessagePacket) NewRecipients(n int32) (User_List, error) {
	l, err := NewUser_List(s.Struct.Segment(), n)
	if err != nil {
		return User_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s IncomingMessagePacket) MentionsConnectorUser() bool {
	return s.Struct.Bit(1)
}

func (s IncomingMessagePacket) SetMentionsConnectorUser(v bool) {
	s.Struct.SetBit(1, v)
}

func (s IncomingMessagePacket) Incoming() bool {
	return s.Struct.Bit(2)
}

func (s IncomingMessagePacket) SetIncoming(v bool) {
	s.Struct.SetBit(2, v)
}

// IncomingMessagePacket_List is a list of IncomingMessagePacket.
type IncomingMessagePacket_List struct{ capnp.List }

// NewIncomingMessagePacket creates a new list of IncomingMessagePacket.
func NewIncomingMessagePacket_List(s *capnp.Segment, sz int32) (IncomingMessagePacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return IncomingMessagePacket_List{l}, err
}

func (s IncomingMessagePacket_List) At(i int) IncomingMessagePacket {
	return IncomingMessagePacket{s.List.Struct(i)}
}

func (s IncomingMessagePacket_List) Set(i int, v IncomingMessagePacket) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s IncomingMessagePacket_List) String() string {
	str, _ := text.MarshalList(0xb0342843020dafb0, s.List)
	return str
}

// IncomingMessagePacket_Future is a wrapper for a IncomingMessagePacket promised by a client call.
type IncomingMessagePacket_Future struct{ *capnp.Future }

func (p IncomingMessagePacket_Future) Struct() (IncomingMessagePacket, error) {
	s, err := p.Future.Struct()
	return IncomingMessagePacket{s}, err
}

func (p IncomingMessagePacket_Future) Timestamp() Timestamp_Future {
	return Timestamp_Future{Future: p.Future.Field(0, nil)}
}

func (p IncomingMessagePacket_Future) Sender() User_Future {
	return User_Future{Future: p.Future.Field(2, nil)}
}

type UserEvent uint16

// UserEvent_TypeID is the unique identifier for the type UserEvent.
const UserEvent_TypeID = 0xf145976cbe38c029

// Values of UserEvent.
const (
	UserEvent_joined UserEvent = 0
	UserEvent_left   UserEvent = 1
)

// String returns the enum's constant name.
func (c UserEvent) String() string {
	switch c {
	case UserEvent_joined:
		return "joined"
	case UserEvent_left:
		return "left"

	default:
		return ""
	}
}

// UserEventFromString returns the enum value with a name,
// or the zero value if there's no such value.
func UserEventFromString(c string) UserEvent {
	switch c {
	case "joined":
		return UserEvent_joined
	case "left":
		return UserEvent_left

	default:
		return 0
	}
}

type UserEvent_List struct{ capnp.List }

func NewUserEvent_List(s *capnp.Segment, sz int32) (UserEvent_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return UserEvent_List{l.List}, err
}

func (l UserEvent_List) At(i int) UserEvent {
	ul := capnp.UInt16List{List: l.List}
	return UserEvent(ul.At(i))
}

func (l UserEvent_List) Set(i int, v UserEvent) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type UserPacket struct{ capnp.Struct }

// UserPacket_TypeID is the unique identifier for the type UserPacket.
const UserPacket_TypeID = 0xd6ef75d402487ecc

func NewUserPacket(s *capnp.Segment) (UserPacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return UserPacket{st}, err
}

func NewRootUserPacket(s *capnp.Segment) (UserPacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return UserPacket{st}, err
}

func ReadRootUserPacket(msg *capnp.Message) (UserPacket, error) {
	root, err := msg.Root()
	return UserPacket{root.Struct()}, err
}

func (s UserPacket) String() string {
	str, _ := text.Marshal(0xd6ef75d402487ecc, s.Struct)
	return str
}

func (s UserPacket) Timestamp() (Timestamp, error) {
	p, err := s.Struct.Ptr(0)
	return Timestamp{Struct: p.Struct()}, err
}

func (s UserPacket) HasTimestamp() bool {
	return s.Struct.HasPtr(0)
}

func (s UserPacket) SetTimestamp(v Timestamp) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTimestamp sets the timestamp field to a newly
// allocated Timestamp struct, preferring placement in s's segment.
func (s UserPacket) NewTimestamp() (Timestamp, error) {
	ss, err := NewTimestamp(s.Struct.Segment())
	if err != nil {
		return Timestamp{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UserPacket) Event() UserEvent {
	return UserEvent(s.Struct.Uint16(0))
}

func (s UserPacket) SetEvent(v UserEvent) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s UserPacket) User() (User, error) {
	p, err := s.Struct.Ptr(1)
	return User{Struct: p.Struct()}, err
}

func (s UserPacket) HasUser() bool {
	return s.Struct.HasPtr(1)
}

func (s UserPacket) SetUser(v User) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewUser sets the user field to a newly
// allocated User struct, preferring placement in s's segment.
func (s UserPacket) NewUser() (User, error) {
	ss, err := NewUser(s.Struct.Segment())
	if err != nil {
		return User{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// UserPacket_List is a list of UserPacket.
type UserPacket_List struct{ capnp.List }

// NewUserPacket creates a new list of UserPacket.
func NewUserPacket_List(s *capnp.Segment, sz int32) (UserPacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return UserPacket_List{l}, err
}

func (s UserPacket_List) At(i int) UserPacket { return UserPacket{s.List.Struct(i)} }

func (s UserPacket_List) Set(i int, v UserPacket) error { return s.List.SetStruct(i, v.Struct) }

func (s UserPacket_List) String() string {
	str, _ := text.MarshalList(0xd6ef75d402487ecc, s.List)
	return str
}

// UserPacket_Future is a wrapper for a UserPacket promised by a client call.
type UserPacket_Future struct{ *capnp.Future }

func (p UserPacket_Future) Struct() (UserPacket, error) {
	s, err := p.Future.Struct()
	return UserPacket{s}, err
}

func (p UserPacket_Future) Timestamp() Timestamp_Future {
	return Timestamp_Future{Future: p.Future.Field(0, nil)}
}

func (p UserPacket_Future) User() User_Future {
	return User_Future{Future: p.Future.Field(1, nil)}
}

type CommandPacket struct{ capnp.Struct }

// CommandPacket_TypeID is the unique identifier for the type CommandPacket.
const CommandPacket_TypeID = 0xd55a13aeeb1a986d

func NewCommandPacket(s *capnp.Segment) (CommandPacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return CommandPacket{st}, err
}

func NewRootCommandPacket(s *capnp.Segment) (CommandPacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return CommandPacket{st}, err
}

func ReadRootCommandPacket(msg *capnp.Message) (CommandPacket, error) {
	root, err := msg.Root()
	return CommandPacket{root.Struct()}, err
}

func (s CommandPacket) String() string {
	str, _ := text.Marshal(0xd55a13aeeb1a986d, s.Struct)
	return str
}

func (s CommandPacket) Timestamp() (Timestamp, error) {
	p, err := s.Struct.Ptr(0)
	return Timestamp{Struct: p.Struct()}, err
}

func (s CommandPacket) HasTimestamp() bool {
	return s.Struct.HasPtr(0)
}

func (s CommandPacket) SetTimestamp(v Timestamp) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTimestamp sets the timestamp field to a newly
// allocated Timestamp struct, preferring placement in s's segment.
func (s CommandPacket) NewTimestamp() (Timestamp, error) {
	ss, err := NewTimestamp(s.Struct.Segment())
	if err != nil {
		return Timestamp{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CommandPacket) Command() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s CommandPacket) HasCommand() bool {
	return s.Struct.HasPtr(1)
}

func (s CommandPacket) CommandBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s CommandPacket) SetCommand(v string) error {
	return s.Struct.SetText(1, v)
}

func (s CommandPacket) Sender() (User, error) {
	p, err := s.Struct.Ptr(2)
	return User{Struct: p.Struct()}, err
}

func (s CommandPacket) HasSender() bool {
	return s.Struct.HasPtr(2)
}

func (s CommandPacket) SetSender(v User) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewSender sets the sender field to a newly
// allocated User struct, preferring placement in s's segment.
func (s CommandPacket) NewSender() (User, error) {
	ss, err := NewUser(s.Struct.Segment())
	if err != nil {
		return User{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s CommandPacket) Private() bool {
	return s.Struct.Bit(0)
}

func (s CommandPacket) SetPrivate(v bool) {
	s.Struct.SetBit(0, v)
}

func (s CommandPacket) Args() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.TextList{List: p.List()}, err
}

func (s CommandPacket) HasArgs() bool {
	return s.Struct.HasPtr(3)
}

func (s CommandPacket) SetArgs(v capnp.TextList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewArgs sets the args field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s CommandPacket) NewArgs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s CommandPacket) ArgString() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s CommandPacket) HasArgString() bool {
	return s.Struct.HasPtr(4)
}

func (s CommandPacket) ArgStringBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s CommandPacket) SetArgString(v string) error {
	return s.Struct.SetText(4, v)
}

// CommandPacket_List is a list of CommandPacket.
type CommandPacket_List struct{ capnp.List }

// NewCommandPacket creates a new list of CommandPacket.
func NewCommandPacket_List(s *capnp.Segment, sz int32) (CommandPacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return CommandPacket_List{l}, err
}

func (s CommandPacket_List) At(i int) CommandPacket { return CommandPacket{s.List.Struct(i)} }

func (s CommandPacket_List) Set(i int, v CommandPacket) error { return s.List.SetStruct(i, v.Struct) }

func (s CommandPacket_List) String() string {
	str, _ := text.MarshalList(0xd55a13aeeb1a986d, s.List)
	return str
}

// CommandPacket_Future is a wrapper for a CommandPacket promised by a client call.
type CommandPacket_Future struct{ *capnp.Future }

func (p CommandPacket_Future) Struct() (CommandPacket, error) {
	s, err := p.Future.Struct()
	return CommandPacket{s}, err
}

func (p CommandPacket_Future) Timestamp() Timestamp_Future {
	return Timestamp_Future{Future: p.Future.Field(0, nil)}
}

func (p CommandPacket_Future) Sender() User_Future {
	return User_Future{Future: p.Future.Field(2, nil)}
}

type OutgoingMessagePacket struct{ capnp.Struct }

// OutgoingMessagePacket_TypeID is the unique identifier for the type OutgoingMessagePacket.
const OutgoingMessagePacket_TypeID = 0xf1a684870430173a

func NewOutgoingMessagePacket(s *capnp.Segment) (OutgoingMessagePacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return OutgoingMessagePacket{st}, err
}

func NewRootOutgoingMessagePacket(s *capnp.Segment) (OutgoingMessagePacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return OutgoingMessagePacket{st}, err
}

func ReadRootOutgoingMessagePacket(msg *capnp.Message) (OutgoingMessagePacket, error) {
	root, err := msg.Root()
	return OutgoingMessagePacket{root.Struct()}, err
}

func (s OutgoingMessagePacket) String() string {
	str, _ := text.Marshal(0xf1a684870430173a, s.Struct)
	return str
}

func (s OutgoingMessagePacket) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s OutgoingMessagePacket) HasMessage() bool {
	return s.Struct.HasPtr(0)
}

func (s OutgoingMessagePacket) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s OutgoingMessagePacket) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s OutgoingMessagePacket) Recipient() (User, error) {
	p, err := s.Struct.Ptr(1)
	return User{Struct: p.Struct()}, err
}

func (s OutgoingMessagePacket) HasRecipient() bool {
	return s.Struct.HasPtr(1)
}

func (s OutgoingMessagePacket) SetRecipient(v User) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewRecipient sets the recipient field to a newly
// allocated User struct, preferring placement in s's segment.
func (s OutgoingMessagePacket) NewRecipient() (User, error) {
	ss, err := NewUser(s.Struct.Segment())
	if err != nil {
		return User{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s OutgoingMessagePacket) Private() bool {
	return s.Struct.Bit(0)
}

func (s OutgoingMessagePacket) SetPrivate(v bool) {
	s.Struct.SetBit(0, v)
}

// OutgoingMessagePacket_List is a list of OutgoingMessagePacket.
type OutgoingMessagePacket_List struct{ capnp.List }

// NewOutgoingMessagePacket creates a new list of OutgoingMessagePacket.
func NewOutgoingMessagePacket_List(s *capnp.Segment, sz int32) (OutgoingMessagePacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return OutgoingMessagePacket_List{l}, err
}

func (s OutgoingMessagePacket_List) At(i int) OutgoingMessagePacket {
	return OutgoingMessagePacket{s.List.Struct(i)}
}

func (s OutgoingMessagePacket_List) Set(i int, v OutgoingMessagePacket) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s OutgoingMessagePacket_List) String() string {
	str, _ := text.MarshalList(0xf1a684870430173a, s.List)
	return str
}

// OutgoingMessagePacket_Future is a wrapper for a OutgoingMessagePacket promised by a client call.
type OutgoingMessagePacket_Future struct{ *capnp.Future }

func (p OutgoingMessagePacket_Future) Struct() (OutgoingMessagePacket, error) {
	s, err := p.Future.Struct()
	return OutgoingMessagePacket{s}, err
}

func (p OutgoingMessagePacket_Future) Recipient() User_Future {
	return User_Future{Future: p.Future.Field(1, nil)}
}

type RegistrationPacket struct{ capnp.Struct }

// RegistrationPacket_TypeID is the unique identifier for the type RegistrationPacket.
const RegistrationPacket_TypeID = 0x93e7595095e6d2c8

func NewRegistrationPacket(s *capnp.Segment) (RegistrationPacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RegistrationPacket{st}, err
}

func NewRootRegistrationPacket(s *capnp.Segment) (RegistrationPacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RegistrationPacket{st}, err
}

func ReadRootRegistrationPacket(msg *capnp.Message) (RegistrationPacket, error) {
	root, err := msg.Root()
	return RegistrationPacket{root.Struct()}, err
}

func (s RegistrationPacket) String() string {
	str, _ := text.Marshal(0x93e7595095e6d2c8, s.Struct)
	return str
}

func (s RegistrationPacket) Commands() (Command_List, error) {
	p, err := s.Struct.Ptr(0)
	return Command_List{List: p.List()}, err
}

func (s RegistrationPacket) HasCommands() bool {
	return s.Struct.HasPtr(0)
}

func (s RegistrationPacket) SetCommands(v Command_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewCommands sets the commands field to a newly
// allocated Command_List, preferring placement in s's segment.
func (s RegistrationPacket) NewCommands(n int32) (Command_List, error) {
	l, err := NewCommand_List(s.Struct.Segment(), n)
	if err != nil {
		return Command_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// RegistrationPacket_List is a list of RegistrationPacket.
type RegistrationPacket_List struct{ capnp.List }

// NewRegistrationPacket creates a new list of RegistrationPacket.
func NewRegistrationPacket_List(s *capnp.Segment, sz int32) (RegistrationPacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return RegistrationPacket_List{l}, err
}

func (s RegistrationPacket_List) At(i int) RegistrationPacket {
	return RegistrationPacket{s.List.Struct(i)}
}

func (s RegistrationPacket_List) Set(i int, v RegistrationPacket) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s RegistrationPacket_List) String() string {
	str, _ := text.MarshalList(0x93e7595095e6d2c8, s.List)
	return str
}

// RegistrationPacket_Future is a wrapper for a RegistrationPacket promised by a client call.
type RegistrationPacket_Future struct{ *capnp.Future }

func (p RegistrationPacket_Future) Struct() (RegistrationPacket, error) {
	s, err := p.Future.Struct()
	return RegistrationPacket{s}, err
}

type ConfirmationPacket struct{ capnp.Struct }

// ConfirmationPacket_TypeID is the unique identifier for the type ConfirmationPacket.
const ConfirmationPacket_TypeID = 0x89b5686804357bc6

func NewConfirmationPacket(s *capnp.Segment) (ConfirmationPacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ConfirmationPacket{st}, err
}

func NewRootConfirmationPacket(s *capnp.Segment) (ConfirmationPacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ConfirmationPacket{st}, err
}

func ReadRootConfirmationPacket(msg *capnp.Message) (ConfirmationPacket, error) {
	root, err := msg.Root()
	return ConfirmationPacket{root.Struct()}, err
}

func (s ConfirmationPacket) String() string {
	str, _ := text.Marshal(0x89b5686804357bc6, s.Struct)
	return str
}

func (s ConfirmationPacket) BotUser() (User, error) {
	p, err := s.Struct.Ptr(0)
	return User{Struct: p.Struct()}, err
}

func (s ConfirmationPacket) HasBotUser() bool {
	return s.Struct.HasPtr(0)
}

func (s ConfirmationPacket) SetBotUser(v User) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBotUser sets the botUser field to a newly
// allocated User struct, preferring placement in s's segment.
func (s ConfirmationPacket) NewBotUser() (User, error) {
	ss, err := NewUser(s.Struct.Segment())
	if err != nil {
		return User{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s ConfirmationPacket) Trigger() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s ConfirmationPacket) HasTrigger() bool {
	return s.Struct.HasPtr(1)
}

func (s ConfirmationPacket) TriggerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s ConfirmationPacket) SetTrigger(v string) error {
	return s.Struct.SetText(1, v)
}

func (s ConfirmationPacket) Users() (User_List, error) {
	p, err := s.Struct.Ptr(2)
	return User_List{List: p.List()}, err
}

func (s ConfirmationPacket) HasUsers() bool {
	return s.Struct.HasPtr(2)
}

func (s ConfirmationPacket) SetUsers(v User_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewUsers sets the users field to a newly
// allocated User_List, preferring placement in s's segment.
func (s ConfirmationPacket) NewUsers(n int32) (User_List, error) {
	l, err := NewUser_List(s.Struct.Segment(), n)
	if err != nil {
		return User_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// ConfirmationPacket_List is a list of ConfirmationPacket.
type ConfirmationPacket_List struct{ capnp.List }

// NewConfirmationPacket creates a new list of ConfirmationPacket.
func NewConfirmationPacket_List(s *capnp.Segment, sz int32) (ConfirmationPacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return ConfirmationPacket_List{l}, err
}

func (s ConfirmationPacket_List) At(i int) ConfirmationPacket {
	return ConfirmationPacket{s.List.Struct(i)}
}

func (s ConfirmationPacket_List) Set(i int, v ConfirmationPacket) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s ConfirmationPacket_List) String() string {
	str, _ := text.MarshalList(0x89b5686804357bc6, s.List)
	return str
}

// ConfirmationPacket_Future is a wrapper for a ConfirmationPacket promised by a client call.
type ConfirmationPacket_Future struct{ *capnp.Future }

func (p ConfirmationPacket_Future) Struct() (ConfirmationPacket, error) {
	s, err := p.Future.Struct()
	return ConfirmationPacket{s}, err
}

func (p ConfirmationPacket_Future) BotUser() User_Future {
	return User_Future{Future: p.Future.Field(0, nil)}
}

const schema_824b0dcb3e553a2a = "x\xda\x94Uo\x88\x14e\x18\x7f~\xef;\xbb\xee\xc2" +
	"\xad{\xc3l$\x87\xb2\xd0';*\xd4K\x8c\xfb\xd0" +
	"\x99\x97\xa4\xa5y\xef\xa6dR\xe0\xb4\xfb\xba\x8e\xee\xcc" +
	",3\xb3R\x94\x1dE\x11\x82\x11\x14\xa1BB\x82\x85" +
	"E\xf9\x8f.\xca2\xee\xa2\xe2$\xfd\xd0\x9f+\xe8\xdb" +
	"\x81\xe5\x87:\xea\xa0\xa0o\x13\xef\xec\x9f\xd9\x9b[\xbc" +
	"\xfa\xf6\xf2\xeeo\x9f\xf7\xf9\xfdy\x9eY\x93\xe3\x1b\xd9" +
	"\xda\xd4\x864\x91\xd8\x92J\x87_?\xb3^\xdb\xbf\x7f" +
	"\xe2\x08\xe99\x84\x83\xc3\xbb\xee\xfd&\xf7\xd0\x0b\x94\xe2" +
	"\xcb\x88\x86\xc0\x19\x8c\x9c:\x1aY~\x8e\x10N\x7f\xf7" +
	"\xeb\x1bc\x8f\xddx=\x01\x86\x02_P\xe0\xcb\x11\xf8" +
	"\x13>B\x08'.\x1c\x1b|\xe0\xfbU\xc7H\xcf\xb1" +
	"\x18L0\xe6\xf8E\xe3\xef\x088\xcfw\x10\xc2\xc9\x95" +
	"\xafl9\xfa\xec\xcc)\x129${0R\xdakF" +
	"N\x8bZ\xd0n\x10\xc2\xf3\xe7rlt\xf5\xdd\xe7\x93" +
	"`\x05\x19\x9a\xd5\x06`\xccG\xc79\xedU\x10\xc2\x81" +
	"\x8f\xb4#'\xcf>}!\x01\x8f\x0a\x9eH_2N" +
	"\xa5\xd5\xe9dZ5l\x1f\x1f\xf8\xed\xac\xb1g&Y" +
	":\xa5 S\xe9\x9f\x8dk\x0a<t%\xfd\xa8\xaa|" +
	"\xf5\xb9-\xec\x87\xc6\x1f?&\xd1L\xa1\xef\xcc|a" +
	"\xac\xcf\xa8\xd3\xda\x8c\x12\xee\xf6\xc9{>\xaf\x1d\xdb<" +
	"\xbfH\x8bk\x99K\xc6L\x04\xfc6\xb3\x81\x10\x0e\xdf" +
	"\xbaF{\xf9\xc5w\xe6{U\x1d\x9a\xc9\x0c\xc0\xf8%" +
	"B\xcfFe\xf9\xa7G\xf1\xd3\xe1;\xfe\xeaa\x9eq" +
	"8\xfb\x9e\xf1RV\x9d\x9e\xcf\x9e\xa3\x87\xc3\xbaY>" +
	"(\x83\xbb\xca\xdc\xac;\xf5\xe1Q\xd7\xd9gy\xb6\x19" +
	"X\xae3\x16\xfdBc\x80\xe8\xe3\x1a\x91\x06\"}\xf3" +
	"&\"\xb1\x91Clc\xd0\x81\x02\xd4\xe5Vuy?" +
	"\x87\x18c\xd0\x19+\x80\x11\xe9\xdb\xd7\xa9(q\x88\x0a" +
	"\xc3\xf8\x93n\xb0\xcb\x97\x1e\xfacW\x09\xe8'\x8c\x07" +
	"\x9eU\xadJ\x0f}\xc4\xd0G(6|\xe9\xf9XN" +
	"\x18\xe3H\xc0\x97\x13\x12\xed\x96d\xd5\xf2\x03/\xd9\xae" +
	"\xd6i7\xf7 \x91\xe8\xe3\x10\xab\x19\xc2\xb2k\xdb\xa6" +
	"S\xf1\x89(~\xa0\xa3U\xe2\x01D\x0f\xec\xf2e\xd1" +
	"+\xb95\x19\xa9\x10\xf1Z\xb5IA\xf5[n#\x02" +
	"\xd3s\xeb\x88\xc6=Ym\xd4Lo\x99\xedV\x8af" +
	"\xc5\xb6\x9c\xc5U\xe0\xa9\x0a\xfd\x9d\xc6\xccA\"\xf18" +
	"\x87\xd8\xdf\xa5\xa3\x1c \x12{9D\x8d\x01-\x19-" +
	"\x05\xacp\x88:\x83\xceY\x01\x9cH\xb7\x15\xad\x1a\x87" +
	"x\x8a!\xefX\xe5\x83m\xf9\xb8Ui\x1f\xf3\x9e[" +
	"\x93\xc8\xc7\xf3F@\x9e\x10\x1ep-GV\xee\x0b\x94" +
	"\x0a\xfd\xf1 4\xedH\xe8\xbb\xd5)\xbb\xb6\xe5T\xb7" +
	"K\xdf7\xabRI\xcce\xa0\x98\xac\xec0\x99(\x11" +
	"\x89\x0f9\xc4d\x17\x93\xcb*\x11\x1fs\x88/\xbb\x12" +
	"15L$>\xe3\x10\xd3\x0c\xe0M&_)\xe0$" +
	"\x87\xb8\xca\xa0k\xbc\x00\x8dH\xbf\xb2\x87HLs\x88" +
	"\xeb\x0cz\x0a\x05\xa4\x88\xf4\xd9\x8bD\xe2:\x87\xf8\x93" +
	"AO\xb3\x02\xd2D\xfa\x9c\x12\xe2w\x0e\xf1\x0fC\x18" +
	"X\xb6\xf4\x03\xd3&\xd4\x173\x1b\xb7\x9b\x1c\xda\xf2\x8c" +
	"\xf8\xd2\xa9\xf4Ld\xdd\xb3\x0e\x99\x81\x04\x88AM\xb4" +
	"'\xcbV\xdd\x92\x0e\xf1\xe0&\xc1\xb4\xa5\xa3R\xe8c" +
	"\xd4u\x1cY\x0e\xdc\xa2\x17%\xbe]\xc5jI\xa9d" +
	"o\xdf-L\xc9N\xcb\x1e\x89\xfa\xaf'2|\xa0\x95" +
	"\xe1\x15\x0c\xa1m\xd5j\x96/\xcb\x94w\x9d\x8a\x8f," +
	"1d\xbb*\xb1\xd6\x14GA\xef\xb2kE\xa7\xda\x09" +
	"e\xd7q\x0eq\xba\xcb\xaeS\xca\x8579\xc4\x99." +
	"\xbb\xdeVv\xbd\xc5!\xde\x8f\xedzW\x01Os\x88" +
	"\xf3]v}\xa0\"z\xa6evJk\xda5U\x8a" +
	"\x8d]\xc2\x99\xd6\\\xfeog\xf2\xa6W\xed8\xa2\xfe" +
	"\xac|0\xbd\xea#\x81g9\x84j\xbb`B\x1de" +
	"K\xef\xddV\x8a\xd7X[\x99\xce\x16\xdb\xa9\x94AS" +
	"\x19\xa1\x08o\xe3\x10\xbb\x97\xe2V\x94\x87\xa4\x13 \x1f" +
	"o\xfa\xe6\x14\xe6\x1b=\xb7\xe1\xe2\xbd1\xe2mV\x15" +
	"T\xa3\x99\xe8m}8Z?\xd9A\xa2\x91\xe6(\xe7" +
	"kr_\x90\x18\xdc\x1d\x8d\xa0\xea\xf6\x1c\xdc\xa5Vy" +
	")\xe6\xdb^AbSL79F\xf1t \xf8\x0f" +
	"\xb3\xb4\x90\xdf\xa8k\xe7\xa3\xa0.\xeck\xf0f\x9f\x98" +
	"\xbd]\x09}B\x99\xb3\xbb\xf9\x89\xc9;\xa6\xdd\xe9j" +
	"\xdc\xacY\xa6/\x93\xe1(6\xba[\xff7\x00\x00\xff" +
	"\xff\x10R\x1dX"

func init() {
	schemas.Register(schema_824b0dcb3e553a2a,
		0x89b5686804357bc6,
		0x93e7595095e6d2c8,
		0x971dd3472a97b1b5,
		0xa1d57c8c488d1cc0,
		0xb0342843020dafb0,
		0xb179ae9d8904b61a,
		0xd55a13aeeb1a986d,
		0xd6ef75d402487ecc,
		0xf145976cbe38c029,
		0xf1a684870430173a,
		0xf42c7dd7018cbb03)
}
