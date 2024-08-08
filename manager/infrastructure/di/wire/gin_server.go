//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/uma-31/switchboard/manager/adapter/http/controller"
	"github.com/uma-31/switchboard/manager/application/usecase"
	"github.com/uma-31/switchboard/manager/domain/repository"
	"github.com/uma-31/switchboard/manager/domain/service"
	"github.com/uma-31/switchboard/manager/infrastructure/http/gin"
	"github.com/uma-31/switchboard/manager/infrastructure/mdns"
	"github.com/uma-31/switchboard/manager/infrastructure/sqlite/gorm"
	"github.com/uma-31/switchboard/manager/infrastructure/wol"
)

func InitializeGinServer(port *gin.ServerPort) (*gin.Server, error) {
	wire.Build(
		gorm.NewSqliteFilePath,
		gorm.NewDB,
		gorm.NewComputerRepository,
		wire.Bind(new(repository.IComputerRepository), new(*gorm.ComputerRepository)),
		wol.NewWakeComputerService,
		wire.Bind(new(service.IWakeComputerService), new(*wol.WakeComputerService)),
		mdns.NewScanComputerService,
		wire.Bind(new(service.IScanComputersService), new(*mdns.ScanComputerService)),
		usecase.NewWakeComputerUseCase,
		usecase.NewGetComputersUseCase,
		usecase.NewSaveComputersUseCase,
		usecase.NewScanComputersUseCase,
		controller.NewComputerController,
		controller.NewComputersController,
		gin.NewRouter,
		gin.NewServer,
	)

	return &gin.Server{}, nil
}
