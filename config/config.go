package config

import "time"

type Config struct {
	ServerPort string
	DBPath     string
}

func Load() *Config {
	return &Config{
		ServerPort: ":8080",
		DBPath:     "data/recipes.db",
	}
}

// DailyRecommendExpiry defines how often daily recommendation rotates
const DailyRecommendExpiry = 24 * time.Hour
