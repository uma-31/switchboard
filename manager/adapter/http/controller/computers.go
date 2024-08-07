package controller

import "github.com/uma-31/switchboard/manager/application/usecase"

// コンピュータ情報の一覧取得に失敗したことを示すエラー。
type GetComputersFailedError struct {
	cause error
}

func (e *GetComputersFailedError) Error() string {
	return "コンピュータ一覧の取得に失敗しました: " + e.cause.Error()
}

// コンピュータのスキャンに失敗したことを示すエラー。
type ScanComputersFailedError struct {
	cause error
}

func (e *ScanComputersFailedError) Error() string {
	return "コンピュータのスキャンに失敗しました: " + e.cause.Error()
}

// コンピュータの情報の保存に失敗したことを示すエラー。
type SaveComputersFailedError struct {
	cause error
}

func (e *SaveComputersFailedError) Error() string {
	return "コンピュータの情報の保存に失敗しました: " + e.cause.Error()
}

// コンピュータに関する操作を提供するコントローラ。
type ComputersController struct {
	getComputersUseCase  *usecase.GetComputersUseCase
	saveComputersUseCase *usecase.SaveComputersUseCase
	scanComputersUseCase *usecase.ScanComputersUseCase
}

// ComputersController のインスタンスを生成する。
func NewComputersController(
	getComputersUseCase *usecase.GetComputersUseCase,
	saveComputersUseCase *usecase.SaveComputersUseCase,
	scanComputerUseCase *usecase.ScanComputersUseCase,
) *ComputersController {
	return &ComputersController{
		getComputersUseCase,
		saveComputersUseCase,
		scanComputerUseCase,
	}
}

// 登録済みのコンピュータ一覧を取得する。
func (c *ComputersController) GetComputers() ([]*ComputerDTO, error) {
	computers, err := c.getComputersUseCase.Execute()
	if err != nil {
		return nil, &GetComputersFailedError{err}
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

// ネットワーク上のコンピュータを検索する。
func (c *ComputersController) ScanComputers() ([]*ComputerDTO, error) {
	computers, err := c.scanComputersUseCase.Execute()
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

// ネットワーク上のコンピュータを検索し、取得した情報を保存する。
func (c *ComputersController) ScanAndSaveComputers() ([]*ComputerDTO, error) {
	computers, err := c.scanComputersUseCase.Execute()
	if err != nil {
		return nil, &ScanComputersFailedError{err}
	}

	if err := c.saveComputersUseCase.Execute(computers); err != nil {
		return nil, &SaveComputersFailedError{err}
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
