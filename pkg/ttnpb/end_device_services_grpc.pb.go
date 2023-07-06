// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
// - protoc             v4.22.2
// source: lorawan-stack/api/end_device_services.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EndDeviceRegistry_Create_FullMethodName                = "/ttn.lorawan.v3.EndDeviceRegistry/Create"
	EndDeviceRegistry_Get_FullMethodName                   = "/ttn.lorawan.v3.EndDeviceRegistry/Get"
	EndDeviceRegistry_GetIdentifiersForEUIs_FullMethodName = "/ttn.lorawan.v3.EndDeviceRegistry/GetIdentifiersForEUIs"
	EndDeviceRegistry_List_FullMethodName                  = "/ttn.lorawan.v3.EndDeviceRegistry/List"
	EndDeviceRegistry_Update_FullMethodName                = "/ttn.lorawan.v3.EndDeviceRegistry/Update"
	EndDeviceRegistry_BatchUpdateLastSeen_FullMethodName   = "/ttn.lorawan.v3.EndDeviceRegistry/BatchUpdateLastSeen"
	EndDeviceRegistry_Delete_FullMethodName                = "/ttn.lorawan.v3.EndDeviceRegistry/Delete"
)

// EndDeviceRegistryClient is the client API for EndDeviceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceRegistryClient interface {
	// Create a new end device within an application.
	//
	// After registering an end device, it also needs to be registered in
	// the NsEndDeviceRegistry that is exposed by the Network Server,
	// the AsEndDeviceRegistry that is exposed by the Application Server,
	// and the JsEndDeviceRegistry that is exposed by the Join Server.
	Create(ctx context.Context, in *CreateEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Get the end device with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Get the identifiers of the end device that has the given EUIs registered.
	GetIdentifiersForEUIs(ctx context.Context, in *GetEndDeviceIdentifiersForEUIsRequest, opts ...grpc.CallOption) (*EndDeviceIdentifiers, error)
	// List end devices in the given application.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error)
	// Update the end device, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Update the last seen timestamp for a batch of end devices.
	BatchUpdateLastSeen(ctx context.Context, in *BatchUpdateEndDeviceLastSeenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the end device with the given IDs.
	//
	// Before deleting an end device it first needs to be deleted from the
	// NsEndDeviceRegistry, the AsEndDeviceRegistry and the JsEndDeviceRegistry.
	// In addition, if the device claimed on a Join Server, it also needs to be
	// unclaimed via the DeviceClaimingServer so it can be claimed in the future.
	// This is NOT done automatically.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type endDeviceRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceRegistryClient(cc grpc.ClientConnInterface) EndDeviceRegistryClient {
	return &endDeviceRegistryClient{cc}
}

func (c *endDeviceRegistryClient) Create(ctx context.Context, in *CreateEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) GetIdentifiersForEUIs(ctx context.Context, in *GetEndDeviceIdentifiersForEUIsRequest, opts ...grpc.CallOption) (*EndDeviceIdentifiers, error) {
	out := new(EndDeviceIdentifiers)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_GetIdentifiersForEUIs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) List(ctx context.Context, in *ListEndDevicesRequest, opts ...grpc.CallOption) (*EndDevices, error) {
	out := new(EndDevices)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) Update(ctx context.Context, in *UpdateEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) BatchUpdateLastSeen(ctx context.Context, in *BatchUpdateEndDeviceLastSeenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_BatchUpdateLastSeen_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceRegistry_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceRegistryServer is the server API for EndDeviceRegistry service.
// All implementations must embed UnimplementedEndDeviceRegistryServer
// for forward compatibility
type EndDeviceRegistryServer interface {
	// Create a new end device within an application.
	//
	// After registering an end device, it also needs to be registered in
	// the NsEndDeviceRegistry that is exposed by the Network Server,
	// the AsEndDeviceRegistry that is exposed by the Application Server,
	// and the JsEndDeviceRegistry that is exposed by the Join Server.
	Create(context.Context, *CreateEndDeviceRequest) (*EndDevice, error)
	// Get the end device with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// Get the identifiers of the end device that has the given EUIs registered.
	GetIdentifiersForEUIs(context.Context, *GetEndDeviceIdentifiersForEUIsRequest) (*EndDeviceIdentifiers, error)
	// List end devices in the given application.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListEndDevicesRequest) (*EndDevices, error)
	// Update the end device, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateEndDeviceRequest) (*EndDevice, error)
	// Update the last seen timestamp for a batch of end devices.
	BatchUpdateLastSeen(context.Context, *BatchUpdateEndDeviceLastSeenRequest) (*emptypb.Empty, error)
	// Delete the end device with the given IDs.
	//
	// Before deleting an end device it first needs to be deleted from the
	// NsEndDeviceRegistry, the AsEndDeviceRegistry and the JsEndDeviceRegistry.
	// In addition, if the device claimed on a Join Server, it also needs to be
	// unclaimed via the DeviceClaimingServer so it can be claimed in the future.
	// This is NOT done automatically.
	Delete(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedEndDeviceRegistryServer()
}

// UnimplementedEndDeviceRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceRegistryServer struct {
}

func (UnimplementedEndDeviceRegistryServer) Create(context.Context, *CreateEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEndDeviceRegistryServer) Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEndDeviceRegistryServer) GetIdentifiersForEUIs(context.Context, *GetEndDeviceIdentifiersForEUIsRequest) (*EndDeviceIdentifiers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentifiersForEUIs not implemented")
}
func (UnimplementedEndDeviceRegistryServer) List(context.Context, *ListEndDevicesRequest) (*EndDevices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEndDeviceRegistryServer) Update(context.Context, *UpdateEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEndDeviceRegistryServer) BatchUpdateLastSeen(context.Context, *BatchUpdateEndDeviceLastSeenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateLastSeen not implemented")
}
func (UnimplementedEndDeviceRegistryServer) Delete(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEndDeviceRegistryServer) mustEmbedUnimplementedEndDeviceRegistryServer() {}

// UnsafeEndDeviceRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceRegistryServer will
// result in compilation errors.
type UnsafeEndDeviceRegistryServer interface {
	mustEmbedUnimplementedEndDeviceRegistryServer()
}

func RegisterEndDeviceRegistryServer(s grpc.ServiceRegistrar, srv EndDeviceRegistryServer) {
	s.RegisterService(&EndDeviceRegistry_ServiceDesc, srv)
}

func _EndDeviceRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).Create(ctx, req.(*CreateEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_GetIdentifiersForEUIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceIdentifiersForEUIsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).GetIdentifiersForEUIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_GetIdentifiersForEUIs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).GetIdentifiersForEUIs(ctx, req.(*GetEndDeviceIdentifiersForEUIsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).List(ctx, req.(*ListEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).Update(ctx, req.(*UpdateEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_BatchUpdateLastSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateEndDeviceLastSeenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).BatchUpdateLastSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_BatchUpdateLastSeen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).BatchUpdateLastSeen(ctx, req.(*BatchUpdateEndDeviceLastSeenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceRegistry_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceRegistry_ServiceDesc is the grpc.ServiceDesc for EndDeviceRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceRegistry",
	HandlerType: (*EndDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EndDeviceRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EndDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "GetIdentifiersForEUIs",
			Handler:    _EndDeviceRegistry_GetIdentifiersForEUIs_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EndDeviceRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EndDeviceRegistry_Update_Handler,
		},
		{
			MethodName: "BatchUpdateLastSeen",
			Handler:    _EndDeviceRegistry_BatchUpdateLastSeen_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EndDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}

const (
	EndDeviceTemplateConverter_ListFormats_FullMethodName = "/ttn.lorawan.v3.EndDeviceTemplateConverter/ListFormats"
	EndDeviceTemplateConverter_Convert_FullMethodName     = "/ttn.lorawan.v3.EndDeviceTemplateConverter/Convert"
)

// EndDeviceTemplateConverterClient is the client API for EndDeviceTemplateConverter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceTemplateConverterClient interface {
	// Returns the configured formats to convert from.
	ListFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EndDeviceTemplateFormats, error)
	// Converts the binary data to a stream of end device templates.
	Convert(ctx context.Context, in *ConvertEndDeviceTemplateRequest, opts ...grpc.CallOption) (EndDeviceTemplateConverter_ConvertClient, error)
}

type endDeviceTemplateConverterClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceTemplateConverterClient(cc grpc.ClientConnInterface) EndDeviceTemplateConverterClient {
	return &endDeviceTemplateConverterClient{cc}
}

func (c *endDeviceTemplateConverterClient) ListFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EndDeviceTemplateFormats, error) {
	out := new(EndDeviceTemplateFormats)
	err := c.cc.Invoke(ctx, EndDeviceTemplateConverter_ListFormats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceTemplateConverterClient) Convert(ctx context.Context, in *ConvertEndDeviceTemplateRequest, opts ...grpc.CallOption) (EndDeviceTemplateConverter_ConvertClient, error) {
	stream, err := c.cc.NewStream(ctx, &EndDeviceTemplateConverter_ServiceDesc.Streams[0], EndDeviceTemplateConverter_Convert_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &endDeviceTemplateConverterConvertClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EndDeviceTemplateConverter_ConvertClient interface {
	Recv() (*EndDeviceTemplate, error)
	grpc.ClientStream
}

type endDeviceTemplateConverterConvertClient struct {
	grpc.ClientStream
}

func (x *endDeviceTemplateConverterConvertClient) Recv() (*EndDeviceTemplate, error) {
	m := new(EndDeviceTemplate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EndDeviceTemplateConverterServer is the server API for EndDeviceTemplateConverter service.
// All implementations must embed UnimplementedEndDeviceTemplateConverterServer
// for forward compatibility
type EndDeviceTemplateConverterServer interface {
	// Returns the configured formats to convert from.
	ListFormats(context.Context, *emptypb.Empty) (*EndDeviceTemplateFormats, error)
	// Converts the binary data to a stream of end device templates.
	Convert(*ConvertEndDeviceTemplateRequest, EndDeviceTemplateConverter_ConvertServer) error
	mustEmbedUnimplementedEndDeviceTemplateConverterServer()
}

// UnimplementedEndDeviceTemplateConverterServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceTemplateConverterServer struct {
}

func (UnimplementedEndDeviceTemplateConverterServer) ListFormats(context.Context, *emptypb.Empty) (*EndDeviceTemplateFormats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFormats not implemented")
}
func (UnimplementedEndDeviceTemplateConverterServer) Convert(*ConvertEndDeviceTemplateRequest, EndDeviceTemplateConverter_ConvertServer) error {
	return status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedEndDeviceTemplateConverterServer) mustEmbedUnimplementedEndDeviceTemplateConverterServer() {
}

// UnsafeEndDeviceTemplateConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceTemplateConverterServer will
// result in compilation errors.
type UnsafeEndDeviceTemplateConverterServer interface {
	mustEmbedUnimplementedEndDeviceTemplateConverterServer()
}

func RegisterEndDeviceTemplateConverterServer(s grpc.ServiceRegistrar, srv EndDeviceTemplateConverterServer) {
	s.RegisterService(&EndDeviceTemplateConverter_ServiceDesc, srv)
}

func _EndDeviceTemplateConverter_ListFormats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceTemplateConverterServer).ListFormats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceTemplateConverter_ListFormats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceTemplateConverterServer).ListFormats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceTemplateConverter_Convert_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConvertEndDeviceTemplateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EndDeviceTemplateConverterServer).Convert(m, &endDeviceTemplateConverterConvertServer{stream})
}

type EndDeviceTemplateConverter_ConvertServer interface {
	Send(*EndDeviceTemplate) error
	grpc.ServerStream
}

type endDeviceTemplateConverterConvertServer struct {
	grpc.ServerStream
}

func (x *endDeviceTemplateConverterConvertServer) Send(m *EndDeviceTemplate) error {
	return x.ServerStream.SendMsg(m)
}

// EndDeviceTemplateConverter_ServiceDesc is the grpc.ServiceDesc for EndDeviceTemplateConverter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceTemplateConverter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceTemplateConverter",
	HandlerType: (*EndDeviceTemplateConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFormats",
			Handler:    _EndDeviceTemplateConverter_ListFormats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Convert",
			Handler:       _EndDeviceTemplateConverter_Convert_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}

const (
	EndDeviceBatchRegistry_Delete_FullMethodName = "/ttn.lorawan.v3.EndDeviceBatchRegistry/Delete"
)

// EndDeviceBatchRegistryClient is the client API for EndDeviceBatchRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceBatchRegistryClient interface {
	// Delete a batch of end devices with the given IDs.
	//
	// This operation is atomic; either all devices are deleted or none.
	// Devices not found are skipped and no error is returned.
	// Before calling this RPC, use the corresponding BatchDelete RPCs
	// of NsEndDeviceRegistry, AsEndDeviceRegistry and
	// optionally the JsEndDeviceRegistry to delete the end devices.
	// If the devices were claimed on a Join Server, use the BatchUnclaim RPC
	// of the DeviceClaimingServer.
	// This is NOT done automatically.
	Delete(ctx context.Context, in *BatchDeleteEndDevicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type endDeviceBatchRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceBatchRegistryClient(cc grpc.ClientConnInterface) EndDeviceBatchRegistryClient {
	return &endDeviceBatchRegistryClient{cc}
}

func (c *endDeviceBatchRegistryClient) Delete(ctx context.Context, in *BatchDeleteEndDevicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceBatchRegistry_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceBatchRegistryServer is the server API for EndDeviceBatchRegistry service.
// All implementations must embed UnimplementedEndDeviceBatchRegistryServer
// for forward compatibility
type EndDeviceBatchRegistryServer interface {
	// Delete a batch of end devices with the given IDs.
	//
	// This operation is atomic; either all devices are deleted or none.
	// Devices not found are skipped and no error is returned.
	// Before calling this RPC, use the corresponding BatchDelete RPCs
	// of NsEndDeviceRegistry, AsEndDeviceRegistry and
	// optionally the JsEndDeviceRegistry to delete the end devices.
	// If the devices were claimed on a Join Server, use the BatchUnclaim RPC
	// of the DeviceClaimingServer.
	// This is NOT done automatically.
	Delete(context.Context, *BatchDeleteEndDevicesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEndDeviceBatchRegistryServer()
}

// UnimplementedEndDeviceBatchRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceBatchRegistryServer struct {
}

func (UnimplementedEndDeviceBatchRegistryServer) Delete(context.Context, *BatchDeleteEndDevicesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEndDeviceBatchRegistryServer) mustEmbedUnimplementedEndDeviceBatchRegistryServer() {
}

// UnsafeEndDeviceBatchRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceBatchRegistryServer will
// result in compilation errors.
type UnsafeEndDeviceBatchRegistryServer interface {
	mustEmbedUnimplementedEndDeviceBatchRegistryServer()
}

func RegisterEndDeviceBatchRegistryServer(s grpc.ServiceRegistrar, srv EndDeviceBatchRegistryServer) {
	s.RegisterService(&EndDeviceBatchRegistry_ServiceDesc, srv)
}

func _EndDeviceBatchRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceBatchRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceBatchRegistry_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceBatchRegistryServer).Delete(ctx, req.(*BatchDeleteEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceBatchRegistry_ServiceDesc is the grpc.ServiceDesc for EndDeviceBatchRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceBatchRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceBatchRegistry",
	HandlerType: (*EndDeviceBatchRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _EndDeviceBatchRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/end_device_services.proto",
}
