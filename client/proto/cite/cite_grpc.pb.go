// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: cite/cite.proto

package cite

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

const (
	CiteService_GetCite_FullMethodName = "/cite.CiteService/GetCite"
)

// CiteServiceClient is the client API for CiteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CiteServiceClient interface {
	GetCite(ctx context.Context, opts ...grpc.CallOption) (CiteService_GetCiteClient, error)
}

type citeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCiteServiceClient(cc grpc.ClientConnInterface) CiteServiceClient {
	return &citeServiceClient{cc}
}

func (c *citeServiceClient) GetCite(ctx context.Context, opts ...grpc.CallOption) (CiteService_GetCiteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CiteService_ServiceDesc.Streams[0], CiteService_GetCite_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &citeServiceGetCiteClient{stream}
	return x, nil
}

type CiteService_GetCiteClient interface {
	Send(*POWRequest) error
	Recv() (*CiteResponse, error)
	grpc.ClientStream
}

type citeServiceGetCiteClient struct {
	grpc.ClientStream
}

func (x *citeServiceGetCiteClient) Send(m *POWRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *citeServiceGetCiteClient) Recv() (*CiteResponse, error) {
	m := new(CiteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CiteServiceServer is the server API for CiteService service.
// All implementations must embed UnimplementedCiteServiceServer
// for forward compatibility
type CiteServiceServer interface {
	GetCite(CiteService_GetCiteServer) error
	mustEmbedUnimplementedCiteServiceServer()
}

// UnimplementedCiteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCiteServiceServer struct {
}

func (UnimplementedCiteServiceServer) GetCite(CiteService_GetCiteServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCite not implemented")
}
func (UnimplementedCiteServiceServer) mustEmbedUnimplementedCiteServiceServer() {}

// UnsafeCiteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CiteServiceServer will
// result in compilation errors.
type UnsafeCiteServiceServer interface {
	mustEmbedUnimplementedCiteServiceServer()
}

func RegisterCiteServiceServer(s grpc.ServiceRegistrar, srv CiteServiceServer) {
	s.RegisterService(&CiteService_ServiceDesc, srv)
}

func _CiteService_GetCite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CiteServiceServer).GetCite(&citeServiceGetCiteServer{stream})
}

type CiteService_GetCiteServer interface {
	Send(*CiteResponse) error
	Recv() (*POWRequest, error)
	grpc.ServerStream
}

type citeServiceGetCiteServer struct {
	grpc.ServerStream
}

func (x *citeServiceGetCiteServer) Send(m *CiteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *citeServiceGetCiteServer) Recv() (*POWRequest, error) {
	m := new(POWRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CiteService_ServiceDesc is the grpc.ServiceDesc for CiteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CiteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cite.CiteService",
	HandlerType: (*CiteServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCite",
			Handler:       _CiteService_GetCite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cite/cite.proto",
}