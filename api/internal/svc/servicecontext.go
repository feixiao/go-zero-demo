package svc

import (
	"context"
	"github.com/feixiao/go-zero-demo/api/internal/config"
	"github.com/feixiao/go-zero-demo/rpc/rpcclient"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"time"
)

type ServiceContext struct {
	Config config.Config

	UserRpc rpcclient.Rpc
}

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	logx.Infof("Call %s method cost: %v\n", method, time.Now().Sub(stime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	ur := rpcclient.NewRpc(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))

	return &ServiceContext{
		UserRpc: ur,
		Config:  c,
	}
}
