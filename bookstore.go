package main

import (
	"flag"
	"fmt"

	"bookstore/internal/config"
	"bookstore/internal/handler"
	"bookstore/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/bookstore-api.yaml", "the config file")

func main() {
	flag.Parse()
	fmt.Printf("1111111111111111\n")
	var c config.Config
	fmt.Printf("22222222222\n")
	conf.MustLoad(*configFile, &c)
	fmt.Printf("3333333333333333\n")
	ctx := svc.NewServiceContext(c)
	fmt.Printf("44444444444444\n")
	server := rest.MustNewServer(c.RestConf)
	fmt.Printf("55555555555555\n")
	defer server.Stop()
	fmt.Printf("666666666666\n")
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("77777777777777777777\n")
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	fmt.Printf("55555555555555555555555555555\n")
}
