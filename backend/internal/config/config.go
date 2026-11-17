package config

import "os"

// Config is everything the process reads from the environment.
// Defaults match deployments/local/docker-compose.yml.
type Config struct {
	Port        string
	DatabaseURL string
	Env         string
}

func Load() Config {
	return Config{
		Port:        getenv("PORT", "8080"),
		DatabaseURL: getenv("DATABASE_URL", "postgres://tawn:tawn@localhost:5433/tawn?sslmode=disable"),
		Env:         getenv("APP_ENV", "local"),
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
