package file

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/felipeversiane/s3filestorage/internal/domain"
	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
	"github.com/google/uuid"
)

const maxFileSize = 10 << 20

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

type FileResponse struct {
	ID  uuid.UUID `json:"id"`
	URL string    `json:"url"`
}

func DomainToFileResponse(file *domain.File) *FileResponse {
	return &FileResponse{
		ID:  file.ID,
		URL: file.URL,
	}
}

func validateFile(fileHeader *multipart.FileHeader) *rest.RestError {

	if fileHeader.Size > maxFileSize {
		return rest.NewBadRequestError("file size exceeds the 10 MB limit")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return rest.NewBadRequestError("file is not an allowed image type")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return rest.NewInternalServerError("failed to open file")
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return rest.NewInternalServerError("failed to read file")
	}

	contentType := http.DetectContentType(buffer)
	if !strings.HasPrefix(contentType, "image/") {
		return rest.NewBadRequestError("file is not an image")
	}

	return nil
}
