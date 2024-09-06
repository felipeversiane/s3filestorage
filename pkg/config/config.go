package config

import (
	"fmt"
	"os"
)

type config struct {
	Api struct {
		Port string
	}
	S3 struct {
		Region   string
		Endpoint string
		Bucket   string
		ACL      string
	}
}

func NewConfig() *config {
	return &config{
		S3: struct {
			Region   string
			Endpoint string
			Bucket   string
			ACL      string
		}{
			Region:   getEnvOrDie("AWS_REGION"),
			Endpoint: getEnvOrDie("AWS_ENDPOINT"),
			Bucket:   getEnvOrDie("AWS_S3_BUCKET"),
			ACL:      getEnvOrDie("AWS_S3_ACL"),
		},
		Api: struct {
			Port string
		}{
			Port: getEnvOrDie("PORT"),
		},
	}
}

func getEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		panic(err)
	}
	return value
}
