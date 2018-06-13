package main

import (
	logHandler "github.com/Bo0km4n/d2d/infra/collector/log/handler/http"
	"github.com/Bo0km4n/d2d/infra/collector/log/repository"
	"github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	"github.com/Bo0km4n/d2d/infra/nats"
	"github.com/gin-gonic/gin"
)

func main() {
	listenAPI()
}

func listenAPI() {
	api := NewRouter()
	api.Run(":" + "8765")
}

func NewRouter() *gin.Engine {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	nats.New("localhost", "4222", "d2d")

	{
		logRepo := repository.NewLogRepository(nats.NATSConn)
		logUC := usecase.NewLogUsecase(logRepo)
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
