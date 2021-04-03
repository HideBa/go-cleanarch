package config

import "os"

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

func GetConfig() (config Config) {
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
