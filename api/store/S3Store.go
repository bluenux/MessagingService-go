package store

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

type S3Store struct {
	fileStore  FileStore
	region     string
	bucket     string
	s3StoreKey string
}

func (s S3Store) Set(token string) bool {
	//prepare

	if s.fileStore.Set(token) != true {
		return false
	}

	file, err := os.Open(s.fileStore.filePath)
	if err != nil {
		log.Printf("Unable to open file %q, %v", s.fileStore.filePath, err)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(s.region)},
	)

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.s3StoreKey),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		log.Printf("Unable to upload %q to %q, %v", s.fileStore.filePath, s.s3StoreKey, err)
		return false
	}

	fmt.Printf("Successfully uploaded %q to %q\n", s.fileStore.filePath, s.s3StoreKey)

	return true

}

func (s S3Store) All() []string {
	//prepare

	file, err := os.Create(s.fileStore.filePath)
	if err != nil {
		log.Printf("Unable to open file %q, %v", s.fileStore.filePath, err)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(s.region)},
	)

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(s.s3StoreKey),
		})
	if err != nil {
		log.Printf("Unable to download item %q, %v", s.s3StoreKey, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return s.fileStore.All()

}
