package config

import (
	"github.com/smallnest/rpcx/"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Add   rpcx.RpcClientConf // 手动代码
	Check rpcx.RpcClientConf // 手动代码
}
