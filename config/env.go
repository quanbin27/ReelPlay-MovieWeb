package config

import (
	"github.com/lpernett/godotenv"
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DSN                    string
	JWTExpirationInSeconds int64
	JWTSecret              string
	EmailUsername          string
	Emailpassword          string
	Emailfrom              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DSN:                    getEnv("DSN", "root:27052002@tcp(127.0.0.1:3306)/ReelPlay?charset=utf8mb4&parseTime=True&loc=Local"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-anymore?"),
		EmailUsername:          getEnv("EMAIL_USERNAME", ""),
		Emailpassword:          getEnv("EMAIL_PASSWORD", ""),
		Emailfrom:              getEnv("EMAIL_FROM", ""),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return i
		}
		return fallback
	}
	return fallback
}
