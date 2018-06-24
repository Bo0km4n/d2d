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
