package lorem_grpc

import (
	"context"
	"errors"
	guuid "github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"log"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// Define service interface
type Service interface {
	//
	Create(ctx context.Context, requestType string, min, max int) (string, error)
}

// Implement service with empty struct
type Storage struct {
}

// Implement service functions
func (Storage) Create(_ context.Context, requestType string, min, max int) (string, error) {
	endpoint := "localhost:9001"
	accessKeyID := "X+qbbAL*mJ:XuOY"
	secretAccessKey := "7puHEDd5EV7OMH[s02J:5I3iKV*!K"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	id := guuid.New()
	objectName := id.String()
	filePath := "main.go"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	return objectName, err
}
