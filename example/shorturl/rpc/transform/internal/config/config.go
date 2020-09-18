package config

import (
	"github.com/wjames2000/go-zero/core/stores/cache"
	"github.com/wjames2000/go-zero/rpcx"
)

type Config struct {
	rpcx.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
}
