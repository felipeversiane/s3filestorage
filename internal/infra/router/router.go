package router

import (
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/file", uploadHandler(aws.S3Client))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
