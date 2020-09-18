package main

import (
	"flag"
	"net/http"

	"github.com/wjames2000/go-zero/core/conf"
	"github.com/wjames2000/go-zero/core/logx"
	"github.com/wjames2000/go-zero/core/service"
	"github.com/wjames2000/go-zero/example/tracing/remote/portal"
	"github.com/wjames2000/go-zero/rest"
	"github.com/wjames2000/go-zero/rest/httpx"
	"github.com/wjames2000/go-zero/rpcx"
)

var (
	configFile = flag.String("f", "config.json", "the config file")
	client     rpcx.Client
)

func handle(w http.ResponseWriter, r *http.Request) {
	conn := client.Conn()
	greet := portal.NewPortalClient(conn)
	resp, err := greet.Portal(r.Context(), &portal.PortalRequest{
		Name: "kevin",
	})
	if err != nil {
		httpx.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	} else {
		httpx.OkJson(w, resp.Response)
	}
}

func main() {
	flag.Parse()

	var c rpcx.RpcClientConf
	conf.MustLoad(*configFile, &c)
	client = rpcx.MustNewClient(c)
	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: 3333,
	})
	defer engine.Stop()

	engine.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: handle,
	})
	engine.Start()
}
