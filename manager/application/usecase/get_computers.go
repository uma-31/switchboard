package usecase

import "github.com/uma-31/switchboard/manager/domain/repository"

// コンピュータ情報の一覧取得に失敗したことを示すエラー。
type FailedToGetComputersError struct {
	cause error
}

func (e *FailedToGetComputersError) Error() string {
	return "コンピュータ一覧の取得に失敗しました: " + e.cause.Error()
}

// 登録済みのコンピュータ一覧の取得機能を提供するユースケース。
type GetComputersUseCase struct {
	computerRepository repository.IComputerRepository
}

// GetComputersUseCase のインスタンスを生成する。
func NewGetComputersUseCase(computerRepository repository.IComputerRepository) *GetComputersUseCase {
	return &GetComputersUseCase{computerRepository}
}

// 登録済みのコンピュータ一覧を取得する。
func (u *GetComputersUseCase) Execute() ([]*ComputerDTO, error) {
	computers, err := u.computerRepository.FindAll()
	if err != nil {
		return nil, &FailedToGetComputersError{err}
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
