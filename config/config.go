package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	TCP TCPConfig
	DB  DBConfig
}

func Get(env string) Config {
	var path string

	if env == ""{
		path = "./.env"
	}else{
		path = fmt.Sprintf("./.env.%s", strings.ToLower(env))
	}

	if err := godotenv.Load(path); err != nil {
		panic("Error loading .env file.")
	}

	return Config{
		TCP: TCPConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DBConfig{
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
