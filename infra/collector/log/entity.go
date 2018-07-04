package log

type LogParam struct {
	Time   uint64  `json:"time"`
	UUID   string  `json:"uuid"`
	Sensor *Sensor `json:"sensor"`
}

type Sensor struct {
	Accel *Accel `json:"accel"`
	Gyro  *Gyro  `json:"gyro"`
	Mag   *Mag   `json:"mag"`
}

type Accel struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Gyro struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Mag struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}
