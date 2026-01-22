package s3

import (
	"bytes"
	"context"
	"errors"
	"log"
	"time"

	app "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/smithy-go"
)

type S3Storage struct {
	S3Client *s3.Client
}

// Creates an instance of s3
func NewS3Storage() (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(app.AppConfig.Storage.S3.Region),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &app.AppConfig.Storage.S3.Url
		o.UsePathStyle = true
		o.Credentials = aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				app.AppConfig.Storage.S3.AccessKey,
				app.AppConfig.Storage.S3.SecretAccessKey,
				"",
			),
		)
	})

	s3 := S3Storage{S3Client: client}
	return &s3, nil
}

// UploadLargeObject uses an upload manager to upload data to an object in a bucket.
// The upload manager breaks large data into parts and uploads the parts concurrently.
func (s *S3Storage) UploadLargeObject(
	ctx context.Context,
	bucketName string,
	objectKey string,
	largeObject []byte,
) error {
	largeBuffer := bytes.NewReader(largeObject)
	var partMiBs int64 = 10
	log.Println("In s3")
	uploader := manager.NewUploader(s.S3Client, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
	})
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   largeBuffer,
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"The maximum size for a multipart upload is 5TB.", bucketName)
		} else {
			log.Printf("Couldn't upload large object to %v:%v. Here's why: %v\n",
				bucketName, objectKey, err)
		}
	} else {
		err = s3.NewObjectExistsWaiter(s.S3Client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(bucketName), Key: aws.String(objectKey)}, time.Minute)
		if err != nil {
			log.Printf("Failed attempt to wait for object %s to exist.\n", objectKey)
		}
	}

	return err
}
