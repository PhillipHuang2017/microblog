package main

import (
	"flag"
	"fmt"

	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/config"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/server"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
