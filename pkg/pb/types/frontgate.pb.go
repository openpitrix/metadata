// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/frontgate.proto

package pbtypes // import "openpitrix.io/metadata/pkg/pb/types"

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

type FrontgateId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateId) Reset()         { *m = FrontgateId{} }
func (m *FrontgateId) String() string { return proto.CompactTextString(m) }
func (*FrontgateId) ProtoMessage()    {}
func (*FrontgateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{0}
}
func (m *FrontgateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateId.Unmarshal(m, b)
}
func (m *FrontgateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateId.Marshal(b, m, deterministic)
}
func (dst *FrontgateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateId.Merge(dst, src)
}
func (m *FrontgateId) XXX_Size() int {
	return xxx_messageInfo_FrontgateId.Size(m)
}
func (m *FrontgateId) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateId.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateId proto.InternalMessageInfo

func (m *FrontgateId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateId) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type FrontgateNodeId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateNodeId) Reset()         { *m = FrontgateNodeId{} }
func (m *FrontgateNodeId) String() string { return proto.CompactTextString(m) }
func (*FrontgateNodeId) ProtoMessage()    {}
func (*FrontgateNodeId) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{1}
}
func (m *FrontgateNodeId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateNodeId.Unmarshal(m, b)
}
func (m *FrontgateNodeId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateNodeId.Marshal(b, m, deterministic)
}
func (dst *FrontgateNodeId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateNodeId.Merge(dst, src)
}
func (m *FrontgateNodeId) XXX_Size() int {
	return xxx_messageInfo_FrontgateNodeId.Size(m)
}
func (m *FrontgateNodeId) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateNodeId.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateNodeId proto.InternalMessageInfo

func (m *FrontgateNodeId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateNodeId) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *FrontgateNodeId) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type FrontgateIdList struct {
	IdList               []string `protobuf:"bytes,1,rep,name=id_list,json=idList,proto3" json:"id_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateIdList) Reset()         { *m = FrontgateIdList{} }
func (m *FrontgateIdList) String() string { return proto.CompactTextString(m) }
func (*FrontgateIdList) ProtoMessage()    {}
func (*FrontgateIdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{2}
}
func (m *FrontgateIdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateIdList.Unmarshal(m, b)
}
func (m *FrontgateIdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateIdList.Marshal(b, m, deterministic)
}
func (dst *FrontgateIdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateIdList.Merge(dst, src)
}
func (m *FrontgateIdList) XXX_Size() int {
	return xxx_messageInfo_FrontgateIdList.Size(m)
}
func (m *FrontgateIdList) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateIdList.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateIdList proto.InternalMessageInfo

func (m *FrontgateIdList) GetIdList() []string {
	if m != nil {
		return m.IdList
	}
	return nil
}

type FrontgateConfig struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	NodeId               string               `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	Host                 string               `protobuf:"bytes,3,opt,name=host,proto3" json:"host"`
	ListenPort           int32                `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port"`
	PilotHost            string               `protobuf:"bytes,5,opt,name=pilot_host,json=pilotHost,proto3" json:"pilot_host"`
	PilotPort            int32                `protobuf:"varint,6,opt,name=pilot_port,json=pilotPort,proto3" json:"pilot_port"`
	NodeList             []*FrontgateEndpoint `protobuf:"bytes,7,rep,name=node_list,json=nodeList,proto3" json:"node_list"`
	EtcdConfig           *EtcdConfig          `protobuf:"bytes,8,opt,name=etcd_config,json=etcdConfig,proto3" json:"etcd_config"`
	ConfdConfig          *ConfdConfig         `protobuf:"bytes,9,opt,name=confd_config,json=confdConfig,proto3" json:"confd_config"`
	LogLevel             string               `protobuf:"bytes,10,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	AutoUpdate           bool                 `protobuf:"varint,11,opt,name=auto_update,json=autoUpdate,proto3" json:"auto_update"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrontgateConfig) Reset()         { *m = FrontgateConfig{} }
