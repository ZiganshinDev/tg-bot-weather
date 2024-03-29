// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: proto/cityservice.proto

package proto

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

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityServiceClient interface {
	GetCityCoordinates(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*CoordinatesReply, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) GetCityCoordinates(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*CoordinatesReply, error) {
	out := new(CoordinatesReply)
	err := c.cc.Invoke(ctx, "/cityservice.CityService/GetCityCoordinates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
// All implementations must embed UnimplementedCityServiceServer
// for forward compatibility
type CityServiceServer interface {
	GetCityCoordinates(context.Context, *CityRequest) (*CoordinatesReply, error)
	mustEmbedUnimplementedCityServiceServer()
}

// UnimplementedCityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (UnimplementedCityServiceServer) GetCityCoordinates(context.Context, *CityRequest) (*CoordinatesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCityCoordinates not implemented")
}
func (UnimplementedCityServiceServer) mustEmbedUnimplementedCityServiceServer() {}

// UnsafeCityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityServiceServer will
// result in compilation errors.
type UnsafeCityServiceServer interface {
	mustEmbedUnimplementedCityServiceServer()
}

func RegisterCityServiceServer(s grpc.ServiceRegistrar, srv CityServiceServer) {
	s.RegisterService(&CityService_ServiceDesc, srv)
}

func _CityService_GetCityCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCityCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cityservice.CityService/GetCityCoordinates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCityCoordinates(ctx, req.(*CityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CityService_ServiceDesc is the grpc.ServiceDesc for CityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cityservice.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCityCoordinates",
			Handler:    _CityService_GetCityCoordinates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cityservice.proto",
}
