// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Subs

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

// SubsClient is the client API for Subs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubsClient interface {
	Subscribe(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error)
	UnSubscribe(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error)
	ShowSubscribes(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Subscribers, error)
	Blacklist(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error)
	UnBlackList(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error)
	ShowBlackList(ctx context.Context, in *User, opts ...grpc.CallOption) (*BlackList, error)
}

type subsClient struct {
	cc grpc.ClientConnInterface
}

func NewSubsClient(cc grpc.ClientConnInterface) SubsClient {
	return &subsClient{cc}
}

func (c *subsClient) Subscribe(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Subs.Subs/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subsClient) UnSubscribe(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Subs.Subs/UnSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subsClient) ShowSubscribes(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Subscribers, error) {
	out := new(Subscribers)
	err := c.cc.Invoke(ctx, "/Subs.Subs/ShowSubscribes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subsClient) Blacklist(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Subs.Subs/Blacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subsClient) UnBlackList(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Subs.Subs/UnBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subsClient) ShowBlackList(ctx context.Context, in *User, opts ...grpc.CallOption) (*BlackList, error) {
	out := new(BlackList)
	err := c.cc.Invoke(ctx, "/Subs.Subs/ShowBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubsServer is the server API for Subs service.
// All implementations must embed UnimplementedSubsServer
// for forward compatibility
type SubsServer interface {
	Subscribe(context.Context, *UserPair) (*Status, error)
	UnSubscribe(context.Context, *UserPair) (*Status, error)
	ShowSubscribes(context.Context, *UserPair) (*Subscribers, error)
	Blacklist(context.Context, *UserPair) (*Status, error)
	UnBlackList(context.Context, *UserPair) (*Status, error)
	ShowBlackList(context.Context, *User) (*BlackList, error)
	mustEmbedUnimplementedSubsServer()
}

// UnimplementedSubsServer must be embedded to have forward compatible implementations.
type UnimplementedSubsServer struct {
}

func (UnimplementedSubsServer) Subscribe(context.Context, *UserPair) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubsServer) UnSubscribe(context.Context, *UserPair) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribe not implemented")
}
func (UnimplementedSubsServer) ShowSubscribes(context.Context, *UserPair) (*Subscribers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowSubscribes not implemented")
}
func (UnimplementedSubsServer) Blacklist(context.Context, *UserPair) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blacklist not implemented")
}
func (UnimplementedSubsServer) UnBlackList(context.Context, *UserPair) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlackList not implemented")
}
func (UnimplementedSubsServer) ShowBlackList(context.Context, *User) (*BlackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowBlackList not implemented")
}
func (UnimplementedSubsServer) mustEmbedUnimplementedSubsServer() {}

// UnsafeSubsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubsServer will
// result in compilation errors.
type UnsafeSubsServer interface {
	mustEmbedUnimplementedSubsServer()
}

func RegisterSubsServer(s grpc.ServiceRegistrar, srv SubsServer) {
	s.RegisterService(&Subs_ServiceDesc, srv)
}

func _Subs_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).Subscribe(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subs_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).UnSubscribe(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subs_ShowSubscribes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).ShowSubscribes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/ShowSubscribes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).ShowSubscribes(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subs_Blacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).Blacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/Blacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).Blacklist(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subs_UnBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).UnBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/UnBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).UnBlackList(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subs_ShowBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubsServer).ShowBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subs.Subs/ShowBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubsServer).ShowBlackList(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Subs_ServiceDesc is the grpc.ServiceDesc for Subs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Subs.Subs",
	HandlerType: (*SubsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Subs_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Subs_UnSubscribe_Handler,
		},
		{
			MethodName: "ShowSubscribes",
			Handler:    _Subs_ShowSubscribes_Handler,
		},
		{
			MethodName: "Blacklist",
			Handler:    _Subs_Blacklist_Handler,
		},
		{
			MethodName: "UnBlackList",
			Handler:    _Subs_UnBlackList_Handler,
		},
		{
			MethodName: "ShowBlackList",
			Handler:    _Subs_ShowBlackList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subs.proto",
}
