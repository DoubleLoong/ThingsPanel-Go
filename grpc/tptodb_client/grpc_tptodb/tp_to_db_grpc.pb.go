// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: tp_to_db.proto

// option java_multiple_files = true;
// option java_package = "io.grpc.examples.helloworld";
// option java_outer_classname = "HelloWorldProto";

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Greeter_SayHello_FullMethodName = "/tptodb.Greeter/SayHello"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tptodb.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tp_to_db.proto",
}

const (
	ThingsPanel_GetDeviceHistory_FullMethodName               = "/tptodb.ThingsPanel/GetDeviceHistory"
	ThingsPanel_GetDeviceAttributesHistory_FullMethodName     = "/tptodb.ThingsPanel/GetDeviceAttributesHistory"
	ThingsPanel_GetDeviceAttributesCurrents_FullMethodName    = "/tptodb.ThingsPanel/GetDeviceAttributesCurrents"
	ThingsPanel_GetDeviceKVDataWithNoAggregate_FullMethodName = "/tptodb.ThingsPanel/GetDeviceKVDataWithNoAggregate"
	ThingsPanel_GetDeviceKVDataWithAggregate_FullMethodName   = "/tptodb.ThingsPanel/GetDeviceKVDataWithAggregate"
)

// ThingsPanelClient is the client API for ThingsPanel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThingsPanelClient interface {
	// 获取设备属性历史数据
	GetDeviceHistory(ctx context.Context, in *GetDeviceHistoryRequest, opts ...grpc.CallOption) (*GetDeviceHistoryReply, error)
	GetDeviceAttributesHistory(ctx context.Context, in *GetDeviceAttributesHistoryRequest, opts ...grpc.CallOption) (*GetDeviceAttributesHistoryReply, error)
	GetDeviceAttributesCurrents(ctx context.Context, in *GetDeviceAttributesCurrentsRequest, opts ...grpc.CallOption) (*GetDeviceAttributesCurrentsReply, error)
	GetDeviceKVDataWithNoAggregate(ctx context.Context, in *GetDeviceKVDataWithNoAggregateRequest, opts ...grpc.CallOption) (*GetDeviceKVDataWithNoAggregateReply, error)
	GetDeviceKVDataWithAggregate(ctx context.Context, in *GetDeviceKVDataWithAggregateRequest, opts ...grpc.CallOption) (*GetDeviceKVDataWithAggregateReply, error)
}

type thingsPanelClient struct {
	cc grpc.ClientConnInterface
}

func NewThingsPanelClient(cc grpc.ClientConnInterface) ThingsPanelClient {
	return &thingsPanelClient{cc}
}

