package config

import (
	"github.com/uma-31/switchboard/agent/domain/valueobject"
	"github.com/uma-31/switchboard/agent/infrastructure/http/gin"
)

// アプリケーションの設定情報。
type Config struct {
	ComputerInfo *valueobject.ComputerInfo
	Port         *gin.ServerPort
}
