package main

import (
	"flag"
	"fmt"

	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/config"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/handler"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
