package model

import "time"

type Log struct {
	ID        int    `gorm:"primary_key"`
	UUID      string `json:"macAddress"`
	Time      uint64
	AccelX    int
	AccelY    int
	AccelZ    int
	GyroX     int
	GyroY     int
	GyroZ     int
	MagX      int
	MagY      int
	MagZ      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AggregatedLog struct {
	ID          int `gorm:"primary_key"`
	Label       string
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
