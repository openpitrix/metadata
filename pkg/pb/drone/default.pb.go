// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/drone/default.proto

package pbdrone // import "openpitrix.io/metadata/pkg/pb/drone"

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
	DroneServicePort     *int32   `protobuf:"varint,1,opt,name=DroneServicePort,def=9112" json:"DroneServicePort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Default) Reset()         { *m = Default{} }
func (m *Default) String() string { return proto.CompactTextString(m) }
func (*Default) ProtoMessage()    {}
func (*Default) Descriptor() ([]byte, []int) {
	return fileDescriptor_default_1f5d7b29bcb54290, []int{0}
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

const Default_Default_DroneServicePort int32 = 9112

func (m *Default) GetDroneServicePort() int32 {
	if m != nil && m.DroneServicePort != nil {
		return *m.DroneServicePort
	}
	return Default_Default_DroneServicePort
}

func init() {
	proto.RegisterType((*Default)(nil), "metadata.drone.Default")
}

func init() {
	proto.RegisterFile("metadata/drone/default.proto", fileDescriptor_default_1f5d7b29bcb54290)
}

var fileDescriptor_default_1f5d7b29bcb54290 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4d, 0x2d, 0x49,
	0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x29, 0xca, 0xcf, 0x4b, 0xd5, 0x4f, 0x49, 0x4d, 0x4b, 0x2c,
	0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xc9, 0xea, 0x81, 0x65, 0x95,
	0xac, 0xb9, 0xd8, 0x5d, 0x20, 0x0a, 0x84, 0x0c, 0xb8, 0x04, 0x5c, 0x40, 0x62, 0xc1, 0xa9, 0x45,
	0x65, 0x99, 0xc9, 0xa9, 0x01, 0xf9, 0x45, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x56, 0x2c,
	0x96, 0x86, 0x86, 0x46, 0x41, 0x18, 0xb2, 0x4e, 0xba, 0x51, 0xda, 0xf9, 0x05, 0xa9, 0x79, 0x05,
	0x99, 0x25, 0x45, 0x99, 0x15, 0x7a, 0x99, 0xf9, 0xfa, 0x70, 0xab, 0x0b, 0xb2, 0xd3, 0xf5, 0x0b,
	0x92, 0x20, 0x2e, 0xb0, 0x2e, 0x48, 0x02, 0xd3, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x13,
	0x55, 0x5a, 0x9a, 0x00, 0x00, 0x00,
}
