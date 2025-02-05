// compile: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/reflog.proto; protoc-go-inject-tag --input=pb/reflog.pb.go
// request: grpcurl -proto pb/reflog.proto -plaintext -d '{"req": "a b"}' localhost:29991 pb.reflogService/Reflog

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pb/reflog.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReflogService_Reflog_FullMethodName = "/pb.ReflogService/Reflog"
)

// ReflogServiceClient is the client API for ReflogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReflogServiceClient interface {
	Reflog(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type reflogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReflogServiceClient(cc grpc.ClientConnInterface) ReflogServiceClient {
	return &reflogServiceClient{cc}
}

func (c *reflogServiceClient) Reflog(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ReflogService_Reflog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflogServiceServer is the server API for ReflogService service.
// All implementations must embed UnimplementedReflogServiceServer
// for forward compatibility.
type ReflogServiceServer interface {
	Reflog(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedReflogServiceServer()
}

// UnimplementedReflogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReflogServiceServer struct{}

func (UnimplementedReflogServiceServer) Reflog(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reflog not implemented")
}
func (UnimplementedReflogServiceServer) mustEmbedUnimplementedReflogServiceServer() {}
func (UnimplementedReflogServiceServer) testEmbeddedByValue()                       {}

// UnsafeReflogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflogServiceServer will
// result in compilation errors.
type UnsafeReflogServiceServer interface {
	mustEmbedUnimplementedReflogServiceServer()
}

func RegisterReflogServiceServer(s grpc.ServiceRegistrar, srv ReflogServiceServer) {
	// If the following call pancis, it indicates UnimplementedReflogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReflogService_ServiceDesc, srv)
}

func _ReflogService_Reflog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflogServiceServer).Reflog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReflogService_Reflog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflogServiceServer).Reflog(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflogService_ServiceDesc is the grpc.ServiceDesc for ReflogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReflogService",
	HandlerType: (*ReflogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reflog",
			Handler:    _ReflogService_Reflog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/reflog.proto",
}
