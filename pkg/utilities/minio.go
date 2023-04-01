package utilities

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection() (*minio.Client, error) {
	endpoint := "localhost:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false

	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	return minioClient, errInit
}

func UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Get Buffer from file
	buffer, err := fileHeader.Open()

	if err != nil {
		return "", err
	}
	defer buffer.Close()

	minioClient, err := MinioConnection()
	if err != nil {
		return "", err
	}

	objectName := fileHeader.Filename
	fileBuffer := buffer
	contentType := fileHeader.Header["Content-Type"][0]
	fileSize := fileHeader.Size

	info, err := minioClient.PutObject(ctx, "asset", objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	log.Println("success upload file", info)
	return `http://localhost:9000/asset/` + objectName, nil
}
