package controllers

import "gorm.io/gorm"

type UserController struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) UserController {
	controller := new(UserController)

	controller.DB = DB

	return *controller
}
