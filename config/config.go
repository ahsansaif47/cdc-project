package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl        string
	ServiceName  string
	ServerHost   string
	ServerPort   string
	DatabaseName string
	RedisUrl     string
	JWTSecret    string
}

var Cfg Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		instance, err := loadConfig()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		Cfg = instance
	})
	return Cfg
}

func loadConfig() (Config, error) {
	err := godotenv.Load(filepath.Join("..", "..", ".env"))

	return Config{
		ServiceName:  os.Getenv("SERVICE_NAME"),
		ServerHost:   os.Getenv("SERVER"),
		ServerPort:   os.Getenv("PORT"),
		DatabaseName: os.Getenv("SERVICE_DATABASE_NAME"),
	}, err

}
