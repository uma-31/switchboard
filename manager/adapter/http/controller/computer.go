package controller

import "github.com/uma-31/switchboard/manager/application/usecase"

// コンピュータの起動に失敗したことを示すエラー。
type FailedToWakeComputerError struct {
	cause error
}

func (e *FailedToWakeComputerError) Error() string {
	return "コンピュータの起動に失敗しました: " + e.cause.Error()
}

// 特定のコンピュータに関する操作を提供するコントローラ。
type ComputerController struct {
	wakeComputerUseCase *usecase.WakeComputerUseCase
}

// ComputerController のインスタンスを生成する。
func NewComputerController(
	wakeComputerUseCase *usecase.WakeComputerUseCase,
) *ComputerController {
	return &ComputerController{wakeComputerUseCase}
}

// 指定した ID のコンピュータを起動する。
func (c *ComputerController) WakeComputer(computerID string) error {
	if err := c.wakeComputerUseCase.Execute(computerID); err != nil {
		return &FailedToWakeComputerError{err}
	}

	return nil
}
