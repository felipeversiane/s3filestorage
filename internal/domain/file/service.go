package file

import (
	"context"
	"mime/multipart"

	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
)

const maxFileSize = 10 << 20

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

type FileServiceInterface interface {
	InsertService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError)
	GetOneService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError)
	DeleteService(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError
}

type fileService struct {
	repository FileRepositoryInterface
}

func NewFileService(repository FileRepositoryInterface) FileServiceInterface {
	return &fileService{repository}
}

func (s *fileService) InsertService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError) {
	return nil, nil
}

func (s *fileService) GetOneService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError) {
	return nil, nil
}

func (s *fileService) DeleteService(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError {
	return nil
}
