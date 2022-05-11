package logic

import (
	"context"

	"Ningxi-Compose/usermanage/rpc/internal/svc"
	"Ningxi-Compose/usermanage/rpc/types/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *client.LoginRequest) (*client.Response, error) {
	// todo: add your logic here and delete this line

	return &client.Response{}, nil
}
