// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: awsume.proto

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

// ArgumentsClient is the client API for Arguments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArgumentsClient interface {
	Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArgumentsMsg, error)
	Post(ctx context.Context, in *ArgumentsMsg, opts ...grpc.CallOption) (*ArgumentsMsg, error)
}

type argumentsClient struct {
	cc grpc.ClientConnInterface
}

func NewArgumentsClient(cc grpc.ClientConnInterface) ArgumentsClient {
	return &argumentsClient{cc}
}

func (c *argumentsClient) Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Arguments/Pre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *argumentsClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArgumentsMsg, error) {
	out := new(ArgumentsMsg)
	err := c.cc.Invoke(ctx, "/proto.Arguments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *argumentsClient) Post(ctx context.Context, in *ArgumentsMsg, opts ...grpc.CallOption) (*ArgumentsMsg, error) {
	out := new(ArgumentsMsg)
	err := c.cc.Invoke(ctx, "/proto.Arguments/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArgumentsServer is the server API for Arguments service.
// All implementations must embed UnimplementedArgumentsServer
// for forward compatibility
type ArgumentsServer interface {
	Pre(context.Context, *Empty) (*Empty, error)
	Get(context.Context, *Empty) (*ArgumentsMsg, error)
	Post(context.Context, *ArgumentsMsg) (*ArgumentsMsg, error)
	mustEmbedUnimplementedArgumentsServer()
}

// UnimplementedArgumentsServer must be embedded to have forward compatible implementations.
type UnimplementedArgumentsServer struct {
}

func (UnimplementedArgumentsServer) Pre(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pre not implemented")
}
func (UnimplementedArgumentsServer) Get(context.Context, *Empty) (*ArgumentsMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArgumentsServer) Post(context.Context, *ArgumentsMsg) (*ArgumentsMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedArgumentsServer) mustEmbedUnimplementedArgumentsServer() {}

// UnsafeArgumentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArgumentsServer will
// result in compilation errors.
type UnsafeArgumentsServer interface {
	mustEmbedUnimplementedArgumentsServer()
}

func RegisterArgumentsServer(s grpc.ServiceRegistrar, srv ArgumentsServer) {
	s.RegisterService(&Arguments_ServiceDesc, srv)
}

func _Arguments_Pre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgumentsServer).Pre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Arguments/Pre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgumentsServer).Pre(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arguments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgumentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Arguments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgumentsServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arguments_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArgumentsMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgumentsServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Arguments/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgumentsServer).Post(ctx, req.(*ArgumentsMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Arguments_ServiceDesc is the grpc.ServiceDesc for Arguments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arguments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Arguments",
	HandlerType: (*ArgumentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pre",
			Handler:    _Arguments_Pre_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Arguments_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Arguments_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awsume.proto",
}

// ProfilesClient is the client API for Profiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesClient interface {
	Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProfilesMsg, error)
	Post(ctx context.Context, in *ProfilesMsg, opts ...grpc.CallOption) (*Empty, error)
}

type profilesClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesClient(cc grpc.ClientConnInterface) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Profiles/Pre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProfilesMsg, error) {
	out := new(ProfilesMsg)
	err := c.cc.Invoke(ctx, "/proto.Profiles/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) Post(ctx context.Context, in *ProfilesMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Profiles/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServer is the server API for Profiles service.
// All implementations must embed UnimplementedProfilesServer
// for forward compatibility
type ProfilesServer interface {
	Pre(context.Context, *Empty) (*Empty, error)
	Get(context.Context, *Empty) (*ProfilesMsg, error)
	Post(context.Context, *ProfilesMsg) (*Empty, error)
	mustEmbedUnimplementedProfilesServer()
}

// UnimplementedProfilesServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServer struct {
}

func (UnimplementedProfilesServer) Pre(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pre not implemented")
}
func (UnimplementedProfilesServer) Get(context.Context, *Empty) (*ProfilesMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProfilesServer) Post(context.Context, *ProfilesMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedProfilesServer) mustEmbedUnimplementedProfilesServer() {}

// UnsafeProfilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServer will
// result in compilation errors.
type UnsafeProfilesServer interface {
	mustEmbedUnimplementedProfilesServer()
}

func RegisterProfilesServer(s grpc.ServiceRegistrar, srv ProfilesServer) {
	s.RegisterService(&Profiles_ServiceDesc, srv)
}

func _Profiles_Pre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).Pre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Profiles/Pre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).Pre(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Profiles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfilesMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Profiles/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).Post(ctx, req.(*ProfilesMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Profiles_ServiceDesc is the grpc.ServiceDesc for Profiles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profiles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pre",
			Handler:    _Profiles_Pre_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Profiles_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Profiles_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awsume.proto",
}

// CredentialsClient is the client API for Credentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialsClient interface {
	Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Post(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type credentialsClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialsClient(cc grpc.ClientConnInterface) CredentialsClient {
	return &credentialsClient{cc}
}

func (c *credentialsClient) Pre(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Credentials/Pre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Credentials/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Post(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Credentials/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsServer is the server API for Credentials service.
// All implementations must embed UnimplementedCredentialsServer
// for forward compatibility
type CredentialsServer interface {
	Pre(context.Context, *Empty) (*Empty, error)
	Get(context.Context, *Empty) (*Empty, error)
	Post(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedCredentialsServer()
}

// UnimplementedCredentialsServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialsServer struct {
}

func (UnimplementedCredentialsServer) Pre(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pre not implemented")
}
func (UnimplementedCredentialsServer) Get(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCredentialsServer) Post(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedCredentialsServer) mustEmbedUnimplementedCredentialsServer() {}

// UnsafeCredentialsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialsServer will
// result in compilation errors.
type UnsafeCredentialsServer interface {
	mustEmbedUnimplementedCredentialsServer()
}

func RegisterCredentialsServer(s grpc.ServiceRegistrar, srv CredentialsServer) {
	s.RegisterService(&Credentials_ServiceDesc, srv)
}

func _Credentials_Pre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Pre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credentials/Pre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Pre(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credentials/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credentials/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Post(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Credentials_ServiceDesc is the grpc.ServiceDesc for Credentials service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Credentials_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Credentials",
	HandlerType: (*CredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pre",
			Handler:    _Credentials_Pre_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Credentials_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Credentials_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awsume.proto",
}

// ProfileNamesClient is the client API for ProfileNames service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileNamesClient interface {
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type profileNamesClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileNamesClient(cc grpc.ClientConnInterface) ProfileNamesClient {
	return &profileNamesClient{cc}
}

func (c *profileNamesClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.ProfileNames/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileNamesServer is the server API for ProfileNames service.
// All implementations must embed UnimplementedProfileNamesServer
// for forward compatibility
type ProfileNamesServer interface {
	Get(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedProfileNamesServer()
}

// UnimplementedProfileNamesServer must be embedded to have forward compatible implementations.
type UnimplementedProfileNamesServer struct {
}

func (UnimplementedProfileNamesServer) Get(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProfileNamesServer) mustEmbedUnimplementedProfileNamesServer() {}

// UnsafeProfileNamesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileNamesServer will
// result in compilation errors.
type UnsafeProfileNamesServer interface {
	mustEmbedUnimplementedProfileNamesServer()
}

func RegisterProfileNamesServer(s grpc.ServiceRegistrar, srv ProfileNamesServer) {
	s.RegisterService(&ProfileNames_ServiceDesc, srv)
}

func _ProfileNames_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileNamesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProfileNames/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileNamesServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileNames_ServiceDesc is the grpc.ServiceDesc for ProfileNames service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileNames_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProfileNames",
	HandlerType: (*ProfileNamesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProfileNames_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awsume.proto",
}
