package config

import (
	"github.com/wjames2000/go-zero/rest"
	"github.com/wjames2000/go-zero/rpcx"
)

type Config struct {
	rest.RestConf
	Rpc rpcx.RpcClientConf
}
