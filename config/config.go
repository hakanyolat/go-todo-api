package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TCP *TCPConfig
	DB  *DBConfig
}

func GetConfig() *Config {
	err := godotenv.Load("./.env")

	if err != nil {
		panic("Error loading .env file.")
	}

	return &Config{
		TCP: &TCPConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: &DBConfig{
			Dialect:  os.Getenv("DATABASE_DIALECT"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Charset:  os.Getenv("DATABASE_CHARSET"),
		},
	}
}
