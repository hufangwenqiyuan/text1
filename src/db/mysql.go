package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:root:test_dome?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	/**
	判断是否链接正常
	*/
	if err != nil {
		fmt.Printf("mysql connect err %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
