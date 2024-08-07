package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uma-31/switchboard/manager/adapter/http/controller"
)

// Gin を使用して提供される HTTP サーバーのルーター。
type Router struct {
	computersController *controller.ComputersController
	computerController  *controller.ComputerController
}

// Router のインスタンスを生成する。
func NewRouter(
	computersController *controller.ComputersController,
	computerController *controller.ComputerController,
) *Router {
	return &Router{
		computersController,
		computerController,
	}
}

// 各エンドポイントとコントローラーの関連付けを行う。
func (r *Router) Register(e *gin.Engine) {
	computersRouter := e.Group("/computers")
	{
		computersRouter.GET("/", func(ctx *gin.Context) {
			computers, err := r.computersController.GetComputers()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"computers": computers})
		})

		computersRouter.GET("/scan", func(ctx *gin.Context) {
			computers, err := r.computersController.ScanComputers()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"computers": computers})
		})

		computersRouter.POST("/scan-and-save", func(ctx *gin.Context) {
			computers, err := r.computersController.ScanAndSaveComputers()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"computers": computers})
		})

		computerRouter := computersRouter.Group("/:computerID")
		{
			computerRouter.POST("/wake", func(ctx *gin.Context) {
				computerID := ctx.Param("computerID")

				if err := r.computerController.WakeComputer(computerID); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

					return
				}

				ctx.Status(http.StatusOK)
			})
		}
	}
}
