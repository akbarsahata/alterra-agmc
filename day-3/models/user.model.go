package models

import (
	"github.com/akbarsahata/alterra-agmc/day-3/models/dto"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (um *UserModel) ToResponseDTO() dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:    um.ID,
		Name:  um.Name,
		Email: um.Email,
	}
}
