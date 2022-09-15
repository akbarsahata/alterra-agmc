package database

import (
	"github.com/akbarsahata/alterra-agmc/day-2/config"
	"github.com/akbarsahata/alterra-agmc/day-2/models"
	"github.com/akbarsahata/alterra-agmc/day-2/models/dto"
)

func GetUsers() ([]models.UserModel, error) {
	var users []models.UserModel

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(data dto.CreateUserBodyDTO) (*models.UserModel, error) {
	user := models.UserModel{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
