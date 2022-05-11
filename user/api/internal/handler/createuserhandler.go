package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qiyaowu-go-zero/user/api/internal/logic"
	"qiyaowu-go-zero/user/api/internal/svc"
	"qiyaowu-go-zero/user/api/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
