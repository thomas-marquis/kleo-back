// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: transaction.proto

package generated

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

// KeloAppServiceClient is the client API for KeloAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeloAppServiceClient interface {
	ListTransactionsByUser(ctx context.Context, in *ListTransactionsByUserRequest, opts ...grpc.CallOption) (*TransactionsListResponse, error)
	GetUserById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type keloAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeloAppServiceClient(cc grpc.ClientConnInterface) KeloAppServiceClient {
	return &keloAppServiceClient{cc}
}

func (c *keloAppServiceClient) ListTransactionsByUser(ctx context.Context, in *ListTransactionsByUserRequest, opts ...grpc.CallOption) (*TransactionsListResponse, error) {
	out := new(TransactionsListResponse)
	err := c.cc.Invoke(ctx, "/KeloAppService/ListTransactionsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keloAppServiceClient) GetUserById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/KeloAppService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeloAppServiceServer is the server API for KeloAppService service.
// All implementations must embed UnimplementedKeloAppServiceServer
// for forward compatibility
type KeloAppServiceServer interface {
	ListTransactionsByUser(context.Context, *ListTransactionsByUserRequest) (*TransactionsListResponse, error)
	GetUserById(context.Context, *GetByIdRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedKeloAppServiceServer()
}

// UnimplementedKeloAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeloAppServiceServer struct {
}

func (UnimplementedKeloAppServiceServer) ListTransactionsByUser(context.Context, *ListTransactionsByUserRequest) (*TransactionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsByUser not implemented")
}
func (UnimplementedKeloAppServiceServer) GetUserById(context.Context, *GetByIdRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedKeloAppServiceServer) mustEmbedUnimplementedKeloAppServiceServer() {}

// UnsafeKeloAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeloAppServiceServer will
// result in compilation errors.
type UnsafeKeloAppServiceServer interface {
	mustEmbedUnimplementedKeloAppServiceServer()
}

func RegisterKeloAppServiceServer(s grpc.ServiceRegistrar, srv KeloAppServiceServer) {
	s.RegisterService(&KeloAppService_ServiceDesc, srv)
}

func _KeloAppService_ListTransactionsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeloAppServiceServer).ListTransactionsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeloAppService/ListTransactionsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeloAppServiceServer).ListTransactionsByUser(ctx, req.(*ListTransactionsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeloAppService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeloAppServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeloAppService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeloAppServiceServer).GetUserById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeloAppService_ServiceDesc is the grpc.ServiceDesc for KeloAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeloAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KeloAppService",
	HandlerType: (*KeloAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransactionsByUser",
			Handler:    _KeloAppService_ListTransactionsByUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _KeloAppService_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
