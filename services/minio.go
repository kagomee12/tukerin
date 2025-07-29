package services

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinioClient() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		fmt.Println("Error initializing Minio client:", err)
		return
	}

	MinioClient = client

	bucketName := os.Getenv("MINIO_BUCKET")
	location := os.Getenv("MINIO_REGION")

	exist, err := client.BucketExists(context.Background(), bucketName)

	if err != nil {
		fmt.Println("Error checking if bucket exists:", err)
		return
	}

	if !exist {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			fmt.Println("Error creating bucket:", err)
			return
		}
	}

	policy := `{
		"Version": "2012-10-17",
		"Statement": [
		{
			"Action": ["s3:GetObject"],
			"Effect": "Allow",
			"Principal": "*",
			"Resource": ["arn:aws:s3:::` + bucketName + `/*"]
		}
	]
	}`

	err = client.SetBucketPolicy(context.Background(), bucketName, policy)

	if err != nil {
		fmt.Println("Error setting bucket policy:", err)
		return
	}

	fmt.Println("Minio client initialized successfully")
	
}