func (m *FrontgateConfig) String() string { return proto.CompactTextString(m) }
func (*FrontgateConfig) ProtoMessage()    {}
func (*FrontgateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{3}
}
func (m *FrontgateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateConfig.Unmarshal(m, b)
}
func (m *FrontgateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateConfig.Marshal(b, m, deterministic)
}
func (dst *FrontgateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateConfig.Merge(dst, src)
}
func (m *FrontgateConfig) XXX_Size() int {
	return xxx_messageInfo_FrontgateConfig.Size(m)
}
func (m *FrontgateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateConfig proto.InternalMessageInfo

func (m *FrontgateConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrontgateConfig) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *FrontgateConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FrontgateConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *FrontgateConfig) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *FrontgateConfig) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func (m *FrontgateConfig) GetNodeList() []*FrontgateEndpoint {
	if m != nil {
		return m.NodeList
	}
	return nil
}

func (m *FrontgateConfig) GetEtcdConfig() *EtcdConfig {
	if m != nil {
		return m.EtcdConfig
	}
	return nil
}

func (m *FrontgateConfig) GetConfdConfig() *ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func (m *FrontgateConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *FrontgateConfig) GetAutoUpdate() bool {
	if m != nil {
		return m.AutoUpdate
	}
	return false
}

type FrontgateEndpoint struct {
	FrontgateId          string   `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	FrontgateNodeId      string   `protobuf:"bytes,2,opt,name=frontgate_node_id,json=frontgateNodeId,proto3" json:"frontgate_node_id"`
	NodeIp               string   `protobuf:"bytes,3,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip"`
	NodePort             int32    `protobuf:"varint,4,opt,name=node_port,json=nodePort,proto3" json:"node_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontgateEndpoint) Reset()         { *m = FrontgateEndpoint{} }
func (m *FrontgateEndpoint) String() string { return proto.CompactTextString(m) }
func (*FrontgateEndpoint) ProtoMessage()    {}
func (*FrontgateEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{4}
}
func (m *FrontgateEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontgateEndpoint.Unmarshal(m, b)
}
func (m *FrontgateEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontgateEndpoint.Marshal(b, m, deterministic)
}
func (dst *FrontgateEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontgateEndpoint.Merge(dst, src)
}
func (m *FrontgateEndpoint) XXX_Size() int {
	return xxx_messageInfo_FrontgateEndpoint.Size(m)
}
func (m *FrontgateEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontgateEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_FrontgateEndpoint proto.InternalMessageInfo

func (m *FrontgateEndpoint) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *FrontgateEndpoint) GetFrontgateNodeId() string {
	if m != nil {
		return m.FrontgateNodeId
	}
	return ""
}

func (m *FrontgateEndpoint) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *FrontgateEndpoint) GetNodePort() int32 {
	if m != nil {
		return m.NodePort
	}
	return 0
}

