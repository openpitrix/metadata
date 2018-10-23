// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/drone/drone.proto

package pbdrone // import "openpitrix.io/metadata/pkg/pb/drone"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "openpitrix.io/metadata/pkg/pb/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DroneServiceClient is the client API for DroneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DroneServiceClient interface {
	DistributeDrone(ctx context.Context, in *types.DroneBinary, opts ...grpc.CallOption) (*types.Empty, error)
	GetPilotVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Version, error)
	GetFrontgateVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Version, error)
	GetDroneVersion(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Version, error)
	GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error)
	SetDroneConfig(ctx context.Context, in *types.DroneConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetConfdConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.ConfdConfig, error)
	SetConfdConfig(ctx context.Context, in *types.ConfdConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetFrontgateConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateConfig, error)
	SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error)
	IsConfdRunning(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Bool, error)
	StartConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	StopConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	GetTemplateFiles(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringList, error)
	GetValues(ctx context.Context, in *types.StringList, opts ...grpc.CallOption) (*types.StringMap, error)
	PingPilot(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	PingFrontgate(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	PingDrone(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	PingMetadataBackend(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	RunCommand(ctx context.Context, in *types.RunCommandOnDroneRequest, opts ...grpc.CallOption) (*types.String, error)
}

type droneServiceClient struct {
	cc *grpc.ClientConn
}

func NewDroneServiceClient(cc *grpc.ClientConn) DroneServiceClient {
	return &droneServiceClient{cc}
}

func (c *droneServiceClient) DistributeDrone(ctx context.Context, in *types.DroneBinary, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/DistributeDrone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetPilotVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Version, error) {
	out := new(types.Version)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetPilotVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetFrontgateVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Version, error) {
	out := new(types.Version)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetFrontgateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetDroneVersion(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Version, error) {
	out := new(types.Version)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetDroneVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error) {
	out := new(types.DroneConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetDroneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetDroneConfig(ctx context.Context, in *types.DroneConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetDroneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetConfdConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.ConfdConfig, error) {
	out := new(types.ConfdConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetConfdConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetConfdConfig(ctx context.Context, in *types.ConfdConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetConfdConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetFrontgateConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateConfig, error) {
	out := new(types.FrontgateConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetFrontgateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetFrontgateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) IsConfdRunning(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Bool, error) {
	out := new(types.Bool)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/IsConfdRunning", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StartConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/StartConfd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StopConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/StopConfd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetTemplateFiles(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringList, error) {
	out := new(types.StringList)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetTemplateFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetValues(ctx context.Context, in *types.StringList, opts ...grpc.CallOption) (*types.StringMap, error) {
	out := new(types.StringMap)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingPilot(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingPilot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingFrontgate(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingFrontgate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingDrone(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingDrone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingMetadataBackend(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingMetadataBackend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) RunCommand(ctx context.Context, in *types.RunCommandOnDroneRequest, opts ...grpc.CallOption) (*types.String, error) {
	out := new(types.String)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DroneServiceServer is the server API for DroneService service.
type DroneServiceServer interface {
	DistributeDrone(context.Context, *types.DroneBinary) (*types.Empty, error)
	GetPilotVersion(context.Context, *types.Empty) (*types.Version, error)
	GetFrontgateVersion(context.Context, *types.Empty) (*types.Version, error)
	GetDroneVersion(context.Context, *types.DroneEndpoint) (*types.Version, error)
	GetDroneConfig(context.Context, *types.Empty) (*types.DroneConfig, error)
	SetDroneConfig(context.Context, *types.DroneConfig) (*types.Empty, error)
	GetConfdConfig(context.Context, *types.Empty) (*types.ConfdConfig, error)
	SetConfdConfig(context.Context, *types.ConfdConfig) (*types.Empty, error)
	GetFrontgateConfig(context.Context, *types.Empty) (*types.FrontgateConfig, error)
	SetFrontgateConfig(context.Context, *types.FrontgateConfig) (*types.Empty, error)
	IsConfdRunning(context.Context, *types.Empty) (*types.Bool, error)
	StartConfd(context.Context, *types.Empty) (*types.Empty, error)
	StopConfd(context.Context, *types.Empty) (*types.Empty, error)
	GetTemplateFiles(context.Context, *types.Empty) (*types.StringList, error)
	GetValues(context.Context, *types.StringList) (*types.StringMap, error)
	PingPilot(context.Context, *types.FrontgateEndpoint) (*types.Empty, error)
	PingFrontgate(context.Context, *types.FrontgateEndpoint) (*types.Empty, error)
	PingDrone(context.Context, *types.Empty) (*types.Empty, error)
	PingMetadataBackend(context.Context, *types.FrontgateEndpoint) (*types.Empty, error)
	RunCommand(context.Context, *types.RunCommandOnDroneRequest) (*types.String, error)
}

func RegisterDroneServiceServer(s *grpc.Server, srv DroneServiceServer) {
	s.RegisterService(&_DroneService_serviceDesc, srv)
}

func _DroneService_DistributeDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneBinary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).DistributeDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/DistributeDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).DistributeDrone(ctx, req.(*types.DroneBinary))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetPilotVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetPilotVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetPilotVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetPilotVersion(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetFrontgateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetFrontgateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetFrontgateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetFrontgateVersion(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetDroneVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetDroneVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetDroneVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetDroneVersion(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetDroneConfig(ctx, req.(*types.DroneConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ConfdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, req.(*types.ConfdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, req.(*types.FrontgateConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_IsConfdRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/IsConfdRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StartConfd(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StopConfd(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetTemplateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetTemplateFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetValues(ctx, req.(*types.StringList))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingPilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingPilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingPilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingPilot(ctx, req.(*types.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingFrontgate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingFrontgate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingFrontgate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingFrontgate(ctx, req.(*types.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingDrone(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingMetadataBackend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingMetadataBackend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingMetadataBackend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingMetadataBackend(ctx, req.(*types.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.RunCommandOnDroneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).RunCommand(ctx, req.(*types.RunCommandOnDroneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DroneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metadata.drone.DroneService",
	HandlerType: (*DroneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeDrone",
			Handler:    _DroneService_DistributeDrone_Handler,
		},
		{
			MethodName: "GetPilotVersion",
			Handler:    _DroneService_GetPilotVersion_Handler,
		},
		{
			MethodName: "GetFrontgateVersion",
			Handler:    _DroneService_GetFrontgateVersion_Handler,
		},
		{
			MethodName: "GetDroneVersion",
			Handler:    _DroneService_GetDroneVersion_Handler,
		},
		{
			MethodName: "GetDroneConfig",
			Handler:    _DroneService_GetDroneConfig_Handler,
		},
		{
			MethodName: "SetDroneConfig",
			Handler:    _DroneService_SetDroneConfig_Handler,
		},
		{
			MethodName: "GetConfdConfig",
			Handler:    _DroneService_GetConfdConfig_Handler,
		},
		{
			MethodName: "SetConfdConfig",
			Handler:    _DroneService_SetConfdConfig_Handler,
		},
		{
			MethodName: "GetFrontgateConfig",
			Handler:    _DroneService_GetFrontgateConfig_Handler,
		},
		{
			MethodName: "SetFrontgateConfig",
			Handler:    _DroneService_SetFrontgateConfig_Handler,
		},
		{
			MethodName: "IsConfdRunning",
			Handler:    _DroneService_IsConfdRunning_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _DroneService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _DroneService_StopConfd_Handler,
		},
		{
			MethodName: "GetTemplateFiles",
			Handler:    _DroneService_GetTemplateFiles_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _DroneService_GetValues_Handler,
		},
		{
			MethodName: "PingPilot",
			Handler:    _DroneService_PingPilot_Handler,
		},
		{
			MethodName: "PingFrontgate",
			Handler:    _DroneService_PingFrontgate_Handler,
		},
		{
			MethodName: "PingDrone",
			Handler:    _DroneService_PingDrone_Handler,
		},
		{
			MethodName: "PingMetadataBackend",
			Handler:    _DroneService_PingMetadataBackend_Handler,
		},
		{
			MethodName: "RunCommand",
			Handler:    _DroneService_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata/drone/drone.proto",
}

func init() { proto.RegisterFile("metadata/drone/drone.proto", fileDescriptor_drone_62968676582979e7) }

var fileDescriptor_drone_62968676582979e7 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x6f, 0x48, 0x19, 0x20, 0x45, 0x5b, 0xfe, 0x08, 0x23, 0x40, 0xdc, 0x90, 0x10, 0x89,
	0x04, 0xa7, 0x82, 0x38, 0xe0, 0xc6, 0x8d, 0x22, 0xa5, 0xb4, 0x8a, 0x51, 0x0f, 0xdc, 0x36, 0xf1,
	0xc4, 0x5a, 0xd5, 0x9e, 0x5d, 0xd6, 0x63, 0x44, 0x3e, 0x04, 0xdf, 0x19, 0x65, 0x37, 0x4e, 0x83,
	0xeb, 0x0d, 0x6a, 0x7b, 0xb1, 0xa5, 0x79, 0x6f, 0x9e, 0x7f, 0x3b, 0xd6, 0x0e, 0x44, 0x25, 0xb2,
	0xcc, 0x24, 0xcb, 0x61, 0x66, 0x35, 0xa1, 0x7f, 0x0e, 0x8c, 0xd5, 0xac, 0x45, 0xbf, 0xd1, 0x06,
	0xae, 0x1a, 0x5d, 0x79, 0x79, 0x65, 0xb0, 0xf2, 0x4f, 0xef, 0xbd, 0xa6, 0x2d, 0x34, 0x2d, 0xb3,
	0x80, 0xb6, 0xf3, 0x8d, 0xe8, 0x55, 0x4b, 0x5b, 0x5a, 0x4d, 0x9c, 0x4b, 0xde, 0xe8, 0x1f, 0xfe,
	0xdc, 0x87, 0x07, 0xa3, 0xb5, 0x3f, 0x45, 0xfb, 0x4b, 0x2d, 0x50, 0x24, 0x70, 0x30, 0x52, 0x15,
	0x5b, 0x35, 0xaf, 0x19, 0x9d, 0x22, 0x5e, 0x0c, 0xb6, 0xa0, 0x1e, 0xc9, 0x95, 0x63, 0x45, 0xd2,
	0xae, 0xa2, 0x27, 0x6d, 0x31, 0x29, 0x0d, 0xaf, 0xc4, 0x57, 0x38, 0x18, 0x23, 0x9f, 0xab, 0x42,
	0xf3, 0x05, 0xda, 0x4a, 0x69, 0x12, 0xdd, 0xce, 0xe8, 0x59, 0xbb, 0xdc, 0xf8, 0x13, 0x38, 0x1c,
	0x23, 0x9f, 0x34, 0xc0, 0xb7, 0x8d, 0x99, 0x38, 0x12, 0x87, 0xdc, 0x94, 0x5e, 0x76, 0x1e, 0x28,
	0xa1, 0xcc, 0x68, 0x45, 0x1c, 0x8e, 0x1a, 0x41, 0xbf, 0x89, 0x3a, 0xd6, 0xb4, 0x54, 0x79, 0x08,
	0xa6, 0x7b, 0x62, 0x9b, 0x9e, 0x11, 0xf4, 0xd3, 0x7f, 0x53, 0xf6, 0xd9, 0x43, 0x03, 0xf6, 0x2c,
	0x6b, 0x4f, 0x76, 0x43, 0x96, 0xdd, 0x1e, 0xcf, 0xb2, 0x5b, 0xd9, 0x67, 0x0f, 0xb1, 0x4c, 0x41,
	0xec, 0xfe, 0xa9, 0xfd, 0x3c, 0xaf, 0xdb, 0xe5, 0x76, 0xdf, 0x14, 0x44, 0x7a, 0x3d, 0xed, 0x7f,
	0x6d, 0x21, 0xb6, 0x2f, 0xd0, 0x9f, 0x54, 0xee, 0x0c, 0xb3, 0x9a, 0x48, 0x51, 0x90, 0xeb, 0x71,
	0xbb, 0x1c, 0x6b, 0x5d, 0x88, 0x4f, 0x00, 0x29, 0x4b, 0xeb, 0x47, 0x14, 0x6a, 0x0d, 0x7c, 0xfa,
	0x08, 0x7a, 0x29, 0x6b, 0x73, 0x9b, 0xd6, 0x04, 0x1e, 0x8d, 0x91, 0xbf, 0x63, 0x69, 0x0a, 0xc9,
	0x78, 0xa2, 0x0a, 0xac, 0x42, 0x09, 0x51, 0xbb, 0x9c, 0xb2, 0x55, 0x94, 0x4f, 0x55, 0xc5, 0x22,
	0x86, 0xde, 0x18, 0xf9, 0x42, 0x16, 0x35, 0x56, 0x62, 0x8f, 0x31, 0x7a, 0xde, 0xad, 0x9d, 0x4a,
	0x23, 0x12, 0xe8, 0x9d, 0x2b, 0xca, 0xdd, 0x55, 0x16, 0x6f, 0x82, 0x7f, 0x61, 0x7b, 0x7b, 0x02,
	0x27, 0x9a, 0xc0, 0xc3, 0x75, 0xcc, 0xd6, 0x7f, 0x87, 0xa8, 0x23, 0x4f, 0xe4, 0x97, 0xd3, 0xcd,
	0xe6, 0x7a, 0x06, 0x87, 0xeb, 0xd6, 0xd3, 0x8d, 0x16, 0xcb, 0xc5, 0x25, 0x52, 0x76, 0x07, 0x96,
	0x6f, 0x00, 0xb3, 0x9a, 0x8e, 0x75, 0x59, 0x4a, 0xca, 0xc4, 0xdb, 0xb6, 0xe9, 0x4a, 0x3b, 0x23,
	0xc7, 0x3b, 0xc3, 0x9f, 0x35, 0x56, 0x1c, 0x3d, 0xed, 0x1e, 0x78, 0xfc, 0xfe, 0xc7, 0x3b, 0x6d,
	0x90, 0x8c, 0x62, 0xab, 0x7e, 0x0f, 0x94, 0x1e, 0x6e, 0xf7, 0xb7, 0xb9, 0xcc, 0x87, 0x66, 0xee,
	0x97, 0xfb, 0x67, 0x33, 0x77, 0xef, 0xf9, 0x3d, 0xb7, 0xc5, 0x3f, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x7b, 0xd3, 0x46, 0x67, 0x06, 0x00, 0x00,
}
