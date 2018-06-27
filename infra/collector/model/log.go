package model

import "time"

type Log struct {
	ID        int `gorm:"primary_key"`
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
	ID           int `gorm:"primary_key"`
	ElapsedTime  uint64
	LambdaAccelX float64
	LambdaAccelY float64
	LambdaAccelZ float64
	LambdaGyroX  float64
	LambdaGyroY  float64
	LambdaGyroZ  float64
	LambdaMagX   float64
	LambdaMagY   float64
	LambdaMagZ   float64
}
