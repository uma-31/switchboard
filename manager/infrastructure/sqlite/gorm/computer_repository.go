package gorm

import (
	"errors"

	"github.com/uma-31/switchboard/manager/domain/entity"
	"gorm.io/gorm"
)

// コンピュータ情報の取得に失敗したことを知らせるエラー。
type FailedToFindComputerError struct {
	cause error
}

func (e *FailedToFindComputerError) Error() string {
	if e.cause != nil {
		return "コンピュータ情報の取得に失敗しました: " + e.cause.Error()
	}

	return "コンピュータ情報の取得に失敗しました。"
}

// コンピュータ情報の保存に失敗したことを知らせるエラー。
type FailedToSaveComputerError struct {
	cause error
}

func (e *FailedToSaveComputerError) Error() string {
	return "コンピュータ情報の保存に失敗しました: " + e.cause.Error()
}

// ComputerEntity を取り扱うリポジトリ。
type ComputerRepository struct {
	db *gorm.DB
}

// ComputerRepository のインスタンスを生成する。
func NewComputerRepository(db *gorm.DB) *ComputerRepository {
	return &ComputerRepository{db}
}

func (r *ComputerRepository) Find(computerID string) (*entity.ComputerEntity, error) {
	var computer Computer
	if err := r.db.First(&computer, "id = ?", computerID).Error; err != nil {
		return nil, &FailedToFindComputerError{err}
	}

	return entity.NewComputerEntity(
		computer.ID,
		computer.Name,
		computer.MacAddress,
	), nil
}

// 全てのコンピュータ情報を取得する。
func (r *ComputerRepository) FindAll() ([]*entity.ComputerEntity, error) {
	var computers []Computer
	if err := r.db.Find(&computers).Error; err != nil {
		return nil, &FailedToFindComputerError{err}
	}

	entities := make([]*entity.ComputerEntity, 0, len(computers))

	for _, computer := range computers {
		entities = append(entities, entity.NewComputerEntity(
			computer.ID,
			computer.Name,
			computer.MacAddress,
		))
	}

	return entities, nil
}

// コンピュータの情報を保存する。
func (r *ComputerRepository) Save(computerEntity *entity.ComputerEntity) error {
	var computer Computer

	// NOTE: 後で FOR UPDATE でロックするようにする。
	if err := r.db.First(&computer, "id = ?", computerEntity.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.db.Create(
				//exhaustruct:ignore
				&Computer{
					ID:         computerEntity.ID,
					Name:       computerEntity.Name,
					MacAddress: computerEntity.MacAddress,
				})

			return nil
		}

		return &FailedToSaveComputerError{err}
	}

	updated := false

	if computer.Name != computerEntity.Name {
		computer.Name = computerEntity.Name
		updated = true
	}

	if computer.MacAddress != computerEntity.MacAddress {
		computer.MacAddress = computerEntity.MacAddress
		updated = true
	}

	if updated {
		if err := r.db.Save(&computer).Error; err != nil {
			return &FailedToSaveComputerError{err}
		}
	}

	return nil
}
