// Code generated by protoc-gen-gogo.
// source: empty/empty.proto
// DO NOT EDIT!

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	empty/empty.proto

It has these top-level messages:
	Empty
*/
package empty

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorEmpty, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

var fileDescriptorEmpty = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcd, 0x2d, 0x28,
	0xa9, 0xd4, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xfc, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0x10, 0x5e, 0x52, 0x69, 0x9a, 0x12, 0x3b, 0x17, 0xab, 0x2b, 0x48, 0xde, 0x29, 0x9c,
	0x4b, 0x38, 0x39, 0x3f, 0x57, 0x0f, 0x4d, 0xde, 0x89, 0x0b, 0x2c, 0x1b, 0x00, 0xe2, 0x06, 0x30,
	0x46, 0xb1, 0x82, 0xcd, 0x5a, 0xc0, 0xc8, 0xf8, 0x83, 0x91, 0x71, 0x11, 0x13, 0xb3, 0x7b, 0x80,
	0xd3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0x96, 0x00, 0xa8, 0x16, 0xbd, 0xf0, 0xd4, 0x9c, 0x1c, 0xef,
	0xbc, 0xfc, 0xf2, 0xbc, 0x90, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x59, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x38, 0xd8, 0x5d, 0x8e, 0x00, 0x00, 0x00,
}