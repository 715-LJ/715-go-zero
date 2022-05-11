package logic

import (
	"context"

	"Ningxi-Compose/usermanage/rpc/internal/svc"
	"Ningxi-Compose/usermanage/rpc/types/client"

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

func (l *UserinfoLogic) Userinfo(in *client.UserinfoRequest) (*client.Response, error) {
	// todo: add your logic here and delete this line

	return &client.Response{}, nil
}
