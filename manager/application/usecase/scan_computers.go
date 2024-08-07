package usecase

import (
	"github.com/uma-31/switchboard/manager/domain/service"
)

// ネットワーク上のコンピュータ検索に失敗したことを示すエラー。
type ScanComputersFailedError struct {
	cause error
}

func (e *ScanComputersFailedError) Error() string {
	return "コンピュータのスキャンに失敗しました: " + e.cause.Error()
}

// ネットワーク上のコンピュータを検索する機能を提供するユースケース。
type ScanComputersUseCase struct {
	scanComputersService service.IScanComputersService
}

// ScanComputersUseCase のインスタンスを生成する。
func NewScanComputersUseCase(
	scanComputersService service.IScanComputersService,
) *ScanComputersUseCase {
	return &ScanComputersUseCase{scanComputersService}
}

// ネットワーク上のコンピュータを検索する。
func (u *ScanComputersUseCase) Execute() ([]*ComputerDTO, error) {
	computers, err := u.scanComputersService.ScanComputers()
	if err != nil {
		return nil, &ScanComputersFailedError{err}
	}

	dtoList := make([]*ComputerDTO, 0, len(computers))

	for _, computer := range computers {
		dtoList = append(dtoList, &ComputerDTO{
			ID:         computer.ID,
			Name:       computer.Name,
			MacAddress: computer.MacAddress,
		})
	}

	return dtoList, nil
}
