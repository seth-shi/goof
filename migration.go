package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


func RegisterMigrate(Up func(gorm.DB) error, Down func(gorm.DB) error) {

	fmt.Println(Up, Down)
}