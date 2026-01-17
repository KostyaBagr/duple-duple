package backup

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/KostyaBagr/duple-duple/internal"
)

// generate a name for dump. It takes outputDir and add current datetime + zipped extension
func dumpFileName(outputDir string, archive bool) string {
	dateTime := internal.CurrentDateTimeRFC3339()
	ext := ".gz"
	if archive {
		ext = ".tar.gz"
	}
	return outputDir + dateTime + ext
}

// Creates a postgres dump
// host - localhost or IP address
// user - db user
// table - table name (in this case we use pg_dump) OR * (in this case we use pg_dump_all)
// outputDir - dir to save dumps
// port - 5432 is a default value
func PostgresDump(host, user, password, db, outputDir, port string) {

	var isCluster bool
	var cmd *exec.Cmd

	if db == "*" {
		isCluster = true
	} else {
		isCluster = false
	}

	dumpFullPath := dumpFileName(outputDir, isCluster)


	if isCluster == false {
		cmd = exec.Command(
			"bash", "-c",
			fmt.Sprintf(
				"pg_dump -h %s -p %s -U %s -d %s | gzip > %s",
				host, port, user, db, dumpFullPath,
			),
		)
	} else {
		cmd = exec.Command(
			"bash", "-c",
			fmt.Sprintf(
				"pg_dumpall -h %s -p %s -U %s | gzip > %s",
				host, port, user, dumpFullPath,
			),
		)
	}
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	err := cmd.Run()

	if err != nil {
		log.Fatalf(
			"pg_dump failed: %v\n",
			err,
		)
	}

}
