/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/KostyaBagr/duple-duple/internal/backup"
	"github.com/spf13/cobra"
)

type BackupFunc func(
	host string,
	user string,
	password string,
	table string,
	output string,
	port string,
)

var (
	dms, host, user, password, db, outputDir, port string
)

var backupFunc = map[string]BackupFunc{
	"postgres": backup.PostgresDump,
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your db",
	Long:  "TODO: change it!",
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			fmt.Printf("Path %s does not exist", outputDir)
			log.Fatalf("Path %s does not exist", outputDir)
		}

		fn, ok := backupFunc[dms]
		if !ok {
			fmt.Print("Invalid DMS")
			log.Fatalf("Passed a wrong value for param dms %v", dms)

		}
		fn(host, user, password, db, outputDir, port)
	},
}

func init() {
	// TODO: Add dbs like mysql mongo for next versions
	backupCmd.Flags().StringVar(&dms, "dms", "postgres", "DMS like postgres, mysql, etc")
	backupCmd.Flags().StringVar(&host, "host", "localhost", "DB host")
	backupCmd.Flags().StringVar(&user, "user", "", "DB user")
	backupCmd.Flags().StringVar(&password, "password", "", "user password")
	backupCmd.Flags().StringVar(&db, "db", "*", "table name to dump or * to dump a cluster")
	backupCmd.Flags().StringVar(&port, "port", "5432", "DB port")
	backupCmd.Flags().StringVar(&outputDir, "outputDir", "", "Directory to save dump")
	if err := backupCmd.MarkFlagRequired("user"); err != nil {
		panic(err)
	}
	if err := backupCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
	if err := backupCmd.MarkFlagRequired("outputDir"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(backupCmd)

}
