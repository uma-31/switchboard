package config

import (
	"fmt"
	"os"

	"github.com/uma-31/switchboard/manager/infrastructure/http/gin"
	"gopkg.in/yaml.v3"
)

type rawConfig struct {
	Port uint16 `yaml:"port"`
}

// 設定情報の読み込みに失敗したことを示す例外。
type FailedToLoadConfigError struct {
	filePath FilePath
	cause    error
}

func (e *FailedToLoadConfigError) Error() string {
	return fmt.Sprintf(
		"設定ファイル '%s' の読み込みに失敗しました: %s",
		e.filePath,
		e.cause.Error(),
	)
}

// 設定情報を取得する。
func Load(configFilePath FilePath) (*Config, error) {
	var raw rawConfig

	data, err := os.ReadFile(string(configFilePath))
	if err != nil {
		return nil, &FailedToLoadConfigError{configFilePath, err}
	}

	err = yaml.Unmarshal(data, &raw)
	if err != nil {
		return nil, &FailedToLoadConfigError{configFilePath, err}
	}

	return &Config{
		Port: gin.NewServerPort(raw.Port),
	}, nil
}
