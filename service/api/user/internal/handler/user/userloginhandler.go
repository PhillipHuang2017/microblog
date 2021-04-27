package handler

import (
	"net/http"

	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/logic/user"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), ctx)
		resp, err := l.UserLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
