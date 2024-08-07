package repository

import "github.com/uma-31/switchboard/manager/domain/entity"

// ComputerEntity を取り扱うリポジトリのインターフェイス。
type IComputerRepository interface {
	// 指定された ID のコンピュータ情報を取得する。
	Find(id string) (*entity.ComputerEntity, error)

	// 全てのコンピュータ情報を取得する。
	FindAll() ([]*entity.ComputerEntity, error)

	// コンピュータ情報を保存する。
	Save(computer *entity.ComputerEntity) error
}
