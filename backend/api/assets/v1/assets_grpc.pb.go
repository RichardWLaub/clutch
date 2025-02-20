// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: assets/v1/assets.proto

package assetsv1

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

// AssetsAPIClient is the client API for AssetsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetsAPIClient interface {
	// Fetch is a simple endpoint that is used to execute middleware (e.g. authentication) before serving an asset.
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type assetsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetsAPIClient(cc grpc.ClientConnInterface) AssetsAPIClient {
	return &assetsAPIClient{cc}
}

func (c *assetsAPIClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/clutch.assets.v1.AssetsAPI/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetsAPIServer is the server API for AssetsAPI service.
// All implementations should embed UnimplementedAssetsAPIServer
// for forward compatibility
type AssetsAPIServer interface {
	// Fetch is a simple endpoint that is used to execute middleware (e.g. authentication) before serving an asset.
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
}

// UnimplementedAssetsAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAssetsAPIServer struct {
}

func (UnimplementedAssetsAPIServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

// UnsafeAssetsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetsAPIServer will
// result in compilation errors.
type UnsafeAssetsAPIServer interface {
	mustEmbedUnimplementedAssetsAPIServer()
}

func RegisterAssetsAPIServer(s grpc.ServiceRegistrar, srv AssetsAPIServer) {
	s.RegisterService(&AssetsAPI_ServiceDesc, srv)
}

func _AssetsAPI_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetsAPIServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.assets.v1.AssetsAPI/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetsAPIServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetsAPI_ServiceDesc is the grpc.ServiceDesc for AssetsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.assets.v1.AssetsAPI",
	HandlerType: (*AssetsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _AssetsAPI_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/v1/assets.proto",
}
