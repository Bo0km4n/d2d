package usecase

import (
	"context"
	"math"

	aggregatedLogRepo "github.com/Bo0km4n/d2d/infra/collector/aggregatedlog/repository"
	logRepo "github.com/Bo0km4n/d2d/infra/collector/log/repository"
	"github.com/Bo0km4n/d2d/infra/collector/model"
)

// LogUC //
type LogUC interface {
	Store(ctx context.Context, item []*model.Log) error
}

type logUC struct {
	logRepo           logRepo.LogRepository
	aggregatedLogRepo aggregatedLogRepo.AggregatedLogRepository
}

// NewLogUsecase //
func NewLogUsecase(repo logRepo.LogRepository, aggrepo aggregatedLogRepo.AggregatedLogRepository) LogUC {
	return &logUC{
		logRepo:           repo,
		aggregatedLogRepo: aggrepo,
	}
}

func (luc *logUC) Store(ctx context.Context, item []*model.Log) error {
	return luc.aggregatedLogRepo.Store(ctx, compressLog(item))
	// return luc.logRepo.StoreDB(ctx, item)
}

func compressLog(logs []*model.Log) *model.AggregatedLog {
	var maxTime uint64
	var minTime uint64

	var maxAccelX float64
	var maxAccelY float64
	var maxAccelZ float64

	var minAccelX float64
	var minAccelY float64
	var minAccelZ float64

	var maxGyroX float64
	var maxGyroY float64
	var maxGyroZ float64

	var minGyroX float64
	var minGyroY float64
	var minGyroZ float64

	var maxMagX float64
	var maxMagY float64
	var maxMagZ float64

	var minMagX float64
	var minMagY float64
	var minMagZ float64

	var deltaAccelX float64
	var deltaAccelY float64
	var deltaAccelZ float64

	var deltaGyroX float64
	var deltaGyroY float64
	var deltaGyroZ float64

	var deltaMagX float64
	var deltaMagY float64
	var deltaMagZ float64

	for _, l := range logs {
		if maxTime < l.Time {
			maxTime = l.Time
		}
		if minTime > l.Time {
			minTime = l.Time
		}

		if l.AccelX > maxAccelX {
			maxAccelX = l.AccelX
		}
		if l.AccelX < minAccelX {
			minAccelX = l.AccelX
		}
		if l.AccelY > maxAccelY {
			minAccelX = l.AccelX
		}
		if l.AccelY < minAccelY {
			minAccelY = l.AccelY
		}
		if l.AccelZ > maxAccelZ {
			maxAccelZ = l.AccelZ
		}
		if l.AccelZ < minAccelZ {
			minAccelZ = l.AccelZ
		}

		if l.GyroX > maxGyroX {
			maxGyroX = l.GyroX
		}
		if l.GyroX < minGyroX {
			minGyroX = l.GyroX
		}
		if l.GyroY > maxGyroY {
			minGyroX = l.GyroX
		}
		if l.GyroY < minGyroY {
			minGyroY = l.GyroY
		}
		if l.GyroZ > maxGyroZ {
			maxGyroZ = l.GyroZ
		}
		if l.GyroZ < minGyroZ {
			minGyroZ = l.GyroZ
		}

		if l.MagX > maxMagX {
			maxMagX = l.MagX
		}
		if l.MagX < minMagX {
			minMagX = l.MagX
		}
		if l.MagY > maxMagY {
			minMagX = l.MagX
		}
		if l.MagY < minMagY {
			minMagY = l.MagY
		}
		if l.MagZ > maxMagZ {
			maxMagZ = l.MagZ
		}
		if l.MagZ < minMagZ {
			minMagZ = l.MagZ
		}
	}

	deltaAccelX = math.Abs(float64(maxAccelX - minAccelX))
	deltaAccelY = math.Abs(float64(maxAccelY - minAccelY))
	deltaAccelZ = math.Abs(float64(maxAccelZ - minAccelZ))

	deltaGyroX = math.Abs(float64(maxGyroX - minGyroX))
	deltaGyroY = math.Abs(float64(maxGyroY - minGyroY))
	deltaGyroZ = math.Abs(float64(maxGyroZ - minGyroZ))

	deltaMagX = math.Abs(float64(maxMagX - minMagX))
	deltaMagY = math.Abs(float64(maxMagY - minMagY))
	deltaMagZ = math.Abs(float64(maxMagZ - minMagZ))

	return &model.AggregatedLog{
		UUID:        logs[0].UUID,
		ElapsedTime: maxTime - minTime,
		DeltaAccelX: deltaAccelX,
		DeltaAccelY: deltaAccelY,
		DeltaAccelZ: deltaAccelZ,
		DeltaGyroX:  deltaGyroX,
		DeltaGyroY:  deltaGyroY,
		DeltaGyroZ:  deltaGyroZ,
		DeltaMagX:   deltaMagX,
		DeltaMagY:   deltaMagY,
		DeltaMagZ:   deltaMagZ,
	}
}