type RunCommandOnFrontgateRequest struct {
	Endpoint             *FrontgateEndpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint"`
	Command              string             `protobuf:"bytes,2,opt,name=command,proto3" json:"command"`
	TimeoutSeconds       int32              `protobuf:"varint,3,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RunCommandOnFrontgateRequest) Reset()         { *m = RunCommandOnFrontgateRequest{} }
func (m *RunCommandOnFrontgateRequest) String() string { return proto.CompactTextString(m) }
func (*RunCommandOnFrontgateRequest) ProtoMessage()    {}
func (*RunCommandOnFrontgateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_frontgate_c1e8993c386293d6, []int{5}
}
func (m *RunCommandOnFrontgateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunCommandOnFrontgateRequest.Unmarshal(m, b)
}
func (m *RunCommandOnFrontgateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunCommandOnFrontgateRequest.Marshal(b, m, deterministic)
}
func (dst *RunCommandOnFrontgateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunCommandOnFrontgateRequest.Merge(dst, src)
}
func (m *RunCommandOnFrontgateRequest) XXX_Size() int {
	return xxx_messageInfo_RunCommandOnFrontgateRequest.Size(m)
}
func (m *RunCommandOnFrontgateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunCommandOnFrontgateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunCommandOnFrontgateRequest proto.InternalMessageInfo

func (m *RunCommandOnFrontgateRequest) GetEndpoint() *FrontgateEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *RunCommandOnFrontgateRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *RunCommandOnFrontgateRequest) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*FrontgateId)(nil), "metadata.types.FrontgateId")
	proto.RegisterType((*FrontgateNodeId)(nil), "metadata.types.FrontgateNodeId")
	proto.RegisterType((*FrontgateIdList)(nil), "metadata.types.FrontgateIdList")
	proto.RegisterType((*FrontgateConfig)(nil), "metadata.types.FrontgateConfig")
	proto.RegisterType((*FrontgateEndpoint)(nil), "metadata.types.FrontgateEndpoint")
	proto.RegisterType((*RunCommandOnFrontgateRequest)(nil), "metadata.types.RunCommandOnFrontgateRequest")
}

func init() {
	proto.RegisterFile("metadata/types/frontgate.proto", fileDescriptor_frontgate_c1e8993c386293d6)
}

var fileDescriptor_frontgate_c1e8993c386293d6 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x93, 0xe6, 0x6b, 0x5c, 0x25, 0xea, 0x5e, 0x7e, 0xfe, 0xa5, 0x7c, 0xb8, 0xb9, 0x10,
	0x15, 0x91, 0x48, 0xe5, 0xc0, 0xa1, 0xa2, 0x07, 0xaa, 0x22, 0x2a, 0x55, 0x80, 0x0c, 0x5c, 0xb8,
	0x58, 0x8e, 0x77, 0x13, 0x56, 0x38, 0x3b, 0x4b, 0x3c, 0x41, 0xf4, 0x3f, 0xe1, 0xc8, 0x91, 0x3f,
	0x13, 0xed, 0xf8, 0xa3, 0x89, 0xe1, 0xd0, 0x93, 0xbd, 0xf3, 0xe6, 0xbd, 0x9d, 0xd9, 0x37, 0x03,
	0x8f, 0xd6, 0x8a, 0x12, 0x99, 0x50, 0x32, 0xa7, 0x5b, 0xab, 0xf2, 0xf9, 0x72, 0x83, 0x86, 0x56,
	0x09, 0xa9, 0x99, 0xdd, 0x20, 0xa1, 0x18, 0x56, 0xf8, 0x8c, 0xf1, 0xf1, 0xff, 0x8d, 0x7c, 0x45,
	0xa9, 0x2c, 0x52, 0xc7, 0xe3, 0x06, 0x94, 0xa2, 0x59, 0x96, 0xd8, 0xe4, 0x05, 0xf8, 0xaf, 0x2b,
	0xe5, 0x6b, 0x29, 0x86, 0xd0, 0xd2, 0x32, 0xf0, 0x42, 0x6f, 0x3a, 0x88, 0x5a, 0x5a, 0x8a, 0x00,
	0x7a, 0x36, 0xb9, 0xcd, 0x30, 0x91, 0x41, 0x8b, 0x83, 0xd5, 0x71, 0xf2, 0x11, 0x46, 0x35, 0xf1,
	0x2d, 0xca, 0x7f, 0x91, 0xff, 0x83, 0x9e, 0x41, 0xa9, 0x62, 0x5d, 0x91, 0xbb, 0xa6, 0x48, 0xdc,
	0x51, 0x6d, 0xef, 0xab, 0x9e, 0xee, 0xa8, 0x5e, 0xcb, 0x1b, 0x9d, 0x93, 0x53, 0xd1, 0x32, 0xce,
	0x74, 0x4e, 0x81, 0x17, 0xb6, 0x9d, 0x8a, 0x66, 0x60, 0xf2, 0xbb, 0xbd, 0x93, 0x7c, 0x89, 0x66,
	0xa9, 0x57, 0xf7, 0x2f, 0x41, 0xc0, 0xc1, 0x17, 0xcc, 0xa9, 0xbc, 0x9f, 0xff, 0xc5, 0x63, 0xf0,
	0xdd, 0x35, 0xca, 0xc4, 0x16, 0x37, 0x14, 0x1c, 0x84, 0xde, 0xb4, 0x13, 0x41, 0x11, 0x7a, 0x8f,
	0x1b, 0x12, 0x0f, 0x01, 0xac, 0xce, 0x90, 0x62, 0xa6, 0x76, 0x98, 0x3a, 0xe0, 0xc8, 0x1b, 0xc7,
	0xaf, 0x61, 0xa6, 0x77, 0x99, 0x5e, 0xc0, 0xcc, 0xbe, 0x80, 0x01, 0xd7, 0xc2, 0xad, 0xf4, 0xc2,
	0xf6, 0xd4, 0x3f, 0x3b, 0x99, 0xed, 0xbb, 0x38, 0xab, 0xfb, 0xb9, 0x32, 0xd2, 0xa2, 0x36, 0x14,
	0xf5, 0x1d, 0x87, 0x1f, 0xe2, 0x1c, 0x7c, 0x67, 0x6a, 0x9c, 0x72, 0xab, 0x41, 0x3f, 0xf4, 0xa6,
	0xfe, 0xd9, 0xb8, 0xa9, 0x70, 0x45, 0xa9, 0x2c, 0x1e, 0x23, 0x02, 0x55, 0xff, 0x8b, 0x0b, 0x38,
	0x64, 0xdb, 0x2b, 0xf6, 0x80, 0xd9, 0xc7, 0x4d, 0xb6, 0xcb, 0xae, 0xe8, 0x7e, 0x7a, 0x77, 0x10,
	0xc7, 0x30, 0xc8, 0x70, 0x15, 0x67, 0xea, 0xbb, 0xca, 0x02, 0xe0, 0xce, 0xfb, 0x19, 0xae, 0x6e,
	0xdc, 0xd9, 0x3d, 0x5c, 0xb2, 0x25, 0x8c, 0xb7, 0x56, 0x26, 0xa4, 0x02, 0x3f, 0xf4, 0xa6, 0xfd,
	0x08, 0x5c, 0xe8, 0x13, 0x47, 0x26, 0x3f, 0x3d, 0x38, 0xfa, 0xab, 0x35, 0x71, 0x02, 0x87, 0xf5,
	0x54, 0xc7, 0xb5, 0x6d, 0xfe, 0x72, 0x67, 0x1e, 0x4f, 0xe1, 0xe8, 0x2e, 0x65, 0xdf, 0xc9, 0xd1,
	0xb2, 0x31, 0x7e, 0xb5, 0xd7, 0xb6, 0x74, 0xb5, 0xf0, 0xda, 0xba, 0xda, 0x19, 0xd8, 0x71, 0x95,
	0x5f, 0xd5, 0xb9, 0x32, 0xf9, 0xe5, 0xc1, 0x83, 0x68, 0x6b, 0x2e, 0x71, 0xbd, 0x4e, 0x8c, 0x7c,
	0x67, 0xea, 0x32, 0x23, 0xf5, 0x6d, 0xab, 0x72, 0x12, 0x2f, 0xa1, 0xaf, 0xca, 0x8a, 0xb9, 0xc2,
	0xfb, 0xb9, 0x56, 0x51, 0xdc, 0xac, 0xa7, 0x85, 0x76, 0xb5, 0x41, 0xe5, 0x51, 0x3c, 0x81, 0x11,
	0xe9, 0xb5, 0xc2, 0x2d, 0xc5, 0xb9, 0x4a, 0xd1, 0xc8, 0x9c, 0xeb, 0xee, 0x44, 0xc3, 0x32, 0xfc,
	0xa1, 0x88, 0xbe, 0x7a, 0xf6, 0xf9, 0x29, 0x5a, 0x65, 0xac, 0xa6, 0x8d, 0xfe, 0x31, 0xd3, 0x38,
	0xaf, 0xf7, 0xd9, 0x7e, 0x5d, 0xcd, 0xed, 0xa2, 0x58, 0xeb, 0x73, 0xbb, 0xe0, 0xef, 0xa2, 0xcb,
	0x9b, 0xfd, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x70, 0x37, 0x0e, 0x42, 0x04, 0x00,
	0x00,
}
