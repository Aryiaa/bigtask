package main

import (
	"app/conf"
	"app/model"

	_"github.com/jinzhu/gorm"
)

func main()  {
	conf.Initconf()
	model.Init()
	model.DB().AutoMigrate(&model.BigdataUser{})
}
