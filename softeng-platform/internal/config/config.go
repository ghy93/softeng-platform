package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		// ✅ 修改默认 URL 为 MySQL 格式
		DatabaseURL: getEnv("DATABASE_URL", "root:password@tcp(127.0.0.1:3306)/softeng?parseTime=true&loc=Local"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
