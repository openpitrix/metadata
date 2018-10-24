// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/pilot/default.proto

package pbpilot // import "openpitrix.io/metadata/pkg/pb/pilot"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Default struct {
	PilotServicePort     *int32   `protobuf:"varint,1,opt,name=PilotServicePort,def=9110" json:"PilotServicePort,omitempty"`
	PilotServiceHost     *string  `protobuf:"bytes,2,opt,name=PilotServiceHost,def=openpitrix-pilot-service" json:"PilotServiceHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Default) Reset()         { *m = Default{} }
func (m *Default) String() string { return proto.CompactTextString(m) }
func (*Default) ProtoMessage()    {}
func (*Default) Descriptor() ([]byte, []int) {
	return fileDescriptor_default_09f22fd2e56ae29d, []int{0}
}
func (m *Default) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Default.Unmarshal(m, b)
}
func (m *Default) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Default.Marshal(b, m, deterministic)
}
func (dst *Default) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Default.Merge(dst, src)
}
func (m *Default) XXX_Size() int {
	return xxx_messageInfo_Default.Size(m)
}
func (m *Default) XXX_DiscardUnknown() {
	xxx_messageInfo_Default.DiscardUnknown(m)
}

var xxx_messageInfo_Default proto.InternalMessageInfo

const Default_Default_PilotServicePort int32 = 9110
const Default_Default_PilotServiceHost string = "openpitrix-pilot-service"

func (m *Default) GetPilotServicePort() int32 {
	if m != nil && m.PilotServicePort != nil {
		return *m.PilotServicePort
	}
	return Default_Default_PilotServicePort
}

func (m *Default) GetPilotServiceHost() string {
	if m != nil && m.PilotServiceHost != nil {
		return *m.PilotServiceHost
	}
	return Default_Default_PilotServiceHost
}

func init() {
	proto.RegisterType((*Default)(nil), "metadata.pilot.Default")
}

func init() {
	proto.RegisterFile("metadata/pilot/default.proto", fileDescriptor_default_09f22fd2e56ae29d)
}

var fileDescriptor_default_09f22fd2e56ae29d = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4d, 0x2d, 0x49,
	0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2f, 0xc8, 0xcc, 0xc9, 0x2f, 0xd1, 0x4f, 0x49, 0x4d, 0x4b, 0x2c,
	0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xc9, 0xea, 0x81, 0x65, 0x95,
	0x1a, 0x19, 0xb9, 0xd8, 0x5d, 0x20, 0x2a, 0x84, 0x0c, 0xb8, 0x04, 0x02, 0x40, 0x82, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x01, 0xf9, 0x45, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x56,
	0x2c, 0x96, 0x86, 0x86, 0x06, 0x41, 0x18, 0xb2, 0x42, 0x2e, 0xa8, 0x3a, 0x3c, 0xf2, 0x8b, 0x4b,
	0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0xad, 0x24, 0xf2, 0x0b, 0x52, 0xf3, 0x0a, 0x32, 0x4b, 0x8a,
	0x32, 0x2b, 0x74, 0xc1, 0x36, 0xe9, 0x16, 0x43, 0xd4, 0x04, 0x61, 0xe8, 0x70, 0xd2, 0x8d, 0xd2,
	0x46, 0xa8, 0xd6, 0xcb, 0xcc, 0xd7, 0x47, 0xf8, 0x20, 0x3b, 0x5d, 0xbf, 0x20, 0x09, 0xe2, 0x11,
	0xeb, 0x82, 0x24, 0x30, 0x0d, 0x08, 0x00, 0x00, 0xff, 0xff, 0x09, 0xf9, 0xce, 0x59, 0xe1, 0x00,
	0x00, 0x00,
}
