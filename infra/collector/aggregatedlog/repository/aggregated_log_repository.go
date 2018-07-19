package repository

import (
	"context"

	"github.com/Bo0km4n/d2d/infra/collector/model"
	"github.com/jinzhu/gorm"
)

type AggregatedLogRepository interface {
	Store(ctx context.Context, item *model.AggregatedLog) (*model.AggregatedLog, error)
}

type aggregatedLogRepository struct {
	DB *gorm.DB
}

func NewAggregatedLogRepository(db *gorm.DB) AggregatedLogRepository {
	repo := aggregatedLogRepository{
		DB: db,
	}

	return &repo
}

func (alr *aggregatedLogRepository) Store(ctx context.Context, item *model.AggregatedLog) (*model.AggregatedLog, error) {
	err := alr.DB.Create(item).Error
	return item, err
}
