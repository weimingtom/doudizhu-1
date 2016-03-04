// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	S2SSystem
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type S2SSystem struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2SSystem) Reset()         { *m = S2SSystem{} }
func (m *S2SSystem) String() string { return proto.CompactTextString(m) }
func (*S2SSystem) ProtoMessage()    {}

type S2SSystem_GetType struct {
	Pid              *int32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2SSystem_GetType) Reset()         { *m = S2SSystem_GetType{} }
func (m *S2SSystem_GetType) String() string { return proto.CompactTextString(m) }
func (*S2SSystem_GetType) ProtoMessage()    {}

func (m *S2SSystem_GetType) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type S2SSystem_LoginInfo struct {
	Pid              *int32  `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Playername       *string `protobuf:"bytes,2,opt,name=playername" json:"playername,omitempty"`
	Passworld        *string `protobuf:"bytes,3,opt,name=passworld" json:"passworld,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2SSystem_LoginInfo) Reset()         { *m = S2SSystem_LoginInfo{} }
func (m *S2SSystem_LoginInfo) String() string { return proto.CompactTextString(m) }
func (*S2SSystem_LoginInfo) ProtoMessage()    {}

func (m *S2SSystem_LoginInfo) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *S2SSystem_LoginInfo) GetPlayername() string {
	if m != nil && m.Playername != nil {
		return *m.Playername
	}
	return ""
}

func (m *S2SSystem_LoginInfo) GetPassworld() string {
	if m != nil && m.Passworld != nil {
		return *m.Passworld
	}
	return ""
}

type S2SSystem_RegisterPlayer struct {
	Pid              *int32  `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Playername       *string `protobuf:"bytes,2,opt,name=playername" json:"playername,omitempty"`
	Passworld        *string `protobuf:"bytes,3,opt,name=passworld" json:"passworld,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2SSystem_RegisterPlayer) Reset()         { *m = S2SSystem_RegisterPlayer{} }
func (m *S2SSystem_RegisterPlayer) String() string { return proto.CompactTextString(m) }
func (*S2SSystem_RegisterPlayer) ProtoMessage()    {}

func (m *S2SSystem_RegisterPlayer) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *S2SSystem_RegisterPlayer) GetPlayername() string {
	if m != nil && m.Playername != nil {
		return *m.Playername
	}
	return ""
}

func (m *S2SSystem_RegisterPlayer) GetPassworld() string {
	if m != nil && m.Passworld != nil {
		return *m.Passworld
	}
	return ""
}

type S2SSystem_LoginResult struct {
	Pid              *int32  `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Result           *int32  `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	Balanceserver    *string `protobuf:"bytes,3,opt,name=balanceserver" json:"balanceserver,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2SSystem_LoginResult) Reset()         { *m = S2SSystem_LoginResult{} }
func (m *S2SSystem_LoginResult) String() string { return proto.CompactTextString(m) }
func (*S2SSystem_LoginResult) ProtoMessage()    {}

func (m *S2SSystem_LoginResult) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *S2SSystem_LoginResult) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *S2SSystem_LoginResult) GetBalanceserver() string {
	if m != nil && m.Balanceserver != nil {
		return *m.Balanceserver
	}
	return ""
}

type S2SSystem_RegisterResult struct {
	Pid              *int32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Result           *int32 `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2SSystem_RegisterResult) Reset()         { *m = S2SSystem_RegisterResult{} }
func (m *S2SSystem_RegisterResult) String() string { return proto.CompactTextString(m) }
func (*S2SSystem_RegisterResult) ProtoMessage()    {}

func (m *S2SSystem_RegisterResult) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *S2SSystem_RegisterResult) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}
