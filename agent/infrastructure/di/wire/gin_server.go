//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/uma-31/switchboard/agent/adapter/http/controller"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
	"github.com/uma-31/switchboard/agent/infrastructure/http/gin"
)

func InitializeGinServer(computerInfo *valueobject.ComputerInfo, port *gin.ServerPort) (*gin.Server, error) {
	wire.Build(
		controller.NewComputerInfoController,
		gin.NewRouter,
		gin.NewServer,
	)

	return &gin.Server{}, nil
}
