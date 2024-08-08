package gorm

import (
	// 最新バージョンになるまで windows 386 がサポートされないので、
	// なんとかしたい。
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// データベースファイルを開けなかったことを示す例外。
type OpenDBFailedError struct {
	cause error
}

func (e *OpenDBFailedError) Error() string {
	return "データベースへの接続に失敗しました: " + e.cause.Error()
}

// 自動マイグレーションに失敗したことを示す例外。
type FailedToAutoMigrateError struct {
	cause error
}

func (e *FailedToAutoMigrateError) Error() string {
	return "自動マイグレーションに失敗しました: " + e.cause.Error()
}

// DB のインスタンスを生成する。
func NewDB(sqliteFilePath SqliteFilePath) (*gorm.DB, error) {
	//exhaustruct:ignore
	gormConfig := &gorm.Config{}

	gormDB, err := gorm.Open(sqlite.Open(string(sqliteFilePath)), gormConfig)
	if err != nil {
		return nil, &OpenDBFailedError{err}
	}

	if err := gormDB.AutoMigrate(
		//exhaustruct:ignore
		&Computer{},
	); err != nil {
		return nil, &FailedToAutoMigrateError{err}
	}

	return gormDB, nil
}
