package main

import (
	"github.com/uma-31/switchboard/manager/infrastructure/config"
	"github.com/uma-31/switchboard/manager/infrastructure/di/wire"
)

func main() {
	configFilePath, err := config.NewFilePath()
	if err != nil {
		panic(err)
	}

	conf, err := config.Load(configFilePath)
	if err != nil {
		panic(err)
	}

	server, err := wire.InitializeGinServer(conf.Port)
	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}
