package config

import "os"

type Config struct {
	DB  *DBConfig
	App *AppConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Dialect  string
	Username string
	Password string
	Name     string
	SSLMode  string
}

type AppConfig struct {
	Title      string
	Host       string
	MapsAPIKey string
}

// GetConfig checks if env has the variables required, if not fallbacks to default values
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     getEnvOrDefault("KUBBE_DB_HOST", "localhost"),
			Port:     getEnvOrDefault("KUBBE_DB_PORT", "5432"),
			Username: getEnvOrDefault("KUBBE_DB_USERNAME", "postgres"),
			Password: getEnvOrDefault("KUBBE_DB_PASSWORD", ""),
			Name:     getEnvOrDefault("KUBBE_DB_NAME", "kubbe_dev"),
			Dialect:  "postgres",
			SSLMode:  "disable",
		},
		App: &AppConfig{
			Title: getEnvOrDefault("KUBBE_APP_TITLE", "Kubbe"),
			Host:  getEnvOrDefault("KUBBE_APP_HOST", "localhost:4000"),
		},
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
