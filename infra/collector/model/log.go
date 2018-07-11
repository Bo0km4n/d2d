package model

import "time"

type Log struct {
	ID        float64 `gorm:"primary_key"`
	UUID      string  `json:"macAddress"`
	Time      uint64
	AccelX    float64
	AccelY    float64
	AccelZ    float64
	GyroX     float64
	GyroY     float64
	GyroZ     float64
	MagX      float64
	MagY      float64
	MagZ      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AggregatedLog struct {
	ID          int `gorm:"primary_key"`
	Label       int
	UUID        string
	ElapsedTime uint64
	DeltaAccelX float64
	DeltaAccelY float64
	DeltaAccelZ float64
	DeltaGyroX  float64
	DeltaGyroY  float64
	DeltaGyroZ  float64
	DeltaMagX   float64
	DeltaMagY   float64
	DeltaMagZ   float64
}