func (c *thingsPanelClient) GetDeviceHistory(ctx context.Context, in *GetDeviceHistoryRequest, opts ...grpc.CallOption) (*GetDeviceHistoryReply, error) {
	out := new(GetDeviceHistoryReply)
	err := c.cc.Invoke(ctx, ThingsPanel_GetDeviceHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsPanelClient) GetDeviceAttributesHistory(ctx context.Context, in *GetDeviceAttributesHistoryRequest, opts ...grpc.CallOption) (*GetDeviceAttributesHistoryReply, error) {
	out := new(GetDeviceAttributesHistoryReply)
	err := c.cc.Invoke(ctx, ThingsPanel_GetDeviceAttributesHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsPanelClient) GetDeviceAttributesCurrents(ctx context.Context, in *GetDeviceAttributesCurrentsRequest, opts ...grpc.CallOption) (*GetDeviceAttributesCurrentsReply, error) {
	out := new(GetDeviceAttributesCurrentsReply)
	err := c.cc.Invoke(ctx, ThingsPanel_GetDeviceAttributesCurrents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsPanelClient) GetDeviceKVDataWithNoAggregate(ctx context.Context, in *GetDeviceKVDataWithNoAggregateRequest, opts ...grpc.CallOption) (*GetDeviceKVDataWithNoAggregateReply, error) {
	out := new(GetDeviceKVDataWithNoAggregateReply)
	err := c.cc.Invoke(ctx, ThingsPanel_GetDeviceKVDataWithNoAggregate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsPanelClient) GetDeviceKVDataWithAggregate(ctx context.Context, in *GetDeviceKVDataWithAggregateRequest, opts ...grpc.CallOption) (*GetDeviceKVDataWithAggregateReply, error) {
	out := new(GetDeviceKVDataWithAggregateReply)
	err := c.cc.Invoke(ctx, ThingsPanel_GetDeviceKVDataWithAggregate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThingsPanelServer is the server API for ThingsPanel service.
// All implementations must embed UnimplementedThingsPanelServer
// for forward compatibility
type ThingsPanelServer interface {
	// 获取设备属性历史数据
	GetDeviceHistory(context.Context, *GetDeviceHistoryRequest) (*GetDeviceHistoryReply, error)
	GetDeviceAttributesHistory(context.Context, *GetDeviceAttributesHistoryRequest) (*GetDeviceAttributesHistoryReply, error)
	GetDeviceAttributesCurrents(context.Context, *GetDeviceAttributesCurrentsRequest) (*GetDeviceAttributesCurrentsReply, error)
	GetDeviceKVDataWithNoAggregate(context.Context, *GetDeviceKVDataWithNoAggregateRequest) (*GetDeviceKVDataWithNoAggregateReply, error)
	GetDeviceKVDataWithAggregate(context.Context, *GetDeviceKVDataWithAggregateRequest) (*GetDeviceKVDataWithAggregateReply, error)
	mustEmbedUnimplementedThingsPanelServer()
}

// UnimplementedThingsPanelServer must be embedded to have forward compatible implementations.
type UnimplementedThingsPanelServer struct {
}

func (UnimplementedThingsPanelServer) GetDeviceHistory(context.Context, *GetDeviceHistoryRequest) (*GetDeviceHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceHistory not implemented")
}
func (UnimplementedThingsPanelServer) GetDeviceAttributesHistory(context.Context, *GetDeviceAttributesHistoryRequest) (*GetDeviceAttributesHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceAttributesHistory not implemented")
}
func (UnimplementedThingsPanelServer) GetDeviceAttributesCurrents(context.Context, *GetDeviceAttributesCurrentsRequest) (*GetDeviceAttributesCurrentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceAttributesCurrents not implemented")
}
func (UnimplementedThingsPanelServer) GetDeviceKVDataWithNoAggregate(context.Context, *GetDeviceKVDataWithNoAggregateRequest) (*GetDeviceKVDataWithNoAggregateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceKVDataWithNoAggregate not implemented")
}
func (UnimplementedThingsPanelServer) GetDeviceKVDataWithAggregate(context.Context, *GetDeviceKVDataWithAggregateRequest) (*GetDeviceKVDataWithAggregateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceKVDataWithAggregate not implemented")
}
func (UnimplementedThingsPanelServer) mustEmbedUnimplementedThingsPanelServer() {}

// UnsafeThingsPanelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThingsPanelServer will
// result in compilation errors.
type UnsafeThingsPanelServer interface {
	mustEmbedUnimplementedThingsPanelServer()
}

func RegisterThingsPanelServer(s grpc.ServiceRegistrar, srv ThingsPanelServer) {
	s.RegisterService(&ThingsPanel_ServiceDesc, srv)
}

func _ThingsPanel_GetDeviceHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsPanelServer).GetDeviceHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThingsPanel_GetDeviceHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsPanelServer).GetDeviceHistory(ctx, req.(*GetDeviceHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingsPanel_GetDeviceAttributesHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceAttributesHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsPanelServer).GetDeviceAttributesHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThingsPanel_GetDeviceAttributesHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsPanelServer).GetDeviceAttributesHistory(ctx, req.(*GetDeviceAttributesHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingsPanel_GetDeviceAttributesCurrents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceAttributesCurrentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsPanelServer).GetDeviceAttributesCurrents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThingsPanel_GetDeviceAttributesCurrents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsPanelServer).GetDeviceAttributesCurrents(ctx, req.(*GetDeviceAttributesCurrentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingsPanel_GetDeviceKVDataWithNoAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceKVDataWithNoAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsPanelServer).GetDeviceKVDataWithNoAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThingsPanel_GetDeviceKVDataWithNoAggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsPanelServer).GetDeviceKVDataWithNoAggregate(ctx, req.(*GetDeviceKVDataWithNoAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingsPanel_GetDeviceKVDataWithAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceKVDataWithAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsPanelServer).GetDeviceKVDataWithAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThingsPanel_GetDeviceKVDataWithAggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsPanelServer).GetDeviceKVDataWithAggregate(ctx, req.(*GetDeviceKVDataWithAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThingsPanel_ServiceDesc is the grpc.ServiceDesc for ThingsPanel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThingsPanel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tptodb.ThingsPanel",
	HandlerType: (*ThingsPanelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceHistory",
			Handler:    _ThingsPanel_GetDeviceHistory_Handler,
		},
		{
			MethodName: "GetDeviceAttributesHistory",
			Handler:    _ThingsPanel_GetDeviceAttributesHistory_Handler,
		},
		{
			MethodName: "GetDeviceAttributesCurrents",
			Handler:    _ThingsPanel_GetDeviceAttributesCurrents_Handler,
		},
		{
			MethodName: "GetDeviceKVDataWithNoAggregate",
			Handler:    _ThingsPanel_GetDeviceKVDataWithNoAggregate_Handler,
		},
		{
			MethodName: "GetDeviceKVDataWithAggregate",
			Handler:    _ThingsPanel_GetDeviceKVDataWithAggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tp_to_db.proto",
}
