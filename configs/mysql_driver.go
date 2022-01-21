package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetMySqlConnection() *gorm.DB {
	dsn := os.Getenv("mySqlConnect")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err) //log.fatall neu co se dung chuong trinh luon - khac fmt o la co in them thoi gian loi
	}

	return db
}
