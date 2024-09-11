package file

import (
	"context"
	"mime/multipart"

	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
	"github.com/google/uuid"
)

type FileServiceInterface interface {
	InsertService(ctx context.Context, fileHeader *multipart.FileHeader) (*FileResponse, *rest.RestError)
	GetOneService(ctx context.Context, id uuid.UUID) (*FileResponse, *rest.RestError)
	DeleteService(ctx context.Context, id uuid.UUID) *rest.RestError
	ListAllService(ctx context.Context) ([]*FileResponse, *rest.RestError)
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

func (s *fileService) GetOneService(ctx context.Context, id uuid.UUID) (*FileResponse, *rest.RestError) {
	domain, err := s.repository.GetOneRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	return DomainToFileResponse(domain), nil
}

func (s *fileService) DeleteService(ctx context.Context, id uuid.UUID) *rest.RestError {
	err := s.repository.DeleteRepository(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *fileService) ListAllService(ctx context.Context) ([]*FileResponse, *rest.RestError) {
	files, err := s.repository.ListAllRepository(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*FileResponse
	for _, file := range files {
		responses = append(responses, DomainToFileResponse(file))
	}

	return responses, nil
}
