package repository

import (
	"context"

	"github.com/Bo0km4n/d2d/infra/collector/model"
	"github.com/jinzhu/gorm"
)

// LogRepository interface
type LogRepository interface {
	StoreDB(ctx context.Context, item []*model.Log) error
	Store(ctx context.Context, item *model.AggregatedLog) error
}

type logRepository struct {
	DB *gorm.DB
}

// NewLogRepository //
func NewLogRepository(db *gorm.DB) LogRepository {
	logRepo := logRepository{
		DB: db,
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

func (lr *logRepository) Store(ctx context.Context, item *model.AggregatedLog) error {
	return lr.DB.Create(item).Error
}
