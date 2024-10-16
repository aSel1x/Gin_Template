package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	Rabbit   RabbitConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Path    string `mapstructure:"APP_PATH"`
	Name    string `mapstructure:"APP_NAME"`
	Secret  string `mapstructure:"APP_SECRET"`
	Version string `mapstructure:"APP_VERSION"`
}

type PostgresConfig struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int    `mapstructure:"POSTGRES_PORT"`
	Database string `mapstructure:"POSTGRES_DB"`
}

type RabbitConfig struct {
	User     string `mapstructure:"RABBIT_USER"`
	Password string `mapstructure:"RABBIT_PASSWORD"`
	Host     string `mapstructure:"RABBIT_HOST"`
	Port     int    `mapstructure:"RABBIT_PORT"`
	Path     string `mapstructure:"RABBIT_PATH"`
}

type RedisConfig struct {
	User     string `mapstructure:"REDIS_USER"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	SSL      bool   `mapstructure:"REDIS_SSL"`
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	DB       int    `mapstructure:"REDIS_DB"`
}

// NewConfig reads configuration from environment variables or a config file.
func NewConfig() (*Config, error) {
	viper.SetConfigName("env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("app.path", "/api")
	viper.SetDefault("app.name", "Template")
	viper.SetDefault("app.secret", "123abc")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("rabbit.port", 5672)
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.ssl", false)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, relying on environment variables: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	return &config, nil
}

func (p PostgresConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", p.User, p.Password, p.Host, p.Port, p.Database)
}

func (r RabbitConfig) DSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", r.User, r.Password, r.Host, r.Port, r.Path)
}

func (r RedisConfig) DSN() string {
	scheme := "redis"
	if r.SSL {
		scheme = "rediss"
	}
	return fmt.Sprintf("%s://%s:%s@%s:%d/%d", scheme, r.User, r.Password, r.Host, r.Port, r.DB)
}
