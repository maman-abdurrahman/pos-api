package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	DBConfig DBConfig
}

type ServerConfig struct {
	AppName  string
	Hostname string
	Port     string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func GetAppConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		Server: ServerConfig{
			AppName:  os.Getenv("APPNAME"),
			Hostname: os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
		},
		DBConfig: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("SSLMode"),
		},
	}
}
