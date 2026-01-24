// Module for the dispathcer that calls different functions for each type of storage
package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	s3 "github.com/KostyaBagr/duple-duple/internal/storage/s3"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)

// Initilizes context for S3, processes name and calls s3 client
// for uploading large files in stream
func s3LoaderDispatcher(filePath string) error {
	fmt.Println("here")
	dumpBytes, err := utils.ConvertFileToBytes(filePath)
	if err != nil {
		log.Printf("Can't dump file to bytes %v", err)
		return errors.New("can't dump file to bytes")
	}

	s, err := s3.NewS3Storage()
	if err != nil {
		log.Printf("Error during init s3 %v", err)
		return errors.New("unable to intialize s3 function")
	}

	fileName, err := utils.GetFileNameFromPath(filePath)
	if err != nil {
		fmt.Printf("Can't parse file name %v", err)
		return errors.New("can't parse file name")
	}

	ctx := context.Background()
	if cfg.AppConfig.Storage.S3.PathInBucket != "" {
		fileName = cfg.AppConfig.Storage.S3.PathInBucket + fileName
	}

	if err = s.UploadLargeObject(
		ctx,
		cfg.AppConfig.Storage.S3.BacketName,
		fileName,
		dumpBytes,
	); err != nil {
		fmt.Printf("Error during uploading file to s3 %v", err)
		return errors.New("Error during uploading file to s3")
	} else {
		fmt.Printf("File %s was successfuly upload to S3\n", filePath)
		return nil
	}
}

// based on storageType calls function for storage processing
func StorageDispatcher(filePath string, storageTypes []string) error {
	// When new methods will be added, call functions via gorutines

	for _, storageType := range storageTypes {

		if storageType == cfg.S3.String() {
			s3LoaderDispatcher(filePath)
		}
	}

	// if a local type of storage was not provided by user delete it beacause it was a temporary file
	if !slices.Contains(storageTypes, cfg.Local.String()) {
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			return errors.New("Error during deliting tmp file")
		}
		fmt.Printf("%s was deleted as temporary file\n", filePath)
	}
	return nil
}
