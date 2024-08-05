package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Gin の HTTP サーバーが使用するポート番号。
type ServerPort struct {
	value uint16
}

// ServerPort のインスタンスを生成する。
func NewServerPort(value uint16) *ServerPort {
	// NOTE: ここでバリデーションをすると良い。
	return &ServerPort{value}
}

// ポート番号を取得する。
func (p *ServerPort) Value() uint16 {
	return p.value
}

// Gin を使用して提供される HTTP サーバー。
type Server struct {
	engine *gin.Engine
	port   *ServerPort
}

// Server のインスタンスを生成する。
func NewServer(router *Router, port *ServerPort) (*Server, error) {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	router.Register(engine)

	return &Server{engine, port}, nil
}

// サーバーを起動する。
func (s *Server) Run() error {
	port := s.port.Value()

	if err := s.engine.Run(
		fmt.Sprintf(":%d", port),
	); err != nil {
		return fmt.Errorf("サーバーの起動に失敗しました: %w", err)
	}

	return nil
}
