package config

import (
	"fmt"
	"os"
)

var Conf Config

type Config struct {
	S3 struct {
		Region   string
		Endpoint string
		Bucket   string
		ACL      string
	}
}

func NewConfig() {
	Conf = Config{
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
