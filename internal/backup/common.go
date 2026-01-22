package backup

import (
	"fmt"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)

// Generates full path. It local type of storage was selected all dump files will be uploaded
// there and will not be deleted. Instead in tmpDirPostres

func dumpFullPath(fileName string) (path string, err error) {
	localStoragePath := cfg.AppConfig.Storage.Local.Path

	isLocalCfgEmpty, err := utils.IsEmpty(localStoragePath)
	if err != nil {
		return "", fmt.Errorf("reading storage config: %w", err)
	}

	if !isLocalCfgEmpty {
		return localStoragePath + "/" + fileName, nil
	}

	return cfg.TmpBackupPath + "/" + fileName, nil
}
