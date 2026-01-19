/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/KostyaBagr/duple-duple/internal/backup"
	cfg "github.com/KostyaBagr/duple-duple/internal/config" 
)

var dbms string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your db",
	Long:  "TODO: change it!",
	Run: func(cmd *cobra.Command, args []string) {
		if dbms == "postgres" {
			backup.PostgresDump(
				cfg.AppConfig.Postgres.Host, 
				cfg.AppConfig.Postgres.User, 
				cfg.AppConfig.Postgres.Password, 
				cfg.AppConfig.Postgres.DB, 
				cfg.AppConfig.Postgres.OutputDir, 
				cfg.AppConfig.Postgres.Port,
			)
		}
	},
}

func init() {
	backupCmd.Flags().StringVar(&dbms, "dbms", "", "DBMS like postgres, mysql, etc")
	if err := backupCmd.MarkFlagRequired("dbms"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(backupCmd)

}
