package models

import (
	"github.com/akbarsahata/alterra-agmc/day-1/lib/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (u *User) DTO() dto.UserDTO {
	return dto.UserDTO{
		Name:  u.Name,
		Email: u.Email,
	}
}
