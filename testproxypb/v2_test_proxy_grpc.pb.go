// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v2_test_proxy.proto

package testproxypb

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

// CloudBigtableV2TestProxyClient is the client API for CloudBigtableV2TestProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudBigtableV2TestProxyClient interface {
	// Client management:
	//
	// Creates a client in the proxy.
	// Each client has its own dedicated channel(s), and can be used concurrently
	// and independently with other clients.
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
	// Removes a client in the proxy.
	RemoveClient(ctx context.Context, in *RemoveClientRequest, opts ...grpc.CallOption) (*RemoveClientResponse, error)
	// Bigtable operations: for each operation, you should use the synchronous or
	// asynchronous variant of the client method based on the `use_async_method`
	// setting of the client instance. For starters, you can choose to implement
	// one variant, and return UNIMPLEMENTED status for the other.
	//
	// Reads a row with the client instance.
	// The result row may not be present in the response.
	// Callers should check for it (e.g. calling has_row() in C++).
	ReadRow(ctx context.Context, in *ReadRowRequest, opts ...grpc.CallOption) (*RowResult, error)
	// Reads rows with the client instance.
	ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (*RowsResult, error)
	// Writes a row with the client instance.
	MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*MutateRowResult, error)
	// Writes multiple rows with the client instance.
	BulkMutateRows(ctx context.Context, in *MutateRowsRequest, opts ...grpc.CallOption) (*MutateRowsResult, error)
	// Performs a check-and-mutate-row operation with the client instance.
	CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResult, error)
	// Obtains a row key sampling with the client instance.
	SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (*SampleRowKeysResult, error)
	// Performs a read-modify-write operation with the client.
	ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*RowResult, error)
}

type cloudBigtableV2TestProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudBigtableV2TestProxyClient(cc grpc.ClientConnInterface) CloudBigtableV2TestProxyClient {
	return &cloudBigtableV2TestProxyClient{cc}
}

