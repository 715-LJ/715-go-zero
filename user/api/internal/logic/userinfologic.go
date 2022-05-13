package logic

import (
	"context"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/zeromicro/go-zero/core/logx"
	"qiyaowu-go-zero/user/api/internal/svc"
	"qiyaowu-go-zero/user/api/internal/types"
	"qiyaowu-go-zero/user/rpc/types/user"
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
	userInfo, err := l.svcCtx.UserRpc.Userinfo(l.ctx, &user.UserinfoRequest{
		Userid: req.Userid,
	})
	if err != nil {
		logx.Errorf("数据不存在：", err)
		return nil, err
	}

	info := make(map[string]interface{}, 0)
	ffjson.Unmarshal([]byte(userInfo.Data), &info)
	return &types.UserinfoResponse{
		Id:       int64(info["id"].(float64)),
		Username: info["username"].(string),
		Email:    info["email"].(string),
	}, nil
}
