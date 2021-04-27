package svc

import (
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/config"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
