package main

import (
	"github.com/uma-31/switchboard/agent/infrastructure/config"
	"github.com/uma-31/switchboard/agent/infrastructure/di/wire"
	"github.com/uma-31/switchboard/agent/infrastructure/mdns"
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

	mdnsServer, err := mdns.NewServer(conf.ComputerInfo.ID(), int(conf.Port.Value()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mdnsServer.Shutdown(); err != nil {
			panic(err)
		}
	}()

	server, err := wire.InitializeGinServer(conf.ComputerInfo, conf.Port)
	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}
