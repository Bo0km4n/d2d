package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: Detailed settings for security
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowMethods("DELETE")
	config.AddAllowHeaders("Authorization", "Cache-Control", "X-Requested-With")
	config.AllowAllOrigins = true
	return cors.New(config)
}
