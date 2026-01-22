/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {

		if dbms == cfg.Postgres.String() {
			backup.PostgresDump(
				cfg.AppConfig.Postgres.Host,
				cfg.AppConfig.Postgres.User,
				cfg.AppConfig.Postgres.Password,
				cfg.AppConfig.Postgres.DB,
				storage,
				cfg.AppConfig.Postgres.Port,
			)
		}

	},
}

func init() {
	backupCmd.Flags().StringVar(&dbms, "dbms", "", "DBMS like postgres, mysql, etc")
	if err := backupCmd.MarkFlagRequired("dbms"); err != nil {
		return
	}

	// TODO: add an ability to specify multiple storages
	backupCmd.Flags().StringVar(&storage, "storage", "", "Storage type to keep your dump")
	if err := backupCmd.MarkFlagRequired("storage"); err != nil {
		return
	}
	rootCmd.AddCommand(backupCmd)

}
