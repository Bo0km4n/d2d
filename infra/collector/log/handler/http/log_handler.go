package http

import (
	"context"
	"log"
	"net/http"

	logEntity "github.com/Bo0km4n/d2d/infra/collector/log"
	logUC "github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	"github.com/Bo0km4n/d2d/infra/collector/model"
	"github.com/k0kubun/pp"

	"github.com/gin-gonic/gin"
)

// HTTPLogHandler //
type HTTPLogHandler struct {
	logUC logUC.LogUC
}

// Store //
func (h *HTTPLogHandler) Store(c *gin.Context) {
	var param logEntity.LogParam

	if err := c.BindJSON(&param); err != nil {
		log.Printf("Bind json error: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	pp.Println(param)
	item := convertParamToLog(&param)

	if err := h.logUC.Store(ctx, item); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

// NewHTTPLogHandler //
func NewHTTPLogHandler(g *gin.Engine, uc logUC.LogUC) {
	h := &HTTPLogHandler{
		logUC: uc,
	}
	v1 := g.Group("/api/v1")
	v1.GET("/log", h.Store)
}

func convertParamToLog(param *logEntity.LogParam) *model.Log {
	return &model.Log{
		AccelX: param.Sensor.Accel.X,
		AccelY: param.Sensor.Accel.Y,
		AccelZ: param.Sensor.Accel.Z,
		GyroX:  param.Sensor.Gyro.X,
		GyroY:  param.Sensor.Gyro.Y,
		GyroZ:  param.Sensor.Gyro.Z,
		MagX:   param.Sensor.Mag.X,
		MagY:   param.Sensor.Mag.Y,
		MagZ:   param.Sensor.Mag.Z,
		Time:   param.Time,
	}
}
