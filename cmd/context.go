package cmd

import "github.com/KostyaBagr/duple-duple/internal/config"

var appConfig *config.Config

func InitConfig(cfg *config.Config) {
	appConfig = cfg
}
