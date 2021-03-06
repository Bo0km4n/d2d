package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var (
	HOST         string
	PORT         string
	USER         string
	PASSWORD     string
	DATABASE     string
	DATABASE_URI string
)

// InitMySQL //
func InitMySQL(host, port, user, password, database string) {
	HOST = host
	PORT = port
	USER = user
	PASSWORD = password
	DATABASE = database
	DATABASE_URI = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		USER, PASSWORD, HOST, PORT, DATABASE,
	)
	gormDB, err := gorm.Open("mysql", DATABASE_URI)
	if err != nil {
		fmt.Println("MYSQL ERROR:", err)
		os.Exit(1)
	}
	gormDB.LogMode(true) // TODO: remove
	gormDB.SingularTable(true)
	// NOTE デフォルトでは関連テーブルの保存を許可しない
	gormDB = gormDB.Set("gorm:save_associations", false)

	DB = gormDB
}
