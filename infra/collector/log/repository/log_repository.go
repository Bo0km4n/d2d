package repository

import (
	"context"
	"encoding/json"

	"github.com/Bo0km4n/d2d/infra/nats"
)

// LogRepository interface
type LogRepository interface {
	StoreToNATS(ctx context.Context, item interface{}) error
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

func (lr *logRepository) StoreToNATS(ctx context.Context, item interface{}) error {
	body, err := json.Marshal(&item)
	if err != nil {
		return err
	}
	return lr.NATSConn.Publish(body)
}
