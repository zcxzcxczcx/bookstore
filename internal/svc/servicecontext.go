package svc

import (
	"bookstore/internal/config"
	"bookstore/rpc/add/adder"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     // 手动代码
	Checker checker.Checker // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(rpcx.MustNewClient(c.Add)),       // 手动代码
		Checker: checker.NewChecker(rpcx.MustNewClient(c.Check)), // 手动代码
	}
}
