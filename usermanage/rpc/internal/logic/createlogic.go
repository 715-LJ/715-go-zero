package logic

import (
	"context"

	"Ningxi-Compose/usermanage/rpc/internal/svc"
	"Ningxi-Compose/usermanage/rpc/types/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *role.RoleRequest) (*role.Response, error) {
	// todo: add your logic here and delete this line

	return &role.Response{}, nil
}
