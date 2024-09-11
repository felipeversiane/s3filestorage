package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/felipeversiane/s3filestorage/internal/infra/api/router"
	"github.com/felipeversiane/s3filestorage/internal/infra/config/log"
	"github.com/gin-gonic/gin"

	"github.com/felipeversiane/s3filestorage/internal/infra/config"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/database"
)

func main() {
	log.Configure()
	config.NewConfig()
	slog.Info("Config init...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := database.Connect(ctx); err != nil {
		panic(err)
	}
	defer database.Close()

	slog.Info("Creating a new AWS-S3 service...")
	err := aws.NewS3Service(
		config.Conf.S3.Bucket,
		config.Conf.S3.Region,
		config.Conf.S3.ACL,
		config.Conf.AWS.ACCESS_KEY,
		config.Conf.AWS.SECRET_ACCESS_KEY,
		config.Conf.S3.Endpoint)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("New AWS-S3 service created...")

	slog.Info("Creating a new AWS-S3 bucket...")
	err = aws.S3Client.CreateBucket(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("New AWS-S3 bucket created...")

	slog.Info("Creating http server...")
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(log.LogMiddleware())
	router.SetupRoutes(g)
	slog.Info(fmt.Sprintf("Server running on port :%s", config.Conf.Api.Port))
	g.Run(":" + config.Conf.Api.Port)

}
