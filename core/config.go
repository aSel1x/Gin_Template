package core

import (
	"fmt"
	"os"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Path    string
	Name    string
	Secret  string
	Version string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (p *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", p.Host, p.Port, p.User, p.Password, p.Database)
}

func NewConfig() (*Config, error) {

	appConfig := AppConfig{
		Path:    os.Getenv("APP_PATH"),
		Name:    os.Getenv("APP_NAME"),
		Secret:  os.Getenv("APP_SECRET_KEY"),
		Version: os.Getenv("APP_VERSION"),
	}

	postgresConfig := PostgresConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	config := Config{
		App:      appConfig,
		Postgres: postgresConfig,
	}

	return &config, nil
}
