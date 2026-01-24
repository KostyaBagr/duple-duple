/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/KostyaBagr/duple-duple/internal/backup"
	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/spf13/cobra"
)

var dbms, storage string

// TODO: add validation on DBMS and STORAGE based on enum mb?

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your db",
	Long:  "TODO: change it!",
	RunE: func(cmd *cobra.Command, args []string) error {
		possibleStorages := []string{
			cfg.Local.String(),
			cfg.S3.String(),
			cfg.GoogleDrive.String(),
		}
		if !slices.Contains(possibleStorages, storage) {
			return fmt.Errorf(
				"Invalid storage type was provided, please choose %s",
				possibleStorages,
			)
		}
		possibleDBMS := []string{
			cfg.Postgres.String(),
		}
		if !slices.Contains(possibleDBMS, dbms) {
			return fmt.Errorf("Invalid DBMS was provided, please choose %s", possibleDBMS)
		}

		storageTypes := strings.Split(storage, "/")
		if dbms == cfg.Postgres.String() {
			backup.PostgresDump(
				cfg.AppConfig.Postgres.Host,
				cfg.AppConfig.Postgres.User,
				cfg.AppConfig.Postgres.Password,
				cfg.AppConfig.Postgres.DB,
				cfg.AppConfig.Postgres.Port,
				storageTypes,
			)
		}
		return nil
	},
}

func init() {
	backupCmd.Flags().StringVar(&dbms, "dbms", "", "DBMS like postgres, mysql, etc")
	if err := backupCmd.MarkFlagRequired("dbms"); err != nil {
		return
	}

	// TODO: add an ability to specify multiple storages
	backupCmd.Flags().StringVar(&storage, "storage", "", "Storage type to keep your dumps")
	if err := backupCmd.MarkFlagRequired("storage"); err != nil {
		return
	}

	// TODO: add validation to dbms and storage
	rootCmd.AddCommand(backupCmd)
}
