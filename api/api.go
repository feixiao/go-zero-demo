package main

import (
	"flag"
	"fmt"
	"github.com/feixiao/go-zero-demo/code"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/feixiao/go-zero-demo/api/internal/config"
	"github.com/feixiao/go-zero-demo/api/internal/handler"
	"github.com/feixiao/go-zero-demo/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {

		logx.Errorf("SetErrorHandler err:%+v", err)
		switch e := err.(type) {
		case *code.Error:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
