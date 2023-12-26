package config

import (
	"web_service/internal/apperrors"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var path = "configs/.env"

type Config struct {
	AppPort  string `required:"true" split_words:"true"`
	Postgres *PostgresConfig
}

type PostgresConfig struct {
	AppPort      string `env:"APP_PORT"`
	LogLevel     string `env:"LOGGER_LEVEL"`
	SqlHost      string `env:"SQL_HOST"`
	SqlPort      string `env:"SQL_PORT"`
	SqlType      string `env:"SQL_TYPE"`
	SqlMode      string `env:"SQL_MODE"`
	UserName     string `env:"USER_NAME"`
	Password     string `env:"PASSWORD"`
	DBName       string `env:"DB_NAME"`
	TimeZone     string `env:"TIME_ZONE"`
	TimeoutQuery string `env:"TIMEOUT_QUERY"`
}

func NewConfig(logger *zap.Logger) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		appErr := apperrors.EnvConfigLoadError.AppendMessage(err)
		return nil, appErr
	}

	confPsql := PostgresConfig{}
	if err := env.Parse(&confPsql); err != nil {
		appErr := apperrors.EnvConfigParseError.AppendMessage(err)
		return nil, appErr
	}

	conf := Config{AppPort: confPsql.AppPort, Postgres: &confPsql}

	logger.Sugar().Info("Config has been parsed")
	return &conf, nil
}
