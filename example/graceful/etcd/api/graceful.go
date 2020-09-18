package main

import (
	"flag"

	"github.com/wjames2000/go-zero/core/conf"
	"github.com/wjames2000/go-zero/example/graceful/etcd/api/config"
	"github.com/wjames2000/go-zero/example/graceful/etcd/api/handler"
	"github.com/wjames2000/go-zero/example/graceful/etcd/api/svc"
	"github.com/wjames2000/go-zero/rest"
	"github.com/wjames2000/go-zero/rpcx"
)

var configFile = flag.String("f", "etc/graceful-api.json", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	client := rpcx.MustNewClient(c.Rpc)
	ctx := &svc.ServiceContext{
		Client: client,
	}

	engine := rest.MustNewServer(c.RestConf)
	defer engine.Stop()

	handler.RegisterHandlers(engine, ctx)
	engine.Start()
}
