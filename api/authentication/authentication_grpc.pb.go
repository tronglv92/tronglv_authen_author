// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: api/authentication/authentication.proto

package authentication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Authentication_Validate_FullMethodName    = "/pmc.inapp.identity.authorization.Authentication/Validate"
	Authentication_ClientToken_FullMethodName = "/pmc.inapp.identity.authorization.Authentication/ClientToken"
	Authentication_EcomToken_FullMethodName   = "/pmc.inapp.identity.authorization.Authentication/EcomToken"
	Authentication_UserToken_FullMethodName   = "/pmc.inapp.identity.authorization.Authentication/UserToken"
)

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	Validate(ctx context.Context, in *TokenAuthReq, opts ...grpc.CallOption) (*TokenAuthResp, error)
	ClientToken(ctx context.Context, in *ClientAuthReq, opts ...grpc.CallOption) (*ClientAuthResp, error)
	EcomToken(ctx context.Context, in *ClientAuthReq, opts ...grpc.CallOption) (*EcomAuthResp, error)
	UserToken(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthResp, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Validate(ctx context.Context, in *TokenAuthReq, opts ...grpc.CallOption) (*TokenAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenAuthResp)
	err := c.cc.Invoke(ctx, Authentication_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ClientToken(ctx context.Context, in *ClientAuthReq, opts ...grpc.CallOption) (*ClientAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientAuthResp)
	err := c.cc.Invoke(ctx, Authentication_ClientToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) EcomToken(ctx context.Context, in *ClientAuthReq, opts ...grpc.CallOption) (*EcomAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EcomAuthResp)
	err := c.cc.Invoke(ctx, Authentication_EcomToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) UserToken(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserAuthResp)
	err := c.cc.Invoke(ctx, Authentication_UserToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	Validate(context.Context, *TokenAuthReq) (*TokenAuthResp, error)
	ClientToken(context.Context, *ClientAuthReq) (*ClientAuthResp, error)
	EcomToken(context.Context, *ClientAuthReq) (*EcomAuthResp, error)
	UserToken(context.Context, *UserAuthReq) (*UserAuthResp, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) Validate(context.Context, *TokenAuthReq) (*TokenAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthenticationServer) ClientToken(context.Context, *ClientAuthReq) (*ClientAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientToken not implemented")
}
func (UnimplementedAuthenticationServer) EcomToken(context.Context, *ClientAuthReq) (*EcomAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EcomToken not implemented")
}
func (UnimplementedAuthenticationServer) UserToken(context.Context, *UserAuthReq) (*UserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Validate(ctx, req.(*TokenAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ClientToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ClientToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_ClientToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ClientToken(ctx, req.(*ClientAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_EcomToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).EcomToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_EcomToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).EcomToken(ctx, req.(*ClientAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_UserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).UserToken(ctx, req.(*UserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pmc.inapp.identity.authorization.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _Authentication_Validate_Handler,
		},
		{
			MethodName: "ClientToken",
			Handler:    _Authentication_ClientToken_Handler,
		},
		{
			MethodName: "EcomToken",
			Handler:    _Authentication_EcomToken_Handler,
		},
		{
			MethodName: "UserToken",
			Handler:    _Authentication_UserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/authentication/authentication.proto",
}
