package service

import "github.com/uma-31/switchboard/manager/domain/entity"

// ネットワーク上のコンピュータを検索するサービス。
type IScanComputersService interface {
	// ネットワーク上のコンピュータを検索する。
	ScanComputers() ([]*entity.ComputerEntity, error)
}
