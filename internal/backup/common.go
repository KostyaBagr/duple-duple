package backup

import (
	"errors"
	"fmt"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)

// Generates full path. It local type of storage was selected all dump files will be uploaded
// there and will not be deleted. Instead in tmpDirPostres
func dumpFullPath(fileName string) (path string) {
	localStoragePath := cfg.AppConfig.Storage.Local.Path

	isLocalCfgEmpty := utils.IsEmpty(localStoragePath)
	if !isLocalCfgEmpty {
		_, err := utils.PathExists(localStoragePath, true)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return localStoragePath + fileName

	}

	return cfg.TmpBackupPath + fileName
}

// This is a discpatcher for dumping.
// Based on dbms var it runs a piplelin
// Returns DumpFileStats (such as time complexity,
// file size and so on),
// path to file (for storagies) and error (optional)
func DumpDispatcher(dbms string) (*DumpFileStats, string, error) {

	if dbms == cfg.Postgres.String() {
		stat, path, err := PostgresDump(
			cfg.AppConfig.Postgres.Host,
			cfg.AppConfig.Postgres.User,
			cfg.AppConfig.Postgres.Password,
			cfg.AppConfig.Postgres.DB,
			cfg.AppConfig.Postgres.Port,
		)
		if err != nil {
			return stat, "", errors.New("Unable to create postgres dump")
		}
		return stat, path, nil
	}
	return &DumpFileStats{}, "", errors.New("Undefined dbms")
}
