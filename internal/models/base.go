package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/mysql")
	fmt.Println("connected to database")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
