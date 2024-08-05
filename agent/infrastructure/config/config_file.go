package config

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

// 設定ファイルのパス。
type FilePath string

// ConfigFilePath のインスタンスを生成する。
func NewFilePath() (FilePath, error) {
	goos := runtime.GOOS

	switch goos {
	case "linux":
		return FilePath("/etc/switchboard/agent/config.yaml"), nil
	case "windows":
		return "", &UnsupportedOSError{goos: goos, supporInFuture: true}
	}

	return "", &UnsupportedOSError{goos: goos, supporInFuture: false}
}
