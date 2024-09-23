package services

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// curl -X PUT --upload-file ./your-file.txt "PRESIGNED_URL"
func GeneratePresignedPutUrl(fileId string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	bucketName := "quickshare-s3"
	objectKey := fileId

	presignedRequest, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}, s3.WithPresignExpires(5*time.Minute)) // 5 minutes expiry

	if err != nil {
		log.Fatalf("Failed to generate presigned URL: %v", err)
	}

	return presignedRequest.URL, err
}

// curl -X POST --upload-file ./your-file.txt "PRESIGNED_URL"
func GeneratePresignedGetUrl(fileId string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	bucketName := "quickshare-s3"
	objectKey := fileId

	presignedRequest, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}, s3.WithPresignExpires(5*time.Minute)) // 5 minutes expiry

	if err != nil {
		log.Fatalf("Failed to generate presigned URL: %v", err)
	}

	return presignedRequest.URL, err
}
