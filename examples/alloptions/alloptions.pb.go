// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alloptions.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	alloptions.proto

It has these top-level messages:
	NoArgs
	StringArg
	SimpleStringReply
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/rapidloop/nrpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NoArgs struct {
}

func (m *NoArgs) Reset()                    { *m = NoArgs{} }
func (m *NoArgs) String() string            { return proto.CompactTextString(m) }
func (*NoArgs) ProtoMessage()               {}
func (*NoArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StringArg struct {
	Arg1 string `protobuf:"bytes,1,opt,name=arg1" json:"arg1,omitempty"`
}

func (m *StringArg) Reset()                    { *m = StringArg{} }
func (m *StringArg) String() string            { return proto.CompactTextString(m) }
func (*StringArg) ProtoMessage()               {}
func (*StringArg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StringArg) GetArg1() string {
	if m != nil {
		return m.Arg1
	}
	return ""
}

type SimpleStringReply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *SimpleStringReply) Reset()                    { *m = SimpleStringReply{} }
func (m *SimpleStringReply) String() string            { return proto.CompactTextString(m) }
func (*SimpleStringReply) ProtoMessage()               {}
func (*SimpleStringReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SimpleStringReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*NoArgs)(nil), "main.NoArgs")
	proto.RegisterType((*StringArg)(nil), "main.StringArg")
	proto.RegisterType((*SimpleStringReply)(nil), "main.SimpleStringReply")
}

func init() { proto.RegisterFile("alloptions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0xd0, 0x4f, 0x6a, 0x3a, 0x31,
	0x14, 0x07, 0x70, 0x86, 0x9f, 0x3f, 0xd1, 0x60, 0xd5, 0xc6, 0x42, 0xe3, 0x50, 0x68, 0x19, 0xba,
	0x68, 0x37, 0x23, 0xda, 0x13, 0xd8, 0x6e, 0xab, 0x14, 0x67, 0xd1, 0xa5, 0xc4, 0x38, 0x8c, 0x29,
	0xf9, 0x47, 0xf2, 0x14, 0xba, 0x75, 0xe9, 0xaa, 0xcb, 0x9e, 0x43, 0x98, 0x0b, 0x74, 0xd9, 0xeb,
	0x78, 0x81, 0x62, 0x32, 0x14, 0x5a, 0xe8, 0x66, 0xe6, 0xe5, 0xe5, 0x9b, 0x7c, 0x92, 0xa0, 0x2e,
	0x15, 0x42, 0x1b, 0xe0, 0x5a, 0xb9, 0xd4, 0x58, 0x0d, 0x1a, 0xd7, 0x24, 0xe5, 0x2a, 0xbe, 0x2e,
	0x38, 0xac, 0xd6, 0x8b, 0x94, 0x69, 0x39, 0xb0, 0xd4, 0xf0, 0xa5, 0xd0, 0xda, 0x0c, 0x94, 0x35,
	0xcc, 0x7f, 0x42, 0x36, 0x69, 0xa0, 0xfa, 0x54, 0x8f, 0x6d, 0xe1, 0x92, 0x4b, 0xd4, 0xcc, 0xc0,
	0x72, 0x55, 0x8c, 0x6d, 0x81, 0x31, 0xaa, 0x51, 0x5b, 0x0c, 0x49, 0x74, 0x15, 0xdd, 0x34, 0x67,
	0xbe, 0x4e, 0x6e, 0xd1, 0x69, 0xc6, 0xa5, 0x11, 0x79, 0x88, 0xcd, 0x72, 0x23, 0x5e, 0xf1, 0x19,
	0xfa, 0x6f, 0x8f, 0x45, 0x95, 0x0c, 0x83, 0xd1, 0x06, 0x75, 0xb3, 0x0d, 0x7b, 0x58, 0x3b, 0xd0,
	0x32, 0x5b, 0x2f, 0x5e, 0x72, 0x06, 0x78, 0x8a, 0x4e, 0x26, 0x10, 0x36, 0x08, 0x4b, 0x3b, 0xe9,
	0xf1, 0x9c, 0xe9, 0x37, 0x1a, 0x9f, 0x57, 0x8d, 0xdf, 0x48, 0xd2, 0xdb, 0xee, 0xfb, 0x1d, 0x09,
	0x73, 0xe7, 0x67, 0xe6, 0xde, 0x88, 0xf1, 0xc7, 0x81, 0xb4, 0x99, 0x27, 0xe6, 0x2e, 0x18, 0x23,
	0xe5, 0xdd, 0x4a, 0x7c, 0xa2, 0x96, 0x4a, 0x87, 0x1f, 0x51, 0x6f, 0x02, 0xcf, 0x1c, 0x56, 0x3f,
	0xdb, 0xad, 0x80, 0x85, 0xcb, 0xff, 0x4d, 0xb7, 0x77, 0xfb, 0xfe, 0x3f, 0x69, 0x86, 0xe1, 0x37,
	0x8a, 0x5b, 0x9f, 0x07, 0xd2, 0x60, 0x82, 0xe7, 0x0a, 0xf8, 0xf2, 0xfe, 0x62, 0x5b, 0x92, 0x9a,
	0xd5, 0x1a, 0x76, 0x25, 0x69, 0x70, 0xe5, 0x80, 0x2a, 0x96, 0xbf, 0x95, 0x24, 0x7a, 0x2f, 0x49,
	0xb4, 0xa8, 0xfb, 0x27, 0xbe, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x97, 0x98, 0xbe, 0xd7, 0xa2,
	0x01, 0x00, 0x00,
}