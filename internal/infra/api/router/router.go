package router

import (
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/domain/file"
)

func SetupRoutes(mux *http.ServeMux) {
	file.FileRouter(mux)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
