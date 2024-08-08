package usecase

import (
	"github.com/uma-31/switchboard/manager/domain/entity"
	"github.com/uma-31/switchboard/manager/domain/repository"
)

// コンピュータの情報の保存に失敗したことを示すエラー。
type SaveComputersFailedError struct {
	cause error
}

func (e *SaveComputersFailedError) Error() string {
	return "コンピュータの情報の保存に失敗しました: " + e.cause.Error()
}

// コンピュータの情報を保存する機能を提供するユースケース。
type SaveComputersUseCase struct {
	computerRepository repository.IComputerRepository
}

// SaveComputersUseCase のインスタンスを生成する。
func NewSaveComputersUseCase(
	computerRepository repository.IComputerRepository,
) *SaveComputersUseCase {
	return &SaveComputersUseCase{computerRepository}
}

// コンピュータの情報を保存する。
func (u *SaveComputersUseCase) Execute(computers []*ComputerDTO) error {
	entities := make([]*entity.ComputerEntity, 0, len(computers))

	for _, dto := range computers {
		entities = append(
			entities,
			entity.NewComputerEntity(dto.ID, dto.Name, dto.MacAddress),
		)
	}

	for _, entity := range entities {
		// NOTE: 一斉保存機能欲しい
		if err := u.computerRepository.Save(entity); err != nil {
			return &SaveComputersFailedError{err}
		}
	}

	return nil
}
