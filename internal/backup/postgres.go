package backup

import (
	"fmt"
	"os"
	"os/exec"

	st "github.com/KostyaBagr/duple-duple/internal/storage"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)

var IslocalStorage bool

// generate a name for dump. It takes datetime + zipped extension
func dumpPostgresFileName(archive bool) string {
	dateTime := utils.CurrentDateTimeRFC3339()
	ext := ".gz"
	if archive {
		ext = ".tar.gz"
	}
	return dateTime + ext
}

// Creates a postgres dump
// host - localhost or IP address
// user - db user
// table - table name (in this case we use pg_dump) OR * (in this case we use pg_dump_all)
// port - 5432 is a default value
// storageTypes - slice of selected storages
func PostgresDump(host, user, password, db, port string, storageTypes []string) {
	var isCluster bool
	var cmd *exec.Cmd

	if db == "*" {
		isCluster = true
	} else {
		isCluster = false
	}

	fileName := dumpPostgresFileName(isCluster)

	fullPath := dumpFullPath(fileName)

	if isCluster == false {
		cmd = exec.Command(
			"bash", "-c",
			fmt.Sprintf(
				"pg_dump -h %s -p %s -U %s -d %s | gzip > %s",
				host, port, user, db, fullPath,
			),
		)
	} else {
		cmd = exec.Command(
			"bash", "-c",
			fmt.Sprintf(
				"pg_dumpall -h %s -p %s -U %s | gzip > %s",
				host, port, user, fullPath,
			),
		)
	}

	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf(
			"pg_dump failed: %v\n",
			err,
		)
		return
	}

	err = st.StorageDispatcher(fullPath, storageTypes)
	if err != nil {
		fmt.Printf("Error in StorageDispatcher %v", err)
		return
	}

}
