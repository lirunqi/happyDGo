package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:wobuhuiwan@tcp(127.0.0.1:3306)/GO_API?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
