package config

import (
	"fmt"
	"sync"

	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		if db == nil {
			InitDB()
		}
	})

	return db
}

func InitDB() {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Env["DB_USER"],
		Env["DB_PASS"],
		Env["DB_HOST"],
		Env["DB_PORT"],
		Env["DB_NAME"],
	)

	if db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		panic(nil)
	}

	fmt.Println("Connected to MySQL")
}

func InitMigration() {
	db.AutoMigrate(&models.User{})

	fmt.Println("Migrations executed successfully")
}
