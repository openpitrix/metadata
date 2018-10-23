// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/task.proto

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

type SubTaskAction int32

const (
	SubTaskAction_NULL               SubTaskAction = 0
	SubTaskAction_StartConfd         SubTaskAction = 1
	SubTaskAction_RegisterMetadata   SubTaskAction = 2
	SubTaskAction_DeregisterMetadata SubTaskAction = 3
	SubTaskAction_RegisterCmd        SubTaskAction = 4
	SubTaskAction_DeregisterCmd      SubTaskAction = 5
	SubTaskAction_GetTaskStatus      SubTaskAction = 6
	SubTaskAction_StopConfd          SubTaskAction = 7
)

var SubTaskAction_name = map[int32]string{
	0: "NULL",
	1: "StartConfd",
	2: "RegisterMetadata",
	3: "DeregisterMetadata",
	4: "RegisterCmd",
	5: "DeregisterCmd",
	6: "GetTaskStatus",
	7: "StopConfd",
}
var SubTaskAction_value = map[string]int32{
	"NULL":               0,
	"StartConfd":         1,
	"RegisterMetadata":   2,
	"DeregisterMetadata": 3,
	"RegisterCmd":        4,
	"DeregisterCmd":      5,
	"GetTaskStatus":      6,
	"StopConfd":          7,
}

func (x SubTaskAction) String() string {
	return proto.EnumName(SubTaskAction_name, int32(x))
}
func (SubTaskAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{0}
}

type SubTaskId struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskId) Reset()         { *m = SubTaskId{} }
func (m *SubTaskId) String() string { return proto.CompactTextString(m) }
func (*SubTaskId) ProtoMessage()    {}
func (*SubTaskId) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{0}
}
func (m *SubTaskId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskId.Unmarshal(m, b)
}
func (m *SubTaskId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskId.Marshal(b, m, deterministic)
}
func (dst *SubTaskId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskId.Merge(dst, src)
}
func (m *SubTaskId) XXX_Size() int {
	return xxx_messageInfo_SubTaskId.Size(m)
}
func (m *SubTaskId) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskId.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskId proto.InternalMessageInfo

