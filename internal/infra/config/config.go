package config

import (
	"fmt"
	"os"
)

var Conf *config

type config struct {
	Api struct {
		Port string
	}
	AWS struct {
		ACCESS_KEY        string
		SECRET_ACCESS_KEY string
	}
	S3 struct {
		Region   string
		Endpoint string
		Bucket   string
		ACL      string
	}
	Database struct {
		Host     string
		Name     string
		Port     string
		User     string
		Password string
	}
}

func NewConfig() {
	Conf = &config{
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
			Port: getEnvOrDie("API_PORT"),
		},
		AWS: struct {
			ACCESS_KEY        string
			SECRET_ACCESS_KEY string
		}{
			ACCESS_KEY:        getEnvOrDie("AWS_ACCESS_KEY"),
			SECRET_ACCESS_KEY: getEnvOrDie("AWS_SECRET_ACCESS_KEY"),
		},
		Database: struct {
			Host     string
			Name     string
			Port     string
			User     string
			Password string
		}{
			Host:     getEnvOrDie("POSTGRES_HOST"),
			Name:     getEnvOrDie("POSTGRES_DB"),
			Port:     getEnvOrDie("POSTGRES_PORT"),
			User:     getEnvOrDie("POSTGRES_USER"),
			Password: getEnvOrDie("POSTGRES_PASSWORD"),
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
