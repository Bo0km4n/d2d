package log

type LogParam struct {
	Time   uint64  `json:"time"`
	UUID   string  `json:"macAdress"`
	Sensor *Sensor `json:"sensor"`
}

type Sensor struct {
	Accel *Accel `json:"accel"`
	Gyro  *Gyro  `json:"gyro"`
	Mag   *Mag   `json:"mag"`
}

type Accel struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Gyro struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Mag struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
