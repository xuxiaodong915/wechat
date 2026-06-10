package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	return &Config{
		ServerPort: port,
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", ""),
		DBName:     getEnv("DB_NAME", "recipes"),
	}
}

// DSN returns MySQL connection string
func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// DailyRecommendExpiry defines how often daily recommendation rotates
const DailyRecommendExpiry = 24 * time.Hour
