package logic

import (
	"context"

	"Ningxi-Compose/user/rpc/internal/svc"
	"Ningxi-Compose/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoRequest) (*user.Response, error) {
	// todo: add your logic here and delete this line

	return &user.Response{}, nil
}
