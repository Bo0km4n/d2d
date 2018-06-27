package db

import (
	"github.com/Bo0km4n/d2d/infra/collector/model"
)

func Migrate() {
	DB.AutoMigrate(
		&model.Log{},
		&model.AggregatedLog{},
	)
}
