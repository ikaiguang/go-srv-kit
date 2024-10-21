// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: testdata/ping-service/api/testdata-service/v1/services/testdata.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/go-srv-kit/testdata/ping-service/api/testdata-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvTestdata_Websocket_FullMethodName           = "/kit.api.testdata.servicev1.SrvTestdata/Websocket"
	SrvTestdata_Get_FullMethodName                 = "/kit.api.testdata.servicev1.SrvTestdata/Get"
	SrvTestdata_Put_FullMethodName                 = "/kit.api.testdata.servicev1.SrvTestdata/Put"
	SrvTestdata_Post_FullMethodName                = "/kit.api.testdata.servicev1.SrvTestdata/Post"
	SrvTestdata_Delete_FullMethodName              = "/kit.api.testdata.servicev1.SrvTestdata/Delete"
	SrvTestdata_Patch_FullMethodName               = "/kit.api.testdata.servicev1.SrvTestdata/Patch"
	SrvTestdata_ServerToClient_FullMethodName      = "/kit.api.testdata.servicev1.SrvTestdata/ServerToClient"
	SrvTestdata_ClientToServer_FullMethodName      = "/kit.api.testdata.servicev1.SrvTestdata/ClientToServer"
	SrvTestdata_BidirectionalStream_FullMethodName = "/kit.api.testdata.servicev1.SrvTestdata/BidirectionalStream"
)

