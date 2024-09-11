package file

import (
	"context"
	"mime/multipart"

	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
)

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
	if err := validateFile(fileHeader); err != nil {
		return nil, err
	}
	domain, err := s.repository.InsertRepository(ctx, fileHeader)
	if err != nil {
		return nil, err
	}
	return DomainToFileResponse(domain), nil
}

func (s *fileService) GetOneService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError) {
	return nil, nil
}

func (s *fileService) DeleteService(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError {
	return nil
}
