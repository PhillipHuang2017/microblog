package svc

import (
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/config"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/sql/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
