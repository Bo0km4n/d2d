package model

import "github.com/jinzhu/gorm"

var labels = map[string]int{
	"joy":    1,
	"anger":  2,
	"sorrow": 3,
}

// 嬉しい
// 悲しみ
// 怒り

type Label struct {
	ID    int `gorm:"primary_key"`
	UID   int
	Label string
}

func InitLabel(db *gorm.DB) {
	for k, v := range labels {
		db.Create(&Label{
			UID:   v,
			Label: k,
		})
	}
}
