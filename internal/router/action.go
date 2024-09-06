package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/services/aws"
)

func uploadHandler(s3Service aws.S3Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		bucketKey := handler.Filename

		url, err := s3Service.UploadFile(context.Background(), bucketKey, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to upload file: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully: %s", url)
	}
}
