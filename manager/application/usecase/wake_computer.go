package usecase

import (
	"github.com/uma-31/switchboard/manager/domain/repository"
	"github.com/uma-31/switchboard/manager/domain/service"
)

// 対象となるコンピュータの情報が取得できなかったことを示すエラー。
type TargetComputerNotFoundError struct {
	cause error
}

func (e *TargetComputerNotFoundError) Error() string {
	return "対象となるコンピュータが見つかりませんでした: " + e.cause.Error()
}

// コンピュータの起動に失敗したことを示すエラー。
type WakeComputerFailedError struct {
	cause error
}

func (e *WakeComputerFailedError) Error() string {
	return "コンピュータの起動に失敗しました: " + e.cause.Error()
}

// コンピュータの起動機能を提供するユースケース。
type WakeComputerUseCase struct {
	computerRepository  repository.IComputerRepository
	wakeComputerService service.IWakeComputerService
}

// WakeComputerUseCase のインスタンスを生成する。
func NewWakeComputerUseCase(
	computerRepository repository.IComputerRepository,
	wakeComputerService service.IWakeComputerService,
) *WakeComputerUseCase {
	return &WakeComputerUseCase{computerRepository, wakeComputerService}
}

// コンピュータを起動する。
func (u *WakeComputerUseCase) Execute(computerID string) error {
	computer, err := u.computerRepository.Find(computerID)
	if err != nil {
		return &TargetComputerNotFoundError{err}
	}

	if err := u.wakeComputerService.Wake(computer); err != nil {
		return &WakeComputerFailedError{err}
	}

	return nil
}
