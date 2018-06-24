package usecase

import (
	"context"

	logRepo "github.com/Bo0km4n/d2d/infra/collector/log/repository"
	"github.com/Bo0km4n/d2d/infra/collector/model"
)

// LogUC //
type LogUC interface {
	Store(ctx context.Context, item []*model.Log) error
}

type logUC struct {
	logRepo logRepo.LogRepository
}

// NewLogUsecase //
func NewLogUsecase(repo logRepo.LogRepository) LogUC {
	return &logUC{
		logRepo: repo,
	}
}

func (luc *logUC) Store(ctx context.Context, item []*model.Log) error {
	return luc.logRepo.StoreDB(ctx, item)
}
