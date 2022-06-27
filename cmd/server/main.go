package main

import (
	"fmt"
	"short-url/config"
	"short-url/pkg/server"
)

func main() {
	address := fmt.Sprintf("%s:%d", config.Cfg.ServerCfg.Host, config.Cfg.ServerCfg.Port)
	server.Run(address)
}
