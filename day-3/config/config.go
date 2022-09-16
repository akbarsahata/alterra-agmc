package config

import (
	"fmt"
	"os"

	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	config := map[string]string{
		"DB_HOST": os.Getenv("DB_HOST"),
		"DB_PORT": os.Getenv("DB_PORT"),
		"DB_NAME": os.Getenv("DB_NAME"),
		"DB_USER": os.Getenv("DB_USER"),
		"DB_PASS": os.Getenv("DB_PASS"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USER"],
		config["DB_PASS"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)

	if DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{}); err != nil {
		panic(nil)
	}

	fmt.Println("Connected to MySQL")
}

func InitMigration() {
	DB.AutoMigrate(&models.UserModel{})

	fmt.Println("Migrations executed successfully")
}
