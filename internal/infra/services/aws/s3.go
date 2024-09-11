package aws

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3Client S3Service

type S3Service interface {
	CreateBucket(ctx context.Context) error
	UploadFile(ctx context.Context, bucketKey string, fileContent multipart.File) (string, error)
	DeleteFile(ctx context.Context, key string) error
}

type s3Service struct {
	Client *s3.S3
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
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretAccessKey, ""),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("unable to create session, %v", err)
	}

	client := s3.New(sess)
	S3Client = &s3Service{
		Client: client,
		Bucket: bucket,
		Region: region,
		ACL:    acl,
	}

	return nil
}

func (s *s3Service) CreateBucket(ctx context.Context) error {
	_, err := s.Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s.Bucket),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(s.Region),
		},
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeBucketAlreadyOwnedByYou {
			return nil
		}
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeBucketAlreadyExists {
			return nil
		}
		return fmt.Errorf("unable to create bucket %s: %v", s.Bucket, err)
	}

	return nil
}

func (s *s3Service) UploadFile(ctx context.Context, bucketKey string, file multipart.File) (string, error) {
	defer file.Close()

	_, err := s.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(bucketKey),
		Body:   file,
		ACL:    aws.String(s.ACL),
	})
	if err != nil {
		return "", fmt.Errorf("unable to upload file %s to bucket %s: %v", bucketKey, s.Bucket, err)
	}

	url := fmt.Sprintf("%s/%s", s.Bucket, bucketKey)
	return url, nil
}

func (s *s3Service) DeleteFile(ctx context.Context, key string) error {
	_, err := s.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error deleting file %s from bucket %s: %v", key, s.Bucket, err)
	}

	return nil
}
