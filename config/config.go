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
	if os.Getenv("GOPATH") == "" {
		panic("GOPATH not found in environment")
	}

	var path string
	if env == ""{
		path = "./.env"
	}else{
		path = fmt.Sprintf("./.env.%s", strings.ToLower(env))
	}

	fullPath := fmt.Sprintf("$GOPATH/src/github.com/hakanyolat/go-todo-api/%s", path)
	if err := godotenv.Load(os.ExpandEnv(fullPath)); err != nil {
		panic(fmt.Sprintf("Error loading .env file. \".env.%s\" not found.", strings.ToLower(env)))
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
