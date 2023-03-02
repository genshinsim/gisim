// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/backend/db.proto

package db

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

// DBStoreClient is the client API for DBStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBStoreClient interface {
	// generic get for pulling from approved db
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetUnfiltered(ctx context.Context, in *GetUnfilteredRequest, opts ...grpc.CallOption) (*GetUnfilteredResponse, error)
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkResponse, error)
	// tagging
	ApproveTag(ctx context.Context, in *ApproveTagRequest, opts ...grpc.CallOption) (*ApproveTagResponse, error)
	RejectTag(ctx context.Context, in *RejectTagRequest, opts ...grpc.CallOption) (*RejectTagResponse, error)
}

type dBStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewDBStoreClient(cc grpc.ClientConnInterface) DBStoreClient {
	return &dBStoreClient{cc}
}

func (c *dBStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetUnfiltered(ctx context.Context, in *GetUnfilteredRequest, opts ...grpc.CallOption) (*GetUnfilteredResponse, error) {
	out := new(GetUnfilteredResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/GetUnfiltered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkResponse, error) {
	out := new(GetWorkResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/GetWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) ApproveTag(ctx context.Context, in *ApproveTagRequest, opts ...grpc.CallOption) (*ApproveTagResponse, error) {
	out := new(ApproveTagResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/ApproveTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) RejectTag(ctx context.Context, in *RejectTagRequest, opts ...grpc.CallOption) (*RejectTagResponse, error) {
	out := new(RejectTagResponse)
	err := c.cc.Invoke(ctx, "/db.DBStore/RejectTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBStoreServer is the server API for DBStore service.
// All implementations must embed UnimplementedDBStoreServer
// for forward compatibility
type DBStoreServer interface {
	// generic get for pulling from approved db
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetUnfiltered(context.Context, *GetUnfilteredRequest) (*GetUnfilteredResponse, error)
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetWork(context.Context, *GetWorkRequest) (*GetWorkResponse, error)
	// tagging
	ApproveTag(context.Context, *ApproveTagRequest) (*ApproveTagResponse, error)
	RejectTag(context.Context, *RejectTagRequest) (*RejectTagResponse, error)
	mustEmbedUnimplementedDBStoreServer()
}

// UnimplementedDBStoreServer must be embedded to have forward compatible implementations.
type UnimplementedDBStoreServer struct {
}

func (UnimplementedDBStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDBStoreServer) GetUnfiltered(context.Context, *GetUnfilteredRequest) (*GetUnfilteredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnfiltered not implemented")
}
func (UnimplementedDBStoreServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedDBStoreServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDBStoreServer) GetWork(context.Context, *GetWorkRequest) (*GetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWork not implemented")
}
func (UnimplementedDBStoreServer) ApproveTag(context.Context, *ApproveTagRequest) (*ApproveTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTag not implemented")
}
func (UnimplementedDBStoreServer) RejectTag(context.Context, *RejectTagRequest) (*RejectTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTag not implemented")
}
func (UnimplementedDBStoreServer) mustEmbedUnimplementedDBStoreServer() {}

// UnsafeDBStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBStoreServer will
// result in compilation errors.
type UnsafeDBStoreServer interface {
	mustEmbedUnimplementedDBStoreServer()
}

func RegisterDBStoreServer(s grpc.ServiceRegistrar, srv DBStoreServer) {
	s.RegisterService(&DBStore_ServiceDesc, srv)
}

func _DBStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetUnfiltered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnfilteredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetUnfiltered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/GetUnfiltered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetUnfiltered(ctx, req.(*GetUnfilteredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/GetWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetWork(ctx, req.(*GetWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_ApproveTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).ApproveTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/ApproveTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).ApproveTag(ctx, req.(*ApproveTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_RejectTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).RejectTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.DBStore/RejectTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).RejectTag(ctx, req.(*RejectTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBStore_ServiceDesc is the grpc.ServiceDesc for DBStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DBStore",
	HandlerType: (*DBStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DBStore_Get_Handler,
		},
		{
			MethodName: "GetUnfiltered",
			Handler:    _DBStore_GetUnfiltered_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _DBStore_GetOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DBStore_Update_Handler,
		},
		{
			MethodName: "GetWork",
			Handler:    _DBStore_GetWork_Handler,
		},
		{
			MethodName: "ApproveTag",
			Handler:    _DBStore_ApproveTag_Handler,
		},
		{
			MethodName: "RejectTag",
			Handler:    _DBStore_RejectTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/backend/db.proto",
}
