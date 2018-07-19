package http

import (
	"context"
	"log"
	"net/http"

	logEntity "github.com/Bo0km4n/d2d/infra/collector/log"
	logUC "github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	mlRepo "github.com/Bo0km4n/d2d/infra/collector/ml/repository"
	"github.com/Bo0km4n/d2d/infra/collector/model"

	"github.com/gin-gonic/gin"
)

// HTTPLogHandler //
type HTTPLogHandler struct {
	logUC logUC.LogUC
}

// Store //
func (h *HTTPLogHandler) Store(c *gin.Context) {
	var param []*logEntity.LogParam

	if err := c.BindJSON(&param); err != nil {
		log.Printf("Bind json error: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	item := convertParamToLogs(param)

	logs, err := h.logUC.Store(ctx, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	preRepo := mlRepo.NewMlRepository()
	preResult, err := preRepo.Predict(logs)

	log.Printf("Emotion: %+v", preResult)
	c.JSON(http.StatusOK, preResult)
}

// NewHTTPLogHandler //
func NewHTTPLogHandler(g *gin.Engine, uc logUC.LogUC) {
	h := &HTTPLogHandler{
		logUC: uc,
	}
	v1 := g.Group("/api/v1")
	v1.POST("/log", h.Store)
}

func convertParamToLogs(param []*logEntity.LogParam) []*model.Log {
	logs := make([]*model.Log, 0)

	for i := range param {
		v := &model.Log{
			UUID:   param[i].UUID,
			AccelX: param[i].Sensor.Accel.X,
			AccelY: param[i].Sensor.Accel.Y,
			AccelZ: param[i].Sensor.Accel.Z,
			GyroX:  param[i].Sensor.Gyro.X,
			GyroY:  param[i].Sensor.Gyro.Y,
			GyroZ:  param[i].Sensor.Gyro.Z,
			MagX:   param[i].Sensor.Mag.X,
			MagY:   param[i].Sensor.Mag.Y,
			MagZ:   param[i].Sensor.Mag.Z,
			Time:   param[i].Time,
		}
		logs = append(logs, v)
	}
	return logs
}
