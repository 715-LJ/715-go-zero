package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"qiyaowu-go-zero/user"
	"qiyaowu-go-zero/user/api/internal/svc"
	"qiyaowu-go-zero/user/api/internal/types"
	models "qiyaowu-go-zero/user/model"
	"strconv"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserinfoRequest) (resp *types.UserinfoResponse, err error) {
	userInt, err := strconv.ParseInt(req.Userid, 10, 64)
	if err != nil {
		return nil, err
	}

	userModel := models.NewUserModel()
	userInfo := userModel.GetItem(userInt)
	switch err {
	case nil:
		return &types.UserinfoResponse{
			Id:       userInfo.Id,
			UserName: userInfo.UserName,
			Email:    userInfo.Email,
		}, nil
	case user.ErrNotFound:
		return nil, err
	default:
		return nil, err
	}
	return
}