func (c *cloudBigtableV2TestProxyClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) RemoveClient(ctx context.Context, in *RemoveClientRequest, opts ...grpc.CallOption) (*RemoveClientResponse, error) {
	out := new(RemoveClientResponse)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/RemoveClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) ReadRow(ctx context.Context, in *ReadRowRequest, opts ...grpc.CallOption) (*RowResult, error) {
	out := new(RowResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/ReadRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (*RowsResult, error) {
	out := new(RowsResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/ReadRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*MutateRowResult, error) {
	out := new(MutateRowResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/MutateRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) BulkMutateRows(ctx context.Context, in *MutateRowsRequest, opts ...grpc.CallOption) (*MutateRowsResult, error) {
	out := new(MutateRowsResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/BulkMutateRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResult, error) {
	out := new(CheckAndMutateRowResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/CheckAndMutateRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (*SampleRowKeysResult, error) {
	out := new(SampleRowKeysResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/SampleRowKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBigtableV2TestProxyClient) ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*RowResult, error) {
	out := new(RowResult)
	err := c.cc.Invoke(ctx, "/bigtable.client.test.CloudBigtableV2TestProxy/ReadModifyWriteRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudBigtableV2TestProxyServer is the server API for CloudBigtableV2TestProxy service.
// All implementations must embed UnimplementedCloudBigtableV2TestProxyServer
// for forward compatibility
type CloudBigtableV2TestProxyServer interface {
	// Client management:
	//
	// Creates a client in the proxy.
	// Each client has its own dedicated channel(s), and can be used concurrently
	// and independently with other clients.
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	// Removes a client in the proxy.
	RemoveClient(context.Context, *RemoveClientRequest) (*RemoveClientResponse, error)
	// Bigtable operations: for each operation, you should use the synchronous or
	// asynchronous variant of the client method based on the `use_async_method`
	// setting of the client instance. For starters, you can choose to implement
	// one variant, and return UNIMPLEMENTED status for the other.
	//
	// Reads a row with the client instance.
	// The result row may not be present in the response.
	// Callers should check for it (e.g. calling has_row() in C++).
	ReadRow(context.Context, *ReadRowRequest) (*RowResult, error)
	// Reads rows with the client instance.
	ReadRows(context.Context, *ReadRowsRequest) (*RowsResult, error)
	// Writes a row with the client instance.
	MutateRow(context.Context, *MutateRowRequest) (*MutateRowResult, error)
	// Writes multiple rows with the client instance.
	BulkMutateRows(context.Context, *MutateRowsRequest) (*MutateRowsResult, error)
	// Performs a check-and-mutate-row operation with the client instance.
	CheckAndMutateRow(context.Context, *CheckAndMutateRowRequest) (*CheckAndMutateRowResult, error)
	// Obtains a row key sampling with the client instance.
	SampleRowKeys(context.Context, *SampleRowKeysRequest) (*SampleRowKeysResult, error)
	// Performs a read-modify-write operation with the client.
	ReadModifyWriteRow(context.Context, *ReadModifyWriteRowRequest) (*RowResult, error)
	mustEmbedUnimplementedCloudBigtableV2TestProxyServer()
}

// UnimplementedCloudBigtableV2TestProxyServer must be embedded to have forward compatible implementations.
type UnimplementedCloudBigtableV2TestProxyServer struct {
}

func (UnimplementedCloudBigtableV2TestProxyServer) CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) RemoveClient(context.Context, *RemoveClientRequest) (*RemoveClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClient not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) ReadRow(context.Context, *ReadRowRequest) (*RowResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRow not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) ReadRows(context.Context, *ReadRowsRequest) (*RowsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRows not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) MutateRow(context.Context, *MutateRowRequest) (*MutateRowResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRow not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) BulkMutateRows(context.Context, *MutateRowsRequest) (*MutateRowsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkMutateRows not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) CheckAndMutateRow(context.Context, *CheckAndMutateRowRequest) (*CheckAndMutateRowResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndMutateRow not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) SampleRowKeys(context.Context, *SampleRowKeysRequest) (*SampleRowKeysResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleRowKeys not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) ReadModifyWriteRow(context.Context, *ReadModifyWriteRowRequest) (*RowResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadModifyWriteRow not implemented")
}
func (UnimplementedCloudBigtableV2TestProxyServer) mustEmbedUnimplementedCloudBigtableV2TestProxyServer() {
}

// UnsafeCloudBigtableV2TestProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudBigtableV2TestProxyServer will
// result in compilation errors.
type UnsafeCloudBigtableV2TestProxyServer interface {
	mustEmbedUnimplementedCloudBigtableV2TestProxyServer()
}

func RegisterCloudBigtableV2TestProxyServer(s grpc.ServiceRegistrar, srv CloudBigtableV2TestProxyServer) {
	s.RegisterService(&CloudBigtableV2TestProxy_ServiceDesc, srv)
}

func _CloudBigtableV2TestProxy_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_RemoveClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).RemoveClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/RemoveClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).RemoveClient(ctx, req.(*RemoveClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_ReadRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).ReadRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/ReadRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).ReadRow(ctx, req.(*ReadRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_ReadRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).ReadRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/ReadRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).ReadRows(ctx, req.(*ReadRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_MutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).MutateRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/MutateRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).MutateRow(ctx, req.(*MutateRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_BulkMutateRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).BulkMutateRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/BulkMutateRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).BulkMutateRows(ctx, req.(*MutateRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_CheckAndMutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndMutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).CheckAndMutateRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/CheckAndMutateRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).CheckAndMutateRow(ctx, req.(*CheckAndMutateRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_SampleRowKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRowKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).SampleRowKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/SampleRowKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).SampleRowKeys(ctx, req.(*SampleRowKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBigtableV2TestProxy_ReadModifyWriteRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadModifyWriteRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBigtableV2TestProxyServer).ReadModifyWriteRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigtable.client.test.CloudBigtableV2TestProxy/ReadModifyWriteRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBigtableV2TestProxyServer).ReadModifyWriteRow(ctx, req.(*ReadModifyWriteRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudBigtableV2TestProxy_ServiceDesc is the grpc.ServiceDesc for CloudBigtableV2TestProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudBigtableV2TestProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bigtable.client.test.CloudBigtableV2TestProxy",
	HandlerType: (*CloudBigtableV2TestProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _CloudBigtableV2TestProxy_CreateClient_Handler,
		},
		{
			MethodName: "RemoveClient",
			Handler:    _CloudBigtableV2TestProxy_RemoveClient_Handler,
		},
		{
			MethodName: "ReadRow",
			Handler:    _CloudBigtableV2TestProxy_ReadRow_Handler,
		},
		{
			MethodName: "ReadRows",
			Handler:    _CloudBigtableV2TestProxy_ReadRows_Handler,
		},
		{
			MethodName: "MutateRow",
			Handler:    _CloudBigtableV2TestProxy_MutateRow_Handler,
		},
		{
			MethodName: "BulkMutateRows",
			Handler:    _CloudBigtableV2TestProxy_BulkMutateRows_Handler,
		},
		{
			MethodName: "CheckAndMutateRow",
			Handler:    _CloudBigtableV2TestProxy_CheckAndMutateRow_Handler,
		},
		{
			MethodName: "SampleRowKeys",
			Handler:    _CloudBigtableV2TestProxy_SampleRowKeys_Handler,
		},
		{
			MethodName: "ReadModifyWriteRow",
			Handler:    _CloudBigtableV2TestProxy_ReadModifyWriteRow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2_test_proxy.proto",
}
