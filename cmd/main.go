package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/router"

	"github.com/felipeversiane/s3filestorage/internal/services/aws"
	"github.com/felipeversiane/s3filestorage/pkg/config"
)

func main() {
	conf := config.NewConfig()

	err := aws.NewS3Service(conf.S3.Bucket, conf.S3.Region, conf.S3.ACL)
	if err != nil {
		slog.Error(fmt.Sprintf("init s3 service error: %s", err))
	}
	aws.S3Client.CreateBucket(context.Background())

	mux := http.NewServeMux()
	router.SetupRoutes(mux)

	http.ListenAndServe(":"+conf.Api.Port, mux)

}
