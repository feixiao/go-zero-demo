package main

import (
	"flag"
	"fmt"

	"github.com/feixiao/go-zero-demo/rpc/internal/config"
	"github.com/feixiao/go-zero-demo/rpc/internal/server"
	"github.com/feixiao/go-zero-demo/rpc/internal/svc"
	"github.com/feixiao/go-zero-demo/rpc/rpc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRpcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterRpcServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
