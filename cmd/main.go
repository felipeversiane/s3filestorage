package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/log"
	"github.com/felipeversiane/s3filestorage/internal/router"

	"github.com/felipeversiane/s3filestorage/internal/config"
	"github.com/felipeversiane/s3filestorage/internal/services/aws"
)

func main() {
	log.Configure()
	conf := config.NewConfig()
	slog.Info("Config init...")

	slog.Info("Creating a new AWS-S3 service...")
	err := aws.NewS3Service(conf.S3.Bucket, conf.S3.Region, conf.S3.ACL, conf.AWS.ACCESS_KEY, conf.AWS.SECRET_ACCESS_KEY, conf.S3.Endpoint)
	if err != nil {
		slog.Error(fmt.Sprintf("init s3 service error: %s", err))
	}
	slog.Info("New AWS-S3 service created...")

	slog.Info("Creating a new AWS-S3 bucket...")
	aws.S3Client.CreateBucket(context.Background())
	slog.Info("New AWS-S3 bucket created...")

	slog.Info("Creating http server...")
	mux := http.NewServeMux()
	router.SetupRoutes(mux)
	handler := log.LogMiddleware(mux)

	slog.Info(fmt.Sprintf("Server running on port :%s", conf.Api.Port))
	http.ListenAndServe(":"+conf.Api.Port, handler)
}
