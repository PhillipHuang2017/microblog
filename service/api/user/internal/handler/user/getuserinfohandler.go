package handler

import (
	"net/http"

	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/logic/user"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserInfoLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
