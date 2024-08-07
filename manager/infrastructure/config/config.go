package config

import "github.com/uma-31/switchboard/manager/infrastructure/http/gin"

// アプリケーションの設定情報。
type Config struct {
	Port *gin.ServerPort
}
