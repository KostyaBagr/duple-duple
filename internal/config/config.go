package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Postgres PostgresConfig
	S3       S3Config
}

type PostgresConfig struct {
	Host      string `validate:"required"`
	User      string `validate:"required"`
	Password  string `validate:"required"`
	DB        string `validate:"required"`
	Port      string `validate:"required"`
	OutputDir string `default:"/tmp/duple-duple/postgres/"` // Now it does not work
}

type S3Config struct {
	Url             string `validate:"required"`
	BacketName      string `validate:"required"`
	AccessKey       string `validate:"required"`
	SecretAccessKey string `validate:"required"`
	Region          string `validate:"required"`
	PathInBucket    string
}

var AppConfig *Config


// Reads .toml config file
func ReadCfgFile() error {
	validate := validator.New()
	f := "config.toml"

	if _, err := os.Stat(f); err != nil {
		return fmt.Errorf("config validation failed: %w", err)

	}

	var cfg Config
	_, err := toml.DecodeFile(f, &cfg)
	if err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	err = validate.Struct(cfg)
	if err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	AppConfig = &cfg
	return nil
}
