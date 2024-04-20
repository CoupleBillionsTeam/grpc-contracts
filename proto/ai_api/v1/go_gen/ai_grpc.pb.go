// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc3
// source: ai.proto

package ai_api

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

// AiApiServiceClient is the client API for AiApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AiApiServiceClient interface {
	GetIdFromAi(ctx context.Context, in *AiApiRequest, opts ...grpc.CallOption) (*AiApiResponse, error)
}

type aiApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAiApiServiceClient(cc grpc.ClientConnInterface) AiApiServiceClient {
	return &aiApiServiceClient{cc}
}

func (c *aiApiServiceClient) GetIdFromAi(ctx context.Context, in *AiApiRequest, opts ...grpc.CallOption) (*AiApiResponse, error) {
	out := new(AiApiResponse)
	err := c.cc.Invoke(ctx, "/ai.v1.AiApiService/GetIdFromAi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiApiServiceServer is the server API for AiApiService service.
// All implementations must embed UnimplementedAiApiServiceServer
// for forward compatibility
type AiApiServiceServer interface {
	GetIdFromAi(context.Context, *AiApiRequest) (*AiApiResponse, error)
	mustEmbedUnimplementedAiApiServiceServer()
}

// UnimplementedAiApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAiApiServiceServer struct {
}

func (UnimplementedAiApiServiceServer) GetIdFromAi(context.Context, *AiApiRequest) (*AiApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdFromAi not implemented")
}
func (UnimplementedAiApiServiceServer) mustEmbedUnimplementedAiApiServiceServer() {}

// UnsafeAiApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AiApiServiceServer will
// result in compilation errors.
type UnsafeAiApiServiceServer interface {
	mustEmbedUnimplementedAiApiServiceServer()
}

func RegisterAiApiServiceServer(s grpc.ServiceRegistrar, srv AiApiServiceServer) {
	s.RegisterService(&AiApiService_ServiceDesc, srv)
}

func _AiApiService_GetIdFromAi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AiApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiApiServiceServer).GetIdFromAi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.v1.AiApiService/GetIdFromAi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiApiServiceServer).GetIdFromAi(ctx, req.(*AiApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AiApiService_ServiceDesc is the grpc.ServiceDesc for AiApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AiApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ai.v1.AiApiService",
	HandlerType: (*AiApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdFromAi",
			Handler:    _AiApiService_GetIdFromAi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai.proto",
}
