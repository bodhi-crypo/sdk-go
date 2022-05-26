// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: injective_exchange_rpc.proto

package injective_exchange_rpcpb

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

// InjectiveExchangeRPCClient is the client API for InjectiveExchangeRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveExchangeRPCClient interface {
	// GetTx gets transaction details by hash.
	GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)
	// PrepareTx generates a Web3-signable body for a Cosmos transaction
	PrepareTx(ctx context.Context, in *PrepareTxRequest, opts ...grpc.CallOption) (*PrepareTxResponse, error)
	// BroadcastTx broadcasts a signed Web3 transaction
	BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error)
}

type injectiveExchangeRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveExchangeRPCClient(cc grpc.ClientConnInterface) InjectiveExchangeRPCClient {
	return &injectiveExchangeRPCClient{cc}
}

func (c *injectiveExchangeRPCClient) GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error) {
	out := new(GetTxResponse)
	err := c.cc.Invoke(ctx, "/injective_exchange_rpc.InjectiveExchangeRPC/GetTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExchangeRPCClient) PrepareTx(ctx context.Context, in *PrepareTxRequest, opts ...grpc.CallOption) (*PrepareTxResponse, error) {
	out := new(PrepareTxResponse)
	err := c.cc.Invoke(ctx, "/injective_exchange_rpc.InjectiveExchangeRPC/PrepareTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExchangeRPCClient) BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error) {
	out := new(BroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/injective_exchange_rpc.InjectiveExchangeRPC/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveExchangeRPCServer is the server API for InjectiveExchangeRPC service.
// All implementations must embed UnimplementedInjectiveExchangeRPCServer
// for forward compatibility
type InjectiveExchangeRPCServer interface {
	// GetTx gets transaction details by hash.
	GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error)
	// PrepareTx generates a Web3-signable body for a Cosmos transaction
	PrepareTx(context.Context, *PrepareTxRequest) (*PrepareTxResponse, error)
	// BroadcastTx broadcasts a signed Web3 transaction
	BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error)
	mustEmbedUnimplementedInjectiveExchangeRPCServer()
}

// UnimplementedInjectiveExchangeRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveExchangeRPCServer struct {
}

func (UnimplementedInjectiveExchangeRPCServer) GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}
func (UnimplementedInjectiveExchangeRPCServer) PrepareTx(context.Context, *PrepareTxRequest) (*PrepareTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareTx not implemented")
}
func (UnimplementedInjectiveExchangeRPCServer) BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTx not implemented")
}
func (UnimplementedInjectiveExchangeRPCServer) mustEmbedUnimplementedInjectiveExchangeRPCServer() {}

// UnsafeInjectiveExchangeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveExchangeRPCServer will
// result in compilation errors.
type UnsafeInjectiveExchangeRPCServer interface {
	mustEmbedUnimplementedInjectiveExchangeRPCServer()
}

func RegisterInjectiveExchangeRPCServer(s grpc.ServiceRegistrar, srv InjectiveExchangeRPCServer) {
	s.RegisterService(&InjectiveExchangeRPC_ServiceDesc, srv)
}

func _InjectiveExchangeRPC_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExchangeRPCServer).GetTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_exchange_rpc.InjectiveExchangeRPC/GetTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExchangeRPCServer).GetTx(ctx, req.(*GetTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExchangeRPC_PrepareTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExchangeRPCServer).PrepareTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_exchange_rpc.InjectiveExchangeRPC/PrepareTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExchangeRPCServer).PrepareTx(ctx, req.(*PrepareTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExchangeRPC_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExchangeRPCServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_exchange_rpc.InjectiveExchangeRPC/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExchangeRPCServer).BroadcastTx(ctx, req.(*BroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveExchangeRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveExchangeRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveExchangeRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_exchange_rpc.InjectiveExchangeRPC",
	HandlerType: (*InjectiveExchangeRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTx",
			Handler:    _InjectiveExchangeRPC_GetTx_Handler,
		},
		{
			MethodName: "PrepareTx",
			Handler:    _InjectiveExchangeRPC_PrepareTx_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _InjectiveExchangeRPC_BroadcastTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "injective_exchange_rpc.proto",
}
