package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uma-31/switchboard/agent/adapter/http/controller"
)

// Gin を使用して提供される HTTP サーバーのルーター。
type Router struct {
	computerInfoController *controller.ComputerInfoController
}

// Router のインスタンスを生成する。
func NewRouter(computerInfoController *controller.ComputerInfoController) *Router {
	return &Router{computerInfoController}
}

// 各エンドポイントとコントローラーの関連付けを行う。
func (r *Router) Register(e *gin.Engine) {
	e.GET("/info", func(ctx *gin.Context) {
		computerInfoDTO, err := r.computerInfoController.GetComputerInfo()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, computerInfoDTO)
	})
}
