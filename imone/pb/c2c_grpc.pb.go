// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// IMServiceClient is the client API for IMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IMServiceClient interface {
	SendMsg(ctx context.Context, in *C2CSendRequest, opts ...grpc.CallOption) (*C2CSendResponse, error)
	PushMsg(ctx context.Context, in *C2CPushRequest, opts ...grpc.CallOption) (*C2CPushResponse, error)
	QueryUserOnline(ctx context.Context, in *QueryUsersOnlineReq, opts ...grpc.CallOption) (*QueryUsersOnlineRsp, error)
}

type iMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIMServiceClient(cc grpc.ClientConnInterface) IMServiceClient {
	return &iMServiceClient{cc}
}

func (c *iMServiceClient) SendMsg(ctx context.Context, in *C2CSendRequest, opts ...grpc.CallOption) (*C2CSendResponse, error) {
	out := new(C2CSendResponse)
	err := c.cc.Invoke(ctx, "/pb.IMService/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) PushMsg(ctx context.Context, in *C2CPushRequest, opts ...grpc.CallOption) (*C2CPushResponse, error) {
	out := new(C2CPushResponse)
	err := c.cc.Invoke(ctx, "/pb.IMService/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) QueryUserOnline(ctx context.Context, in *QueryUsersOnlineReq, opts ...grpc.CallOption) (*QueryUsersOnlineRsp, error) {
	out := new(QueryUsersOnlineRsp)
	err := c.cc.Invoke(ctx, "/pb.IMService/QueryUserOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMServiceServer is the server API for IMService service.
// All implementations must embed UnimplementedIMServiceServer
// for forward compatibility
type IMServiceServer interface {
	SendMsg(context.Context, *C2CSendRequest) (*C2CSendResponse, error)
	PushMsg(context.Context, *C2CPushRequest) (*C2CPushResponse, error)
	QueryUserOnline(context.Context, *QueryUsersOnlineReq) (*QueryUsersOnlineRsp, error)
	mustEmbedUnimplementedIMServiceServer()
}

// UnimplementedIMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIMServiceServer struct {
}

func (UnimplementedIMServiceServer) SendMsg(context.Context, *C2CSendRequest) (*C2CSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedIMServiceServer) PushMsg(context.Context, *C2CPushRequest) (*C2CPushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (UnimplementedIMServiceServer) QueryUserOnline(context.Context, *QueryUsersOnlineReq) (*QueryUsersOnlineRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserOnline not implemented")
}
func (UnimplementedIMServiceServer) mustEmbedUnimplementedIMServiceServer() {}

// UnsafeIMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IMServiceServer will
// result in compilation errors.
type UnsafeIMServiceServer interface {
	mustEmbedUnimplementedIMServiceServer()
}

func RegisterIMServiceServer(s grpc.ServiceRegistrar, srv IMServiceServer) {
	s.RegisterService(&IMService_ServiceDesc, srv)
}

func _IMService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2CSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IMService/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).SendMsg(ctx, req.(*C2CSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2CPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IMService/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).PushMsg(ctx, req.(*C2CPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_QueryUserOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUsersOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).QueryUserOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IMService/QueryUserOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).QueryUserOnline(ctx, req.(*QueryUsersOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IMService_ServiceDesc is the grpc.ServiceDesc for IMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IMService",
	HandlerType: (*IMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _IMService_SendMsg_Handler,
		},
		{
			MethodName: "PushMsg",
			Handler:    _IMService_PushMsg_Handler,
		},
		{
			MethodName: "QueryUserOnline",
			Handler:    _IMService_QueryUserOnline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "c2c.proto",
}
