package gorm

import (
	"fmt"
	"runtime"
)

// サポートされていない OS であることを示す例外。
type UnsupportedOSError struct {
	goos           string
	supporInFuture bool
}

func (e *UnsupportedOSError) Error() string {
	var unsupportDuration string

	if e.supporInFuture {
		unsupportDuration = "現在"
	}

	return fmt.Sprintf("%s は%s未対応です。", e.goos, unsupportDuration)
}

// Sqlite データベースファイルのパス。
type SqliteFilePath string

// SqliteFilePath のインスタンスを生成する。
func NewSqliteFilePath() (SqliteFilePath, error) {
	goos := runtime.GOOS

	switch goos {
	case "linux":
		return SqliteFilePath("/etc/switchboard/manager/data.db"), nil
	case "windows":
		return "", &UnsupportedOSError{goos: goos, supporInFuture: true}
	}

	return "", &UnsupportedOSError{goos: goos, supporInFuture: false}
}
