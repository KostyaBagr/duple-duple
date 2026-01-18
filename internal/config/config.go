package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host      string
	User      string
	Password  string
	DB        string
	Port      string
	OutputDir string
}

// Reads .toml config file
func ReadCfgFile() (*Config, error) {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		return nil, errors.New("No config file")
	}
	var cfg Config
	_, err := toml.DecodeFile(f, &cfg)
	if err != nil {
		return nil, errors.New("Fail to decode")
	}

	return &cfg, nil
}
