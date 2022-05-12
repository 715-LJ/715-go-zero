package svc

import (
	"qiyaowu-go-zero/user/api/internal/config"
	models "qiyaowu-go-zero/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *models.UserModel
	//UserRpc   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//UserModel: models.NewUserModel(),
	}
}
