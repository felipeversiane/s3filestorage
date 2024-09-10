package aws

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var S3Client S3Service

type S3Service interface {
	CreateBucket(ctx context.Context) error
	UploadFile(ctx context.Context, bucketKey string, fileContent multipart.File) (string, error)
	DeleteFile(ctx context.Context, key string) error
}

type s3Service struct {
	Client *s3.Client
	Bucket string
	Region string
	ACL    string
}

func NewS3Service(
	bucket,
	region,
	acl,
	accessKey,
	secretAccessKey,
	endpoint string,
) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretAccessKey, "")),
	)
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %v", err)
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	S3Client = &s3Service{
		Client: client,
		Bucket: bucket,
		Region: region,
		ACL:    acl,
	}

	return nil
}

func (s *s3Service) CreateBucket(ctx context.Context) error {
	_, err := s.Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(s.Bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(s.Region),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create bucket %s: %v", s.Bucket, err)
	}

	fmt.Printf("Bucket %s successfully created\n", s.Bucket)
	return nil
}

func (s *s3Service) UploadFile(ctx context.Context, bucketKey string, file multipart.File) (string, error) {
	defer file.Close()

	_, err := s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(bucketKey),
		Body:   file,
		ACL:    types.ObjectCannedACL(s.ACL),
	})
	if err != nil {
		return "", fmt.Errorf("unable to upload file %s to bucket %s: %v", bucketKey, s.Bucket, err)
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, bucketKey)
	return url, nil
}

func (s *s3Service) DeleteFile(ctx context.Context, key string) error {
	_, err := s.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error deleting file %s from bucket %s: %v", key, s.Bucket, err)
	}

	fmt.Printf("File %s successfully deleted from bucket %s\n", key, s.Bucket)
	return nil
}
