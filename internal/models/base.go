package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:password@tcp(db:3306)/myapp_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connected to database")

	if !Db.HasTable(&User{}) {
		Db.CreateTable(&User{})
	}
}
