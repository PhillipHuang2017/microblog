// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "gitee.com/phillip_huang/redis-demo/service/api/user/internal/handler/user"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
	)
}