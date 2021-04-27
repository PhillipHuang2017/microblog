package config

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/tal-tech/go-zero/core/stores/cache"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct{
		DataSource string
	}
	CacheRedis cache.CacheConf   // 缓存结点配置列表
	UserPasswordSalt string // 用户密码加盐
}
