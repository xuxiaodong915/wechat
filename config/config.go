package config

import (
	"os"
	"time"
)

type Config struct {
	ServerPort string
	DBPath     string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/recipes.db"
	}
	return &Config{
		ServerPort: port,
		DBPath:     dbPath,
	}
}

// DailyRecommendExpiry defines how often daily recommendation rotates
const DailyRecommendExpiry = 24 * time.Hour
