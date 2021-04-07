package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

func Run() {
	dsn := "root:db1900@tcp(127.0.0.1:3306)/dog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// 迁移
	if err := db.AutoMigrate(NewPost(), NewPostCate()); err != nil {
		log.Fatal(err)
	}
}

type Model struct {
	gorm.Model
}
