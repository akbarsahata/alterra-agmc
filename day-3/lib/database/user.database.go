package database

import (
	"github.com/akbarsahata/alterra-agmc/day-3/config"
	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dto"
)

func GetUsers() ([]models.UserModel, error) {
	var users []models.UserModel

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(userID uint) (*models.UserModel, error) {
	user := new(models.UserModel)

	if err := config.DB.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
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

func UpdateUser(userID uint, data dto.UpdateUserBodyDTO) (*models.UserModel, error) {
	err := config.DB.Model(&models.UserModel{}).
		Where("id = ?", userID).
		Updates(models.UserModel{
			Name:  data.Name,
			Email: data.Email,
		}).Error

	if err != nil {
		return nil, err
	}

	updated := new(models.UserModel)

	if err = config.DB.Where("id = ?", userID).First(updated).Error; err != nil {
		return nil, err
	}

	return updated, nil
}

func DeleteUser(userID uint) error {
	if err := config.DB.Delete(&models.UserModel{}, userID).Error; err != nil {
		return err
	}

	return nil
}
