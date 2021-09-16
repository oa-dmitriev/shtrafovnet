// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package legalinfo

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

// LegalInfoFetcherClient is the client API for LegalInfoFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LegalInfoFetcherClient interface {
	GetInfoByInn(ctx context.Context, in *Inn, opts ...grpc.CallOption) (*Info, error)
}

type legalInfoFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewLegalInfoFetcherClient(cc grpc.ClientConnInterface) LegalInfoFetcherClient {
	return &legalInfoFetcherClient{cc}
}

func (c *legalInfoFetcherClient) GetInfoByInn(ctx context.Context, in *Inn, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/proto.LegalInfoFetcher/GetInfoByInn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LegalInfoFetcherServer is the server API for LegalInfoFetcher service.
// All implementations must embed UnimplementedLegalInfoFetcherServer
// for forward compatibility
type LegalInfoFetcherServer interface {
	GetInfoByInn(context.Context, *Inn) (*Info, error)
	mustEmbedUnimplementedLegalInfoFetcherServer()
}

// UnimplementedLegalInfoFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedLegalInfoFetcherServer struct {
}

func (UnimplementedLegalInfoFetcherServer) GetInfoByInn(context.Context, *Inn) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByInn not implemented")
}
func (UnimplementedLegalInfoFetcherServer) mustEmbedUnimplementedLegalInfoFetcherServer() {}

// UnsafeLegalInfoFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LegalInfoFetcherServer will
// result in compilation errors.
type UnsafeLegalInfoFetcherServer interface {
	mustEmbedUnimplementedLegalInfoFetcherServer()
}

func RegisterLegalInfoFetcherServer(s grpc.ServiceRegistrar, srv LegalInfoFetcherServer) {
	s.RegisterService(&LegalInfoFetcher_ServiceDesc, srv)
}

func _LegalInfoFetcher_GetInfoByInn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Inn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegalInfoFetcherServer).GetInfoByInn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LegalInfoFetcher/GetInfoByInn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegalInfoFetcherServer).GetInfoByInn(ctx, req.(*Inn))
	}
	return interceptor(ctx, in, info, handler)
}

// LegalInfoFetcher_ServiceDesc is the grpc.ServiceDesc for LegalInfoFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LegalInfoFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LegalInfoFetcher",
	HandlerType: (*LegalInfoFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfoByInn",
			Handler:    _LegalInfoFetcher_GetInfoByInn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "legalinfo.proto",
}
