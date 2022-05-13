// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"qiyaowu-go-zero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest    = user.LoginRequest
	RegisterRequest = user.RegisterRequest
	Response        = user.Response
	UserinfoRequest = user.UserinfoRequest

	User interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
		Userinfo(ctx context.Context, in *UserinfoRequest, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Userinfo(ctx context.Context, in *UserinfoRequest, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Userinfo(ctx, in, opts...)
}
