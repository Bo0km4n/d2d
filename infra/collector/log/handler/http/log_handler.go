package http

import (
	"context"
	"log"
	"net/http"

	logUC "github.com/Bo0km4n/d2d/infra/collector/log/usecase"
	"github.com/k0kubun/pp"

	"github.com/gin-gonic/gin"
)

// HTTPLogHandler //
type HTTPLogHandler struct {
	logUC logUC.LogUC
}

// Store //
func (h *HTTPLogHandler) Store(c *gin.Context) {
	var item interface{}

	if err := c.BindJSON(&item); err != nil {
		log.Printf("Bind json error: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	pp.Println(item)

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
