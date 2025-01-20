// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Posts

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

// PostsClient is the client API for Posts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsClient interface {
	AddPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	ChangePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	DeletePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	GetUserPosts(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*PostsPack, error)
	GetFeed(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostsPack, error)
}

type postsClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsClient(cc grpc.ClientConnInterface) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) AddPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Posts.Posts/AddPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) ChangePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Posts.Posts/ChangePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) DeletePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Posts.Posts/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetUserPosts(ctx context.Context, in *UserPair, opts ...grpc.CallOption) (*PostsPack, error) {
	out := new(PostsPack)
	err := c.cc.Invoke(ctx, "/Posts.Posts/GetUserPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetFeed(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostsPack, error) {
	out := new(PostsPack)
	err := c.cc.Invoke(ctx, "/Posts.Posts/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServer is the server API for Posts service.
// All implementations must embed UnimplementedPostsServer
// for forward compatibility
type PostsServer interface {
	AddPost(context.Context, *Post) (*Status, error)
	ChangePost(context.Context, *Post) (*Status, error)
	DeletePost(context.Context, *Post) (*Status, error)
	GetUserPosts(context.Context, *UserPair) (*PostsPack, error)
	GetFeed(context.Context, *Post) (*PostsPack, error)
	mustEmbedUnimplementedPostsServer()
}

// UnimplementedPostsServer must be embedded to have forward compatible implementations.
type UnimplementedPostsServer struct {
}

func (UnimplementedPostsServer) AddPost(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPost not implemented")
}
func (UnimplementedPostsServer) ChangePost(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePost not implemented")
}
func (UnimplementedPostsServer) DeletePost(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostsServer) GetUserPosts(context.Context, *UserPair) (*PostsPack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPosts not implemented")
}
func (UnimplementedPostsServer) GetFeed(context.Context, *Post) (*PostsPack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedPostsServer) mustEmbedUnimplementedPostsServer() {}

// UnsafePostsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsServer will
// result in compilation errors.
type UnsafePostsServer interface {
	mustEmbedUnimplementedPostsServer()
}

func RegisterPostsServer(s grpc.ServiceRegistrar, srv PostsServer) {
	s.RegisterService(&Posts_ServiceDesc, srv)
}

func _Posts_AddPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).AddPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts.Posts/AddPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).AddPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_ChangePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).ChangePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts.Posts/ChangePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).ChangePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts.Posts/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).DeletePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetUserPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetUserPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts.Posts/GetUserPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetUserPosts(ctx, req.(*UserPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Posts.Posts/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetFeed(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

// Posts_ServiceDesc is the grpc.ServiceDesc for Posts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Posts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Posts.Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPost",
			Handler:    _Posts_AddPost_Handler,
		},
		{
			MethodName: "ChangePost",
			Handler:    _Posts_ChangePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Posts_DeletePost_Handler,
		},
		{
			MethodName: "GetUserPosts",
			Handler:    _Posts_GetUserPosts_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _Posts_GetFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts.proto",
}
