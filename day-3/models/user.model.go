package models

import (
	"github.com/akbarsahata/alterra-agmc/day-3/models/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (um *User) ToResponseDTO() dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:    um.ID,
		Name:  um.Name,
		Email: um.Email,
	}
}
