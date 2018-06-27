package usecase

import (
	"context"
	"math"

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

func compressLog(logs []*model.Log) *model.AggregatedLog {
	var maxTime uint64
	var minTime uint64

	var maxAccelX int
	var maxAccelY int
	var maxAccelZ int

	var minAccelX int
	var minAccelY int
	var minAccelZ int

	var maxGyroX int
	var maxGyroY int
	var maxGyroZ int

	var minGyroX int
	var minGyroY int
	var minGyroZ int

	var maxMagX int
	var maxMagY int
	var maxMagZ int

	var minMagX int
	var minMagY int
	var minMagZ int

	var lambdaAccelX float64
	var lambdaAccelY float64
	var lambdaAccelZ float64

	var lambdaGyroX float64
	var lambdaGyroY float64
	var lambdaGyroZ float64

	var lambdaMagX float64
	var lambdaMagY float64
	var lambdaMagZ float64

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

	lambdaAccelX = math.Abs(float64(maxAccelX - minAccelX))
	lambdaAccelY = math.Abs(float64(maxAccelY - minAccelY))
	lambdaAccelZ = math.Abs(float64(maxAccelZ - minAccelZ))

	lambdaGyroX = math.Abs(float64(maxGyroX - minGyroX))
	lambdaGyroY = math.Abs(float64(maxGyroY - minGyroY))
	lambdaGyroZ = math.Abs(float64(maxGyroZ - minGyroZ))

	lambdaMagX = math.Abs(float64(maxMagX - minMagX))
	lambdaMagY = math.Abs(float64(maxMagY - minMagY))
	lambdaMagZ = math.Abs(float64(maxMagZ - minMagZ))

	return &model.AggregatedLog{
		ElapsedTime:  maxTime - minTime,
		LambdaAccelX: lambdaAccelX,
		LambdaAccelY: lambdaAccelY,
		LambdaAccelZ: lambdaAccelZ,
		LambdaGyroX:  lambdaGyroX,
		LambdaGyroY:  lambdaGyroY,
		LambdaGyroZ:  lambdaGyroZ,
		LambdaMagX:   lambdaMagX,
		LambdaMagY:   lambdaMagY,
		LambdaMagZ:   lambdaMagZ,
	}

	return nil
}
