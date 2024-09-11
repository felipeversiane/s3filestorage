package file

import (
	"mime/multipart"

	"github.com/felipeversiane/s3filestorage/internal/domain"
	"github.com/google/uuid"
)

type FileRequest struct {
	File *multipart.FileHeader `json:"file" binding:"required,filesize"`
}

type FileResponse struct {
	ID  uuid.UUID `json:"id"`
	URL string    `json:"url"`
}

func DomainToFileResponse(file domain.File) FileResponse {
	return FileResponse{
		ID:  file.ID,
		URL: file.URL,
	}
}
