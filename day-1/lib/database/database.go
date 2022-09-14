package database

import (
	"github.com/akbarsahata/alterra-agmc/day-1/config"
	"github.com/akbarsahata/alterra-agmc/day-1/lib/dto"
	"github.com/akbarsahata/alterra-agmc/day-1/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(data dto.UserDTO) (*models.User, error) {
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
