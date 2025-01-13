package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel                 string
	LogFileName              string
	BindAddress              string
	DatabaseConnectString    string
	OutsideServerBindAddress string
}

func New() (*Config, error) {
	if err := godotenv.Load("./configs/songLibrary.env"); err != nil {
		return nil, err
	}
	return &Config{
		LogLevel:                 getEnv("SONGLIBRARY_SERVER_LOG_LEVEL", "debug"),
		LogFileName:              getEnv("SONGLIBRARY_SERVER_LOG_FILE_NAME", "log.txt"),
		BindAddress:              getEnv("SONGLIBRARY_SERVER_BIND_ADDRESS", ":8000"),
		DatabaseConnectString:    getEnv("SONGLIBRARY_DATABASE_CONNECT_STRING", "host=localhost database=song_library port=5432 sslmode=disable user=postgres password=1234"),
		OutsideServerBindAddress: getEnv("SONGLIBRARY_OUTSIDE_SERVER_BIND_ADDRESS", "http://localhost:8001"),
	}, nil
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
