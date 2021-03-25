package model

import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println(db)
	} else {

	}
}
