package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Postgres PostgresConfig
	S3 S3Config
}

type PostgresConfig struct {
	Host      string
	User      string
	Password  string
	DB        string
	Port      string
	OutputDir string
}

type S3Config struct {
	Url string
	BacketName string
	AccessKey string
	SecretAccessKey string
	Region string
}
//  TODO: add here a method whuch will form connection url


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