func (m *SubTaskId) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type SubTaskMessage struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	Directive            string   `protobuf:"bytes,3,opt,name=directive,proto3" json:"directive"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskMessage) Reset()         { *m = SubTaskMessage{} }
func (m *SubTaskMessage) String() string { return proto.CompactTextString(m) }
func (*SubTaskMessage) ProtoMessage()    {}
func (*SubTaskMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{1}
}
func (m *SubTaskMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskMessage.Unmarshal(m, b)
}
func (m *SubTaskMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskMessage.Marshal(b, m, deterministic)
}
func (dst *SubTaskMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskMessage.Merge(dst, src)
}
func (m *SubTaskMessage) XXX_Size() int {
	return xxx_messageInfo_SubTaskMessage.Size(m)
}
func (m *SubTaskMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskMessage proto.InternalMessageInfo

func (m *SubTaskMessage) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTaskMessage) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTaskMessage) GetDirective() string {
	if m != nil {
		return m.Directive
	}
	return ""
}

type SubTaskStatus struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskStatus) Reset()         { *m = SubTaskStatus{} }
func (m *SubTaskStatus) String() string { return proto.CompactTextString(m) }
func (*SubTaskStatus) ProtoMessage()    {}
func (*SubTaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{2}
}
func (m *SubTaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskStatus.Unmarshal(m, b)
}
func (m *SubTaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskStatus.Marshal(b, m, deterministic)
}
func (dst *SubTaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskStatus.Merge(dst, src)
}
func (m *SubTaskStatus) XXX_Size() int {
	return xxx_messageInfo_SubTaskStatus.Size(m)
}
func (m *SubTaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskStatus proto.InternalMessageInfo

func (m *SubTaskStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTaskStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
// {
// "action": "StartConfd",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", "timeout": 600}
// }
type SubTask_StartConfd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_StartConfd) Reset()         { *m = SubTask_StartConfd{} }
func (m *SubTask_StartConfd) String() string { return proto.CompactTextString(m) }
func (*SubTask_StartConfd) ProtoMessage()    {}
func (*SubTask_StartConfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{3}
}
func (m *SubTask_StartConfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_StartConfd.Unmarshal(m, b)
}
func (m *SubTask_StartConfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_StartConfd.Marshal(b, m, deterministic)
}
func (dst *SubTask_StartConfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_StartConfd.Merge(dst, src)
}
func (m *SubTask_StartConfd) XXX_Size() int {
	return xxx_messageInfo_SubTask_StartConfd.Size(m)
}
func (m *SubTask_StartConfd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_StartConfd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_StartConfd proto.InternalMessageInfo

func (m *SubTask_StartConfd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_StartConfd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_StartConfd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_StartConfd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_StartConfd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

//
// {
// "action": "StopConfd",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", "timeout": 600}
// }
type SubTask_StopConfd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_StopConfd) Reset()         { *m = SubTask_StopConfd{} }
func (m *SubTask_StopConfd) String() string { return proto.CompactTextString(m) }
func (*SubTask_StopConfd) ProtoMessage()    {}
func (*SubTask_StopConfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{4}
}
func (m *SubTask_StopConfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_StopConfd.Unmarshal(m, b)
}
func (m *SubTask_StopConfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_StopConfd.Marshal(b, m, deterministic)
}
func (dst *SubTask_StopConfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_StopConfd.Merge(dst, src)
}
func (m *SubTask_StopConfd) XXX_Size() int {
	return xxx_messageInfo_SubTask_StopConfd.Size(m)
}
func (m *SubTask_StopConfd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_StopConfd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_StopConfd proto.InternalMessageInfo

func (m *SubTask_StopConfd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_StopConfd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_StopConfd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_StopConfd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_StopConfd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

//
// {
// "action": "RegisterMetadata",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
// }
type SubTask_RegisterMetadata struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	Cnodes               string   `protobuf:"bytes,4,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,6,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_RegisterMetadata) Reset()         { *m = SubTask_RegisterMetadata{} }
func (m *SubTask_RegisterMetadata) String() string { return proto.CompactTextString(m) }
func (*SubTask_RegisterMetadata) ProtoMessage()    {}
func (*SubTask_RegisterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{5}
}
func (m *SubTask_RegisterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_RegisterMetadata.Unmarshal(m, b)
}
func (m *SubTask_RegisterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_RegisterMetadata.Marshal(b, m, deterministic)
}
func (dst *SubTask_RegisterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_RegisterMetadata.Merge(dst, src)
}
func (m *SubTask_RegisterMetadata) XXX_Size() int {
	return xxx_messageInfo_SubTask_RegisterMetadata.Size(m)
}
func (m *SubTask_RegisterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_RegisterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_RegisterMetadata proto.InternalMessageInfo

func (m *SubTask_RegisterMetadata) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_RegisterMetadata) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
// {
// "action": "DeregisterMetadata",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
// }
type SubTask_DeregisterMetadata struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_DeregisterMetadata) Reset()         { *m = SubTask_DeregisterMetadata{} }
func (m *SubTask_DeregisterMetadata) String() string { return proto.CompactTextString(m) }
func (*SubTask_DeregisterMetadata) ProtoMessage()    {}
func (*SubTask_DeregisterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{6}
}
func (m *SubTask_DeregisterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Unmarshal(m, b)
}
func (m *SubTask_DeregisterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Marshal(b, m, deterministic)
}
func (dst *SubTask_DeregisterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_DeregisterMetadata.Merge(dst, src)
}
func (m *SubTask_DeregisterMetadata) XXX_Size() int {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Size(m)
}
func (m *SubTask_DeregisterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_DeregisterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_DeregisterMetadata proto.InternalMessageInfo

func (m *SubTask_DeregisterMetadata) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_DeregisterMetadata) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
// {
// "action": "RegisterCmd",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
// }
type SubTask_RegisterCmd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_RegisterCmd) Reset()         { *m = SubTask_RegisterCmd{} }
func (m *SubTask_RegisterCmd) String() string { return proto.CompactTextString(m) }
func (*SubTask_RegisterCmd) ProtoMessage()    {}
func (*SubTask_RegisterCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{7}
}
func (m *SubTask_RegisterCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_RegisterCmd.Unmarshal(m, b)
}
func (m *SubTask_RegisterCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_RegisterCmd.Marshal(b, m, deterministic)
}
func (dst *SubTask_RegisterCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_RegisterCmd.Merge(dst, src)
}
func (m *SubTask_RegisterCmd) XXX_Size() int {
	return xxx_messageInfo_SubTask_RegisterCmd.Size(m)
}
func (m *SubTask_RegisterCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_RegisterCmd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_RegisterCmd proto.InternalMessageInfo

func (m *SubTask_RegisterCmd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_RegisterCmd) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
// {
// "action": "DeregisterCmd",
// "taskId": "t-abcdefgh",
// "directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
// }
type SubTask_DeregisterCmd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_DeregisterCmd) Reset()         { *m = SubTask_DeregisterCmd{} }
func (m *SubTask_DeregisterCmd) String() string { return proto.CompactTextString(m) }
func (*SubTask_DeregisterCmd) ProtoMessage()    {}
func (*SubTask_DeregisterCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{8}
}
func (m *SubTask_DeregisterCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_DeregisterCmd.Unmarshal(m, b)
}
func (m *SubTask_DeregisterCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_DeregisterCmd.Marshal(b, m, deterministic)
}
func (dst *SubTask_DeregisterCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_DeregisterCmd.Merge(dst, src)
}
func (m *SubTask_DeregisterCmd) XXX_Size() int {
	return xxx_messageInfo_SubTask_DeregisterCmd.Size(m)
}
func (m *SubTask_DeregisterCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_DeregisterCmd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_DeregisterCmd proto.InternalMessageInfo

func (m *SubTask_DeregisterCmd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_DeregisterCmd) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
// {
// "action": "GetTaskStatus",
// "taskId": "t-abcdefgh"
// }
type SubTask_GetTaskStatus struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_GetTaskStatus) Reset()         { *m = SubTask_GetTaskStatus{} }
func (m *SubTask_GetTaskStatus) String() string { return proto.CompactTextString(m) }
func (*SubTask_GetTaskStatus) ProtoMessage()    {}
func (*SubTask_GetTaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_2fd763836eca3bb9, []int{9}
}
func (m *SubTask_GetTaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_GetTaskStatus.Unmarshal(m, b)
}
func (m *SubTask_GetTaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_GetTaskStatus.Marshal(b, m, deterministic)
}
func (dst *SubTask_GetTaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_GetTaskStatus.Merge(dst, src)
}
func (m *SubTask_GetTaskStatus) XXX_Size() int {
	return xxx_messageInfo_SubTask_GetTaskStatus.Size(m)
}
func (m *SubTask_GetTaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_GetTaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_GetTaskStatus proto.InternalMessageInfo

func (m *SubTask_GetTaskStatus) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_GetTaskStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterType((*SubTaskId)(nil), "metadata.types.SubTaskId")
	proto.RegisterType((*SubTaskMessage)(nil), "metadata.types.SubTaskMessage")
	proto.RegisterType((*SubTaskStatus)(nil), "metadata.types.SubTaskStatus")
	proto.RegisterType((*SubTask_StartConfd)(nil), "metadata.types.SubTask_StartConfd")
	proto.RegisterType((*SubTask_StopConfd)(nil), "metadata.types.SubTask_StopConfd")
	proto.RegisterType((*SubTask_RegisterMetadata)(nil), "metadata.types.SubTask_RegisterMetadata")
	proto.RegisterType((*SubTask_DeregisterMetadata)(nil), "metadata.types.SubTask_DeregisterMetadata")
	proto.RegisterType((*SubTask_RegisterCmd)(nil), "metadata.types.SubTask_RegisterCmd")
	proto.RegisterType((*SubTask_DeregisterCmd)(nil), "metadata.types.SubTask_DeregisterCmd")
	proto.RegisterType((*SubTask_GetTaskStatus)(nil), "metadata.types.SubTask_GetTaskStatus")
	proto.RegisterEnum("metadata.types.SubTaskAction", SubTaskAction_name, SubTaskAction_value)
}

