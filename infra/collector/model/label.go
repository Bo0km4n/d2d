package model

import "github.com/jinzhu/gorm"

var labels = map[string]int{
	"negative": 1,
	"positive": 2,
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
