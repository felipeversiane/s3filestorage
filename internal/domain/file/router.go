package file

import (
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/database"
)

func FileRouter(mux *http.ServeMux) {
	handler := NewFileHandler(NewFileService(NewFileRepository(database.Connection, aws.S3Client)))
	mux.HandleFunc("POST /api/v1/file", handler.InsertHandler)
	mux.HandleFunc("GET /api/v1/file/{id}", handler.GetOneHandler)
	mux.HandleFunc("DELETE /api/v1/file", handler.DeleteHandler)

}
