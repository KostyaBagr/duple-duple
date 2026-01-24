package backup

import (
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
