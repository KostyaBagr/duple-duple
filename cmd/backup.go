/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/KostyaBagr/duple-duple/internal/backup"
)

var dbms string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your db",
	Long:  "TODO: change it!",
	Run: func(cmd *cobra.Command, args []string) {
		if dbms == "postgres" {
			backup.PostgresDump(
				appConfig.Postgres.Host, 
				appConfig.Postgres.User, 
				appConfig.Postgres.Password, 
				appConfig.Postgres.DB, 
				appConfig.Postgres.OutputDir, 
				appConfig.Postgres.Port,
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
