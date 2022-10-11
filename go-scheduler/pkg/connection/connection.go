package connection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	USERNAME := "root"
	PASSWORD := "kush@123"
	HOST := "127.0.0.1:3306"
	DBNAME := "Interview Creation Portal"
	connectDB := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, DBNAME)
	db, err := gorm.Open("mysql", connectDB)
	fmt.Println("EE")
	if err != nil {
		panic(err)
	}
	return db
}
