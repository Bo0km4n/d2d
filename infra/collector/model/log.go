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
	ID          int     `gorm:"primary_key" json:"id"`
	Label       int     `json:"label"`
	UUID        string  `json:"uuid"`
	ElapsedTime uint64  `json:"elapsed_time"`
	DeltaAccelX float64 `json:"delta_accel_x"`
	DeltaAccelY float64 `json:"delta_accel_y"`
	DeltaAccelZ float64 `json:"delta_accel_z"`
	DeltaGyroX  float64 `json:"delta_gyro_x"`
	DeltaGyroY  float64 `json:"delta_gyro_y"`
	DeltaGyroZ  float64 `json:"delta_gyro_z"`
	DeltaMagX   float64 `json:"delta_mag_x"`
	DeltaMagY   float64 `json:"delta_mag_y"`
	DeltaMagZ   float64 `json:"delta_mag_z"`
}
