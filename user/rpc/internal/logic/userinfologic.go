package logic

import (
	"context"
	"fmt"
	models "qiyaowu-go-zero/user/model"

	"qiyaowu-go-zero/user/rpc/internal/svc"
	"qiyaowu-go-zero/user/rpc/types/user"

	"github.com/pquerna/ffjson/ffjson"
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
	fmt.Println(in.Userid)
	userInfo := models.NewUserModel().GetItem(in.Userid)

	userInfonJson, _ := ffjson.Marshal(userInfo)

	return &user.Response{
		Code:    0,
		Message: "ok",
		Data:    string(userInfonJson),
	}, nil

}
