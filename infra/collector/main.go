package main

import (
	"log"

	"github.com/Bo0km4n/d2d/infra/collector/db"
	logHandler "github.com/Bo0km4n/d2d/infra/collector/log/handler/http"
	logRepo "github.com/Bo0km4n/d2d/infra/collector/log/repository"
	logUC "github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	"github.com/Bo0km4n/d2d/infra/collector/middleware"
	"github.com/Bo0km4n/d2d/infra/nats"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	nats.New("localhost", "4222", "d2d")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("local")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config err", err)
	}

	db.MySQL()
	db.Migrate()
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
		logRepo := logRepo.NewLogRepository(nats.NATSConn, db.DB)
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
