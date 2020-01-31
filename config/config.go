package config

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

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "localhost",
			Port:     "5432",
			Dialect:  "postgres",
			Username: "postgres",
			Password: "",
			Name:     "kubbe_dev",
			SSLMode:  "disable",
		},
		App: &AppConfig{
			Title: "Kubbe",
			Host:  "localhost:4000",
		},
	}
}
