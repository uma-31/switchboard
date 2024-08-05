package config

import (
	"fmt"
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
	"github.com/uma-31/switchboard/agent/infrastructure/http/gin"
	"gopkg.in/yaml.v3"
)

type rawConfig struct {
	Port     uint16 `yaml:"port"`
	Computer struct {
		ID   string `yaml:"id"`
		Name string `yaml:"name"`
	} `yaml:"computer"`
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

// デフォルトのコンピューター ID の生成に失敗したことを示す例外。
type FailedToGenerateDefaultComputerIDError struct {
	cause error
}

func (e *FailedToGenerateDefaultComputerIDError) Error() string {
	return "デフォルトのコンピューター ID の生成に失敗しました: " + e.cause.Error()
}

// 設定情報が不正であることを示す例外。
type InvalidConfigError struct {
	filePath FilePath
	name     string
	cause    error
}

func (e *InvalidConfigError) Error() string {
	return fmt.Sprintf(
		"設定項目 '%s' が不正です(設定ファイル: '%s'): %s",
		e.filePath,
		e.name,
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

	computerConfig := &raw.Computer

	if computerConfig.ID == "" {
		id, err := machineid.ID()
		if err != nil {
			return nil, &FailedToGenerateDefaultComputerIDError{err}
		}

		computerConfig.ID = id
	}

	if computerConfig.Name == "" {
		computerConfig.Name = computerConfig.ID
	}

	computerID, err := valueobject.NewComputerID(computerConfig.ID)
	if err != nil {
		return nil, &InvalidConfigError{configFilePath, "computer.id", err}
	}

	computerName, err := valueobject.NewComputerName(computerConfig.Name)
	if err != nil {
		return nil, &InvalidConfigError{configFilePath, "computer.name", err}
	}

	return &Config{
		ComputerInfo: valueobject.NewComputerInfo(computerID, computerName),
		Port:         gin.NewServerPort(raw.Port),
	}, nil
}
