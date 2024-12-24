// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: eotsmanager.proto

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

const (
	EOTSManager_Ping_FullMethodName                     = "/proto.EOTSManager/Ping"
	EOTSManager_CreateKey_FullMethodName                = "/proto.EOTSManager/CreateKey"
	EOTSManager_CreateRandomnessPairList_FullMethodName = "/proto.EOTSManager/CreateRandomnessPairList"
	EOTSManager_KeyRecord_FullMethodName                = "/proto.EOTSManager/KeyRecord"
	EOTSManager_SignEOTS_FullMethodName                 = "/proto.EOTSManager/SignEOTS"
	EOTSManager_UnsafeSignEOTS_FullMethodName           = "/proto.EOTSManager/UnsafeSignEOTS"
	EOTSManager_SignSchnorrSig_FullMethodName           = "/proto.EOTSManager/SignSchnorrSig"
)

// EOTSManagerClient is the client API for EOTSManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EOTSManagerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// CreateKey generates and saves an EOTS key
	CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*CreateKeyResponse, error)
	// CreateRandomnessPairList returns a list of Schnorr randomness pairs
	CreateRandomnessPairList(ctx context.Context, in *CreateRandomnessPairListRequest, opts ...grpc.CallOption) (*CreateRandomnessPairListResponse, error)
	// KeyRecord returns the key record
	KeyRecord(ctx context.Context, in *KeyRecordRequest, opts ...grpc.CallOption) (*KeyRecordResponse, error)
	// SignEOTS signs an EOTS with the EOTS private key and the relevant randomness
	SignEOTS(ctx context.Context, in *SignEOTSRequest, opts ...grpc.CallOption) (*SignEOTSResponse, error)
	// UnsafeSignEOTS used only for testing purpose. Use SignEOTS for real operations
	UnsafeSignEOTS(ctx context.Context, in *SignEOTSRequest, opts ...grpc.CallOption) (*SignEOTSResponse, error)
	// SignSchnorrSig signs a Schnorr sig with the EOTS private key
	SignSchnorrSig(ctx context.Context, in *SignSchnorrSigRequest, opts ...grpc.CallOption) (*SignSchnorrSigResponse, error)
}

type eOTSManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewEOTSManagerClient(cc grpc.ClientConnInterface) EOTSManagerClient {
	return &eOTSManagerClient{cc}
}

