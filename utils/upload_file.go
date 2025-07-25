package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"tukerin/services"

	"github.com/lithammer/shortuuid/v4"
	"github.com/minio/minio-go/v7"
)

func UploadFile(fileHeader *multipart.FileHeader, bucketName string) (string, error) {
	file, err := fileHeader.Open()

	if err != nil {
		return "", err
	}
	defer file.Close()

	uniqueName := shortuuid.New() + "-" + fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")
	fileSize := fileHeader.Size

	_, err = services.MinioClient.PutObject(
		context.Background(),
		bucketName,
		uniqueName,
		file,
		fileSize,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s",os.Getenv("MINIO_ENDPOINT"), bucketName, uniqueName)

	return url, nil
}

func UploadMultipleFiles(files []*multipart.FileHeader, bucketName string) ([]string, error) {
	var urls []string

	for _, fileHeader := range files {
		url, err := UploadFile(fileHeader, bucketName)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}