package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/KostyaBagr/duple-duple/internal/utils"
	"github.com/go-playground/validator/v10"
)

const TmpBackupPath = "/tmp/duple-duple/backup" // temporary path for dumps

// TODO: add config for notifications
// Type of storage. For version 1 it is s3 (implemented) and googleDrive (not implemented)
type StorageType int

const (
	S3 StorageType = 1 << iota
	GoogleDrive
	Local
	StLen int = iota
)

var StorageTypeName = map[StorageType]string{
	S3:          "S3",
	GoogleDrive: "googleDrive",
	Local:       "local",
}

func (st StorageType) String() string {
	return StorageTypeName[st]
}

// Type of database management system. For now it is just postgres
type DBMSType int

const (
	Postgres DBMSType = iota
)

var DBMSTypeName = map[DBMSType]string{
	Postgres: "postgres",
}

func (dt DBMSType) String() string {
	return DBMSTypeName[dt]
}

// Toml config
type Config struct {
	Postgres PostgresConfig `validate:"required"`
	Storage  StorageConfig  `validate:"required"`
}

type PostgresConfig struct {
	Host     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DB       string `validate:"required"`
	Port     string `validate:"required"`
}

type S3Config struct {
	Url             string `validate:"required"`
	BacketName      string `validate:"required"`
	AccessKey       string `validate:"required"`
	SecretAccessKey string `validate:"required"`
	Region          string `validate:"required"`
	PathInBucket    string
}

type StorageLocalConfig struct {
	Path string
}

type StorageConfig struct {
	S3    S3Config
	Local StorageLocalConfig
}

var AppConfig *Config
var SelectedStorages *[]string

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
		return fmt.Errorf("Config validation failed: %w", err)
	}

	AppConfig = &cfg

	_, err = utils.PathExists(TmpBackupPath, true)
	if err != nil {
		return fmt.Errorf("Error during creating default dir %w", err)
	}

	err = validateConfigSchema()
	if err != nil {
		return fmt.Errorf("Error in config file: %w", err)
	}
	return nil
}

// Validates toml config file
func validateConfigSchema() error {
	var selectedStorage []string


	s3CfgEmpty, _ := utils.IsEmpty(AppConfig.Storage.S3)
	localCfgEmpty, _ := utils.IsEmpty(AppConfig.Storage.Local)

	if s3CfgEmpty == true && localCfgEmpty == true {
		return errors.New("Storage config was not provided")
	}

	if !s3CfgEmpty {
		selectedStorage = append(selectedStorage, S3.String())
	}

	if !localCfgEmpty {
		isPath, err := utils.PathExists(AppConfig.Storage.Local.Path, false)
		if err != nil {
			return errors.New("error during checking the path")
		}
		if !isPath {
			return errors.New("incorrect path for local storage")
		}
		selectedStorage = append(selectedStorage, Local.String())
	}

	SelectedStorages = &selectedStorage

	return nil
}
