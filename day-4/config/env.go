package config

import (
	"os"
)

var (
	Env map[string]string
)

func InitEnv() {
	once.Do(func() {
		Env = map[string]string{
			"DB_HOST":    os.Getenv("DB_HOST"),
			"DB_PORT":    os.Getenv("DB_PORT"),
			"DB_NAME":    os.Getenv("DB_NAME"),
			"DB_USER":    os.Getenv("DB_USER"),
			"DB_PASS":    os.Getenv("DB_PASS"),
			"JWT_SECRET": os.Getenv("JWT_SECRET"),
		}
	})
}