func (c *eOTSManagerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, EOTSManager_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*CreateKeyResponse, error) {
	out := new(CreateKeyResponse)
	err := c.cc.Invoke(ctx, EOTSManager_CreateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) CreateRandomnessPairList(ctx context.Context, in *CreateRandomnessPairListRequest, opts ...grpc.CallOption) (*CreateRandomnessPairListResponse, error) {
	out := new(CreateRandomnessPairListResponse)
	err := c.cc.Invoke(ctx, EOTSManager_CreateRandomnessPairList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) KeyRecord(ctx context.Context, in *KeyRecordRequest, opts ...grpc.CallOption) (*KeyRecordResponse, error) {
	out := new(KeyRecordResponse)
	err := c.cc.Invoke(ctx, EOTSManager_KeyRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) SignEOTS(ctx context.Context, in *SignEOTSRequest, opts ...grpc.CallOption) (*SignEOTSResponse, error) {
	out := new(SignEOTSResponse)
	err := c.cc.Invoke(ctx, EOTSManager_SignEOTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) UnsafeSignEOTS(ctx context.Context, in *SignEOTSRequest, opts ...grpc.CallOption) (*SignEOTSResponse, error) {
	out := new(SignEOTSResponse)
	err := c.cc.Invoke(ctx, EOTSManager_UnsafeSignEOTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eOTSManagerClient) SignSchnorrSig(ctx context.Context, in *SignSchnorrSigRequest, opts ...grpc.CallOption) (*SignSchnorrSigResponse, error) {
	out := new(SignSchnorrSigResponse)
	err := c.cc.Invoke(ctx, EOTSManager_SignSchnorrSig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EOTSManagerServer is the server API for EOTSManager service.
// All implementations must embed UnimplementedEOTSManagerServer
// for forward compatibility
type EOTSManagerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// CreateKey generates and saves an EOTS key
	CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error)
	// CreateRandomnessPairList returns a list of Schnorr randomness pairs
	CreateRandomnessPairList(context.Context, *CreateRandomnessPairListRequest) (*CreateRandomnessPairListResponse, error)
	// KeyRecord returns the key record
	KeyRecord(context.Context, *KeyRecordRequest) (*KeyRecordResponse, error)
	// SignEOTS signs an EOTS with the EOTS private key and the relevant randomness
	SignEOTS(context.Context, *SignEOTSRequest) (*SignEOTSResponse, error)
	// UnsafeSignEOTS used only for testing purpose. Use SignEOTS for real operations
	UnsafeSignEOTS(context.Context, *SignEOTSRequest) (*SignEOTSResponse, error)
	// SignSchnorrSig signs a Schnorr sig with the EOTS private key
	SignSchnorrSig(context.Context, *SignSchnorrSigRequest) (*SignSchnorrSigResponse, error)
	mustEmbedUnimplementedEOTSManagerServer()
}

// UnimplementedEOTSManagerServer must be embedded to have forward compatible implementations.
type UnimplementedEOTSManagerServer struct {
}

func (UnimplementedEOTSManagerServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedEOTSManagerServer) CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKey not implemented")
}
func (UnimplementedEOTSManagerServer) CreateRandomnessPairList(context.Context, *CreateRandomnessPairListRequest) (*CreateRandomnessPairListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRandomnessPairList not implemented")
}
func (UnimplementedEOTSManagerServer) KeyRecord(context.Context, *KeyRecordRequest) (*KeyRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyRecord not implemented")
}
func (UnimplementedEOTSManagerServer) SignEOTS(context.Context, *SignEOTSRequest) (*SignEOTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignEOTS not implemented")
}
func (UnimplementedEOTSManagerServer) UnsafeSignEOTS(context.Context, *SignEOTSRequest) (*SignEOTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsafeSignEOTS not implemented")
}
func (UnimplementedEOTSManagerServer) SignSchnorrSig(context.Context, *SignSchnorrSigRequest) (*SignSchnorrSigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignSchnorrSig not implemented")
}
func (UnimplementedEOTSManagerServer) mustEmbedUnimplementedEOTSManagerServer() {}

// UnsafeEOTSManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EOTSManagerServer will
// result in compilation errors.
type UnsafeEOTSManagerServer interface {
	mustEmbedUnimplementedEOTSManagerServer()
}

func RegisterEOTSManagerServer(s grpc.ServiceRegistrar, srv EOTSManagerServer) {
	s.RegisterService(&EOTSManager_ServiceDesc, srv)
}

func _EOTSManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_CreateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).CreateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_CreateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).CreateKey(ctx, req.(*CreateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_CreateRandomnessPairList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRandomnessPairListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).CreateRandomnessPairList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_CreateRandomnessPairList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).CreateRandomnessPairList(ctx, req.(*CreateRandomnessPairListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_KeyRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).KeyRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_KeyRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).KeyRecord(ctx, req.(*KeyRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_SignEOTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignEOTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).SignEOTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_SignEOTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).SignEOTS(ctx, req.(*SignEOTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_UnsafeSignEOTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignEOTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).UnsafeSignEOTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_UnsafeSignEOTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).UnsafeSignEOTS(ctx, req.(*SignEOTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EOTSManager_SignSchnorrSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignSchnorrSigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EOTSManagerServer).SignSchnorrSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EOTSManager_SignSchnorrSig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EOTSManagerServer).SignSchnorrSig(ctx, req.(*SignSchnorrSigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EOTSManager_ServiceDesc is the grpc.ServiceDesc for EOTSManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EOTSManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EOTSManager",
	HandlerType: (*EOTSManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EOTSManager_Ping_Handler,
		},
		{
			MethodName: "CreateKey",
			Handler:    _EOTSManager_CreateKey_Handler,
		},
		{
			MethodName: "CreateRandomnessPairList",
			Handler:    _EOTSManager_CreateRandomnessPairList_Handler,
		},
		{
			MethodName: "KeyRecord",
			Handler:    _EOTSManager_KeyRecord_Handler,
		},
		{
			MethodName: "SignEOTS",
			Handler:    _EOTSManager_SignEOTS_Handler,
		},
		{
			MethodName: "UnsafeSignEOTS",
			Handler:    _EOTSManager_UnsafeSignEOTS_Handler,
		},
		{
			MethodName: "SignSchnorrSig",
			Handler:    _EOTSManager_SignSchnorrSig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eotsmanager.proto",
}