package config

import (
	"os"
)

type AppConfig struct {
	JWTSecret string
}

func Load() AppConfig {
	return AppConfig{
		JWTSecret: getEnv("JWT_SECRET", "change-me"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}




