// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"Ningxi-Compose/user/rpc/internal/logic"
	"Ningxi-Compose/user/rpc/internal/svc"
	"Ningxi-Compose/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.Response, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.Response, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Userinfo(ctx context.Context, in *user.UserinfoRequest) (*user.Response, error) {
	l := logic.NewUserinfoLogic(ctx, s.svcCtx)
	return l.Userinfo(in)
}
