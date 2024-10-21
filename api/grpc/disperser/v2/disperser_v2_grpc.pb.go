// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: disperser/v2/disperser_v2.proto

package v2

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
	Disperser_DisperseBlob_FullMethodName      = "/disperser.v2.Disperser/DisperseBlob"
	Disperser_GetBlobStatus_FullMethodName     = "/disperser.v2.Disperser/GetBlobStatus"
	Disperser_GetBlobCommitment_FullMethodName = "/disperser.v2.Disperser/GetBlobCommitment"
)

// DisperserClient is the client API for Disperser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisperserClient interface {
	// DisperseBlob accepts blob to disperse from clients.
	// This executes the dispersal asynchronously, i.e. it returns once the request
	// is accepted. The client could use GetBlobStatus() API to poll the the
	// processing status of the blob.
	DisperseBlob(ctx context.Context, in *DisperseBlobRequest, opts ...grpc.CallOption) (*DisperseBlobReply, error)
	// GetBlobStatus is meant to be polled for the blob status.
	GetBlobStatus(ctx context.Context, in *BlobStatusRequest, opts ...grpc.CallOption) (*BlobStatusReply, error)
	// GetBlobCommitment is a utility method that calculates commitment for a blob payload.
	GetBlobCommitment(ctx context.Context, in *BlobCommitmentRequest, opts ...grpc.CallOption) (*BlobCommitmentReply, error)
}

type disperserClient struct {
	cc grpc.ClientConnInterface
}

func NewDisperserClient(cc grpc.ClientConnInterface) DisperserClient {
	return &disperserClient{cc}
}

func (c *disperserClient) DisperseBlob(ctx context.Context, in *DisperseBlobRequest, opts ...grpc.CallOption) (*DisperseBlobReply, error) {
	out := new(DisperseBlobReply)
	err := c.cc.Invoke(ctx, Disperser_DisperseBlob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disperserClient) GetBlobStatus(ctx context.Context, in *BlobStatusRequest, opts ...grpc.CallOption) (*BlobStatusReply, error) {
	out := new(BlobStatusReply)
	err := c.cc.Invoke(ctx, Disperser_GetBlobStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disperserClient) GetBlobCommitment(ctx context.Context, in *BlobCommitmentRequest, opts ...grpc.CallOption) (*BlobCommitmentReply, error) {
	out := new(BlobCommitmentReply)
	err := c.cc.Invoke(ctx, Disperser_GetBlobCommitment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisperserServer is the server API for Disperser service.
// All implementations must embed UnimplementedDisperserServer
// for forward compatibility
type DisperserServer interface {
	// DisperseBlob accepts blob to disperse from clients.
	// This executes the dispersal asynchronously, i.e. it returns once the request
	// is accepted. The client could use GetBlobStatus() API to poll the the
	// processing status of the blob.
	DisperseBlob(context.Context, *DisperseBlobRequest) (*DisperseBlobReply, error)
	// GetBlobStatus is meant to be polled for the blob status.
	GetBlobStatus(context.Context, *BlobStatusRequest) (*BlobStatusReply, error)
	// GetBlobCommitment is a utility method that calculates commitment for a blob payload.
	GetBlobCommitment(context.Context, *BlobCommitmentRequest) (*BlobCommitmentReply, error)
	mustEmbedUnimplementedDisperserServer()
}

// UnimplementedDisperserServer must be embedded to have forward compatible implementations.
type UnimplementedDisperserServer struct {
}

func (UnimplementedDisperserServer) DisperseBlob(context.Context, *DisperseBlobRequest) (*DisperseBlobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisperseBlob not implemented")
}
func (UnimplementedDisperserServer) GetBlobStatus(context.Context, *BlobStatusRequest) (*BlobStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobStatus not implemented")
}
func (UnimplementedDisperserServer) GetBlobCommitment(context.Context, *BlobCommitmentRequest) (*BlobCommitmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobCommitment not implemented")
}
func (UnimplementedDisperserServer) mustEmbedUnimplementedDisperserServer() {}

// UnsafeDisperserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisperserServer will
// result in compilation errors.
type UnsafeDisperserServer interface {
	mustEmbedUnimplementedDisperserServer()
}

func RegisterDisperserServer(s grpc.ServiceRegistrar, srv DisperserServer) {
	s.RegisterService(&Disperser_ServiceDesc, srv)
}

func _Disperser_DisperseBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisperseBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).DisperseBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_DisperseBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).DisperseBlob(ctx, req.(*DisperseBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disperser_GetBlobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).GetBlobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_GetBlobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).GetBlobStatus(ctx, req.(*BlobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disperser_GetBlobCommitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlobCommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).GetBlobCommitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_GetBlobCommitment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).GetBlobCommitment(ctx, req.(*BlobCommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Disperser_ServiceDesc is the grpc.ServiceDesc for Disperser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Disperser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "disperser.v2.Disperser",
	HandlerType: (*DisperserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisperseBlob",
			Handler:    _Disperser_DisperseBlob_Handler,
		},
		{
			MethodName: "GetBlobStatus",
			Handler:    _Disperser_GetBlobStatus_Handler,
		},
		{
			MethodName: "GetBlobCommitment",
			Handler:    _Disperser_GetBlobCommitment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "disperser/v2/disperser_v2.proto",
}