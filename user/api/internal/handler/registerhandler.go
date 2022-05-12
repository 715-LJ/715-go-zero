package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qiyaowu-go-zero/user/api/internal/logic"
	"qiyaowu-go-zero/user/api/internal/svc"
	"qiyaowu-go-zero/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