// SrvTestdataClient is the client API for SrvTestdata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvTestdataClient interface {
	// Websocket websocket
	Websocket(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// Get Get
	Get(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// Put Put
	Put(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// Post Post
	Post(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// Delete Delete
	Delete(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// Patch Patch
	Patch(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error)
	// ServerToClient A server-to-client streaming RPC.
	ServerToClient(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (SrvTestdata_ServerToClientClient, error)
	// ClientToServer A client-to-server streaming RPC.
	ClientToServer(ctx context.Context, opts ...grpc.CallOption) (SrvTestdata_ClientToServerClient, error)
	// BidirectionalStream A Bidirectional streaming RPC.
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (SrvTestdata_BidirectionalStreamClient, error)
}

type srvTestdataClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvTestdataClient(cc grpc.ClientConnInterface) SrvTestdataClient {
	return &srvTestdataClient{cc}
}

func (c *srvTestdataClient) Websocket(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Websocket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) Get(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) Put(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) Post(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Post_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) Delete(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) Patch(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (*resources.TestResp, error) {
	out := new(resources.TestResp)
	err := c.cc.Invoke(ctx, SrvTestdata_Patch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvTestdataClient) ServerToClient(ctx context.Context, in *resources.TestReq, opts ...grpc.CallOption) (SrvTestdata_ServerToClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &SrvTestdata_ServiceDesc.Streams[0], SrvTestdata_ServerToClient_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &srvTestdataServerToClientClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SrvTestdata_ServerToClientClient interface {
	Recv() (*resources.TestResp, error)
	grpc.ClientStream
}

type srvTestdataServerToClientClient struct {
	grpc.ClientStream
}

func (x *srvTestdataServerToClientClient) Recv() (*resources.TestResp, error) {
	m := new(resources.TestResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *srvTestdataClient) ClientToServer(ctx context.Context, opts ...grpc.CallOption) (SrvTestdata_ClientToServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &SrvTestdata_ServiceDesc.Streams[1], SrvTestdata_ClientToServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &srvTestdataClientToServerClient{stream}
	return x, nil
}

type SrvTestdata_ClientToServerClient interface {
	Send(*resources.TestReq) error
	CloseAndRecv() (*resources.TestResp, error)
	grpc.ClientStream
}

type srvTestdataClientToServerClient struct {
	grpc.ClientStream
}

func (x *srvTestdataClientToServerClient) Send(m *resources.TestReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *srvTestdataClientToServerClient) CloseAndRecv() (*resources.TestResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(resources.TestResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *srvTestdataClient) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (SrvTestdata_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SrvTestdata_ServiceDesc.Streams[2], SrvTestdata_BidirectionalStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &srvTestdataBidirectionalStreamClient{stream}
	return x, nil
}

type SrvTestdata_BidirectionalStreamClient interface {
	Send(*resources.TestReq) error
	Recv() (*resources.TestResp, error)
	grpc.ClientStream
}

type srvTestdataBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *srvTestdataBidirectionalStreamClient) Send(m *resources.TestReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *srvTestdataBidirectionalStreamClient) Recv() (*resources.TestResp, error) {
	m := new(resources.TestResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SrvTestdataServer is the server API for SrvTestdata service.
// All implementations must embed UnimplementedSrvTestdataServer
// for forward compatibility
type SrvTestdataServer interface {
	// Websocket websocket
	Websocket(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// Get Get
	Get(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// Put Put
	Put(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// Post Post
	Post(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// Delete Delete
	Delete(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// Patch Patch
	Patch(context.Context, *resources.TestReq) (*resources.TestResp, error)
	// ServerToClient A server-to-client streaming RPC.
	ServerToClient(*resources.TestReq, SrvTestdata_ServerToClientServer) error
	// ClientToServer A client-to-server streaming RPC.
	ClientToServer(SrvTestdata_ClientToServerServer) error
	// BidirectionalStream A Bidirectional streaming RPC.
	BidirectionalStream(SrvTestdata_BidirectionalStreamServer) error
	mustEmbedUnimplementedSrvTestdataServer()
}

// UnimplementedSrvTestdataServer must be embedded to have forward compatible implementations.
type UnimplementedSrvTestdataServer struct {
}

func (UnimplementedSrvTestdataServer) Websocket(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Websocket not implemented")
}
func (UnimplementedSrvTestdataServer) Get(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSrvTestdataServer) Put(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedSrvTestdataServer) Post(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedSrvTestdataServer) Delete(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSrvTestdataServer) Patch(context.Context, *resources.TestReq) (*resources.TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedSrvTestdataServer) ServerToClient(*resources.TestReq, SrvTestdata_ServerToClientServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerToClient not implemented")
}
func (UnimplementedSrvTestdataServer) ClientToServer(SrvTestdata_ClientToServerServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientToServer not implemented")
}
func (UnimplementedSrvTestdataServer) BidirectionalStream(SrvTestdata_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}
func (UnimplementedSrvTestdataServer) mustEmbedUnimplementedSrvTestdataServer() {}

// UnsafeSrvTestdataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvTestdataServer will
// result in compilation errors.
type UnsafeSrvTestdataServer interface {
	mustEmbedUnimplementedSrvTestdataServer()
}

func RegisterSrvTestdataServer(s grpc.ServiceRegistrar, srv SrvTestdataServer) {
	s.RegisterService(&SrvTestdata_ServiceDesc, srv)
}

func _SrvTestdata_Websocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Websocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Websocket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Websocket(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Get(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Put(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Post(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Delete(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvTestdataServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvTestdata_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvTestdataServer).Patch(ctx, req.(*resources.TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvTestdata_ServerToClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(resources.TestReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SrvTestdataServer).ServerToClient(m, &srvTestdataServerToClientServer{stream})
}

type SrvTestdata_ServerToClientServer interface {
	Send(*resources.TestResp) error
	grpc.ServerStream
}

type srvTestdataServerToClientServer struct {
	grpc.ServerStream
}

func (x *srvTestdataServerToClientServer) Send(m *resources.TestResp) error {
	return x.ServerStream.SendMsg(m)
}

func _SrvTestdata_ClientToServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SrvTestdataServer).ClientToServer(&srvTestdataClientToServerServer{stream})
}

type SrvTestdata_ClientToServerServer interface {
	SendAndClose(*resources.TestResp) error
	Recv() (*resources.TestReq, error)
	grpc.ServerStream
}

type srvTestdataClientToServerServer struct {
	grpc.ServerStream
}

func (x *srvTestdataClientToServerServer) SendAndClose(m *resources.TestResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *srvTestdataClientToServerServer) Recv() (*resources.TestReq, error) {
	m := new(resources.TestReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SrvTestdata_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SrvTestdataServer).BidirectionalStream(&srvTestdataBidirectionalStreamServer{stream})
}

type SrvTestdata_BidirectionalStreamServer interface {
	Send(*resources.TestResp) error
	Recv() (*resources.TestReq, error)
	grpc.ServerStream
}

type srvTestdataBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *srvTestdataBidirectionalStreamServer) Send(m *resources.TestResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *srvTestdataBidirectionalStreamServer) Recv() (*resources.TestReq, error) {
	m := new(resources.TestReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SrvTestdata_ServiceDesc is the grpc.ServiceDesc for SrvTestdata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvTestdata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kit.api.testdata.servicev1.SrvTestdata",
	HandlerType: (*SrvTestdataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Websocket",
			Handler:    _SrvTestdata_Websocket_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SrvTestdata_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _SrvTestdata_Put_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _SrvTestdata_Post_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SrvTestdata_Delete_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _SrvTestdata_Patch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerToClient",
			Handler:       _SrvTestdata_ServerToClient_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientToServer",
			Handler:       _SrvTestdata_ClientToServer_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _SrvTestdata_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testdata/ping-service/api/testdata-service/v1/services/testdata.service.v1.proto",
}
