package backup

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/KostyaBagr/duple-duple/internal"
	app "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/storage/s3"
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

	fileName := dumpFileName(isCluster)
	fullPath := outputDir + fileName
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
	// TODO: replace this S3 init function to common module
	s, err := s3.NewS3Storage()
	if err != nil {
		log.Printf("Error during init s3 %v", err)
		return
	}
	ctx := context.Background()
	dumpBytes, err := internal.ConvertFileToBytes(fullPath)
	log.Println("File was converted to bytes")
	if err != nil {
		log.Print(err)
		return
	}
	if app.AppConfig.S3.PathInBucket != "" {
		fileName = app.AppConfig.S3.PathInBucket + fileName
	}
	if err = s.UploadLargeObject(ctx, app.AppConfig.S3.BacketName, fileName, dumpBytes); err != nil {
		log.Print(err)
		return
	} else {
		log.Println("File was successfuly upload to S3")
	}

}
