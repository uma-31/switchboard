package service

import "github.com/uma-31/switchboard/manager/domain/entity"

// コンピュータの起動を提供するサービスのインターフェイス。
type IWakeComputerService interface {
	// コンピュータを起動する。
	Wake(computer *entity.ComputerEntity) error
}
