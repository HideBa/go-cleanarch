package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

func EnvLoad() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetConfig() (config Config) {
	env := Env()
	if env == "development" {
		EnvLoad()
	}
	config = Config{
		DBConfig: DBConfig{
			DBUrl: os.Getenv("DB_URL"),
		},
		ServerConfig: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
	return
}

func Env() string {
	env := os.Getenv("GIN_MODE")
	if (env == "") || (env == "debug") {
		env = "development"
		return env
	}
	env = "production"
	return env
}