func init() { proto.RegisterFile("metadata/types/task.proto", fileDescriptor_task_2fd763836eca3bb9) }

var fileDescriptor_task_2fd763836eca3bb9 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x71, 0x9b, 0xdd, 0x6d, 0xa6, 0x24, 0xb8, 0xa6, 0x84, 0x2d, 0xe2, 0x50, 0x56, 0x1c,
	0x2a, 0x10, 0xc9, 0x81, 0x23, 0x17, 0xa0, 0x48, 0x10, 0xa9, 0xe5, 0x90, 0xc0, 0x85, 0x4b, 0xe4,
	0xc4, 0x6e, 0x64, 0x45, 0x59, 0x5b, 0xf6, 0x04, 0xd1, 0x77, 0xa1, 0xaf, 0xc1, 0x53, 0x20, 0x2e,
	0xbc, 0x10, 0x8a, 0xf7, 0x6f, 0x5b, 0x2d, 0x52, 0x0f, 0xa0, 0xe6, 0xb4, 0x9a, 0xf9, 0x46, 0xe3,
	0xdf, 0x67, 0x7b, 0xc7, 0x70, 0xb0, 0x94, 0xc8, 0x05, 0x47, 0x3e, 0xc0, 0x73, 0x23, 0xdd, 0x00,
	0xb9, 0x5b, 0xf4, 0x8d, 0xd5, 0xa8, 0x59, 0xb7, 0x90, 0xfa, 0x5e, 0x4a, 0x9e, 0x42, 0x7b, 0xbc,
	0x9a, 0x7e, 0xe2, 0x6e, 0x31, 0x14, 0xec, 0x21, 0x44, 0xeb, 0xd2, 0x89, 0x12, 0x31, 0x39, 0x24,
	0x47, 0xed, 0x51, 0x88, 0x5e, 0x48, 0x26, 0xd0, 0xcd, 0xab, 0x4e, 0xa5, 0x73, 0x7c, 0x2e, 0x59,
	0x0f, 0x42, 0x3e, 0x43, 0xa5, 0xd3, 0xa2, 0x32, 0x8b, 0xea, 0x2d, 0xb6, 0xea, 0x2d, 0xd8, 0x63,
	0x68, 0x0b, 0x65, 0xe5, 0x0c, 0xd5, 0x57, 0x19, 0x6f, 0x7b, 0xa9, 0x4a, 0x24, 0xaf, 0xa1, 0x93,
	0x2f, 0x30, 0x46, 0x8e, 0x2b, 0xd7, 0x88, 0xb2, 0x5e, 0xd8, 0xf9, 0x92, 0xa2, 0x7f, 0x16, 0x25,
	0x17, 0x04, 0x58, 0xde, 0x62, 0x32, 0x46, 0x6e, 0xf1, 0x58, 0xa7, 0x67, 0xe2, 0xe6, 0x9c, 0x4f,
	0xe0, 0xee, 0x99, 0xd5, 0x29, 0xce, 0x39, 0xca, 0xb5, 0x9a, 0xa1, 0xee, 0x96, 0xb9, 0xa1, 0x60,
	0x07, 0xb0, 0x23, 0xac, 0x4e, 0xe5, 0x44, 0x99, 0xb8, 0xe5, 0xe5, 0xc8, 0xc7, 0x43, 0xc3, 0x62,
	0x88, 0x50, 0x2d, 0xa5, 0x5e, 0x61, 0x1c, 0x1c, 0x92, 0xa3, 0x60, 0x54, 0x84, 0xc9, 0x77, 0x02,
	0x7b, 0x15, 0x9f, 0x36, 0xb7, 0x0c, 0xef, 0x07, 0x81, 0xb8, 0xc0, 0x1b, 0xc9, 0xb9, 0x72, 0x28,
	0xed, 0x69, 0x7e, 0x55, 0xfe, 0x09, 0x65, 0x0f, 0xc2, 0x59, 0xaa, 0x85, 0x74, 0x39, 0x63, 0x1e,
	0x35, 0x23, 0xb2, 0x7d, 0x08, 0xac, 0x44, 0x7b, 0x1e, 0x87, 0x3e, 0x9f, 0x05, 0xc9, 0x6f, 0x02,
	0x8f, 0x0a, 0xf0, 0x77, 0xd2, 0xfe, 0x0f, 0xf4, 0xbf, 0x6c, 0x70, 0xe5, 0x2a, 0x68, 0x72, 0x15,
	0x36, 0xb8, 0x8a, 0xea, 0xae, 0x7e, 0x12, 0xb8, 0x7f, 0xf5, 0x38, 0x8e, 0x97, 0x62, 0x53, 0xed,
	0xfc, 0x22, 0xf0, 0xe0, 0xfa, 0x21, 0x6d, 0xb0, 0xa1, 0x0f, 0x95, 0x9f, 0xf7, 0x12, 0x6b, 0x73,
	0xeb, 0xa6, 0x7e, 0x9e, 0x5d, 0x90, 0x72, 0xf4, 0xbd, 0xc9, 0x4a, 0x77, 0xa0, 0xf5, 0xf1, 0xf3,
	0xc9, 0x09, 0xbd, 0xc3, 0xba, 0x00, 0xd5, 0x28, 0xa3, 0x84, 0xed, 0x03, 0xbd, 0xfa, 0x6f, 0xd2,
	0x2d, 0xd6, 0x03, 0x76, 0xfd, 0xe2, 0xd3, 0x6d, 0x76, 0x0f, 0x76, 0x6b, 0x57, 0x87, 0xb6, 0xd8,
	0x1e, 0x74, 0x2e, 0x6d, 0x3e, 0x0d, 0xd6, 0xa9, 0x4b, 0xfc, 0x34, 0x64, 0x1d, 0x68, 0x97, 0xf3,
	0x89, 0x46, 0x6f, 0x5f, 0x7c, 0x79, 0xae, 0x8d, 0x4c, 0x8d, 0x42, 0xab, 0xbe, 0xf5, 0x95, 0x1e,
	0x94, 0x6f, 0x8b, 0x59, 0xcc, 0x07, 0x66, 0x9a, 0x3d, 0x31, 0xaf, 0xcc, 0xd4, 0x7f, 0xa7, 0xa1,
	0x7f, 0x66, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xd2, 0x94, 0x05, 0x83, 0x06, 0x00,
	0x00,
}
