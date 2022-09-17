package dao

import (
	"github.com/akbarsahata/alterra-agmc/day-3/lib/auth"
	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dto"
	"gorm.io/gorm"
)

type UserDAO interface {
	GetMany() ([]models.User, error)
	GetOne(params *models.User) (*models.User, error)
	GetOneByID(userID uint) (*models.User, error)
	CreateOne(data dto.CreateUserBodyDTO) (*models.User, error)
	UpdateOne(userID uint, data dto.UpdateUserBodyDTO) (*models.User, error)
	DeleteOne(userID uint) error
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{db: db}
}

func (dao *userDAO) GetMany() ([]models.User, error) {
	var users []models.User

	if err := dao.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (dao *userDAO) GetOne(params *models.User) (*models.User, error) {
	user := new(models.User)

	if err := dao.db.Where(params).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (dao *userDAO) GetOneByID(userID uint) (*models.User, error) {
	user := new(models.User)

	if err := dao.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (dao *userDAO) CreateOne(data dto.CreateUserBodyDTO) (*models.User, error) {
	password, err := auth.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: password,
	}

	if err := dao.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (dao *userDAO) UpdateOne(userID uint, data dto.UpdateUserBodyDTO) (*models.User, error) {
	err := dao.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(models.User{
			Name:  data.Name,
			Email: data.Email,
		}).Error

	if err != nil {
		return nil, err
	}

	updated := new(models.User)

	if err = dao.db.Where("id = ?", userID).First(updated).Error; err != nil {
		return nil, err
	}

	return updated, nil
}

func (dao *userDAO) DeleteOne(userID uint) error {
	if err := dao.db.Delete(&models.User{}, userID).Error; err != nil {
		return err
	}

	return nil
}
