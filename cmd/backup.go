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
	"github.com/KostyaBagr/duple-duple/internal/notifications"
	st "github.com/KostyaBagr/duple-duple/internal/storage"
	"github.com/KostyaBagr/duple-duple/internal/utils"
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

		storageTypes := strings.Split(storage, ",")

		if !utils.SliceIsSubSlice(possibleStorages, storageTypes) {
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

		dumpStat, path, err := backup.DumpDispatcher(dbms)

		if err != nil {
			fmt.Printf("Error in DBMS dispatcher %v")
			return err
		}

		dumpStat.Storages = storageTypes
		err = st.StorageDispatcher(path, storageTypes)

		if err != nil {
			fmt.Printf("Error in StorageDispatcher %v", err)
			return err
		}
		// send in another thread for future
		notifications.NotificationDumpDispatcher(
			cfg.AppConfig.Notifications.Email.Receiver,
			*dumpStat,
		)

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
