package config

import (
	"github.com/lpernett/godotenv"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DSN        string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DSN:        getEnv("DSN", "root:27052002@tcp(127.0.0.1:3306)/ReelPlay?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
