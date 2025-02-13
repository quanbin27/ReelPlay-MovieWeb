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
	GgClientID             string
	GgClientSecret         string
	GgClientCallBackURL    string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", ""),
		DSN:                    getEnv("DSN", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-anymore?"),
		EmailUsername:          getEnv("EMAIL_USERNAME", ""),
		Emailpassword:          getEnv("EMAIL_PASSWORD", ""),
		Emailfrom:              getEnv("EMAIL_FROM", ""),
		GgClientID:             getEnv("CLIENT_ID", ""),
		GgClientSecret:         getEnv("CLIENT_SECRET", ""),
		GgClientCallBackURL:    getEnv("CLIENT_CALLBACK_URL", "http://localhost:8080/auth/google/callback"),
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
