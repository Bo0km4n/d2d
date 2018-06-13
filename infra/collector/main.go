package main

import (
	logHandler "github.com/Bo0km4n/d2d/infra/collector/log/handler/http"
	logRepo "github.com/Bo0km4n/d2d/infra/collector/log/repository"
	logUC "github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	"github.com/Bo0km4n/d2d/infra/collector/middleware"
	"github.com/Bo0km4n/d2d/infra/nats"
	"github.com/gin-gonic/gin"
)

func init() {
	nats.New("localhost", "4222", "d2d")
}

func main() {
	listenAPI()
}

func listenAPI() {
	api := NewRouter()
	api.Run(":" + "8765")
}

// NewRouter //
func NewRouter() *gin.Engine {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	api.Use(middleware.Cors())

	{
		logRepo := logRepo.NewLogRepository(nats.NATSConn)
		logUC := logUC.NewLogUsecase(logRepo)
		logHandler.NewHTTPLogHandler(api, logUC)
	}

	healthCheck(api)
	return api
}

func healthCheck(api *gin.Engine) {
	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
}
