// Module for the dispathcer that calls different functions for each type of storage
package storage

import (
	"context"
	"errors"
	"log"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	s3 "github.com/KostyaBagr/duple-duple/internal/storage/s3"

)


// Initilizes context for S3, processes name and calls s3 client
// for uploading large files in stream
func s3LoaderDispatcher(fileName string, fileBytes []byte) error {
	s, err := s3.NewS3Storage()
	if err != nil {
		log.Printf("Error during init s3 %v", err)
		return errors.New("Unable to intialize s3 function")
	}

	ctx := context.Background()
	if cfg.AppConfig.S3.PathInBucket != "" {
		fileName = cfg.AppConfig.S3.PathInBucket + fileName
	}

	if err = s.UploadLargeObject(ctx, cfg.AppConfig.S3.BacketName, fileName, fileBytes); err != nil {
		log.Print(err)
		return errors.New("Error during uploading file to s3")
	} else {
		log.Println("File was successfuly upload to S3")
	}
	return nil
}

// based on storageType calls function for storage processing
func StorageDispatcher(storageType, fileName string, fileBytes []byte) error {
	switch storageType {
	case "s3":
		s3LoaderDispatcher(fileName, fileBytes)
	default:
		return errors.New("Invalid storageType was providen")
	}
	
	return nil
}
