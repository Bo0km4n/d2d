package repository

import (
	"context"

	"github.com/Bo0km4n/d2d/infra/collector/model"
	"github.com/Bo0km4n/d2d/infra/nats"
	"github.com/jinzhu/gorm"
)

// LogRepository interface
type LogRepository interface {
	StoreDB(ctx context.Context, item []*model.Log) error
}

type logRepository struct {
	NATSConn *nats.NATS
	DB       *gorm.DB
}

// NewLogRepository //
func NewLogRepository(conn *nats.NATS, db *gorm.DB) LogRepository {
	logRepo := logRepository{
		NATSConn: conn,
		DB:       db,
	}
	return &logRepo
}

func (lr *logRepository) StoreDB(ctx context.Context, item []*model.Log) error {
	for i := range item {
		if err := lr.DB.Create(item[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
