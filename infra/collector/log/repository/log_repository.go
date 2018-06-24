package repository

import (
	"context"
	"encoding/json"

	"github.com/Bo0km4n/d2d/infra/collector/model"
	"github.com/Bo0km4n/d2d/infra/nats"
)

// LogRepository interface
type LogRepository interface {
	StoreDB(ctx context.Context, item *model.Log) error
}

type logRepository struct {
	NATSConn *nats.NATS
}

// NewLogRepository //
func NewLogRepository(conn *nats.NATS) LogRepository {
	logRepo := logRepository{
		NATSConn: conn,
	}
	return &logRepo
}

func (lr *logRepository) StoreDB(ctx context.Context, item *model.Log) error {
	body, err := json.Marshal(&item)
	if err != nil {
		return err
	}
	return lr.NATSConn.Publish(body)
}
