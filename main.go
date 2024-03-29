package main

import (
	guuid "github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
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

	objectName, n := Create(minioClient, bucketName)

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func Create(minioClient *minio.Client, bucketName string) (string, int64) {
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
	return objectName, n
}

func Get(minioClient *minio.Client, bucketName string, fileName string) error {
	err := minioClient.FGetObject(bucketName, fileName, "retrievedObject.go", minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
