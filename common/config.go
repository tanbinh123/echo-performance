package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var Db *gorm.DB

func InitDB() {
	db, _ := gorm.Open("mysql", os.Getenv("DB_SOURCE"))
	defer db.Close()

	Db = db
}
