package backup

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/KostyaBagr/duple-duple/internal"
	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	st "github.com/KostyaBagr/duple-duple/internal/storage"
)

// generate a name for dump. It takes datetime + zipped extension
func dumpFileName(archive bool) string {
	dateTime := internal.CurrentDateTimeRFC3339()
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
// storage - a type of storage to keep your backups
// port - 5432 is a default value
func PostgresDump(host, user, password, db, storage, port string) {
	var isCluster bool
	var cmd *exec.Cmd

	if db == "*" {
		isCluster = true
	} else {
		isCluster = false
	}

	fileName := dumpFileName(isCluster)

	fullPath := cfg.AppConfig.Postgres.OutputDir + fileName
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
		log.Printf(
			"pg_dump failed: %v\n",
			err,
		)
		return
	}
	
	log.Printf("Created dump %s", fileName)
	dumpBytes, err := internal.ConvertFileToBytes(fullPath)
	if err != nil {
		log.Println(err)
		return
	}

	err = st.StorageDispatcher("s3", fileName, dumpBytes)
	if err != nil {
		log.Println(err)
		fmt.Printf("Invalid type of storage %v", storage)
		return
	}

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("Postgres dump was successfully uploaded!")

}
