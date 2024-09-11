package file

import (
	"context"
	"mime/multipart"

	"github.com/felipeversiane/s3filestorage/internal/domain"
	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepositoryInterface interface {
	InsertRepository(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.File, *rest.RestError)
	GetOneRepository(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.File, *rest.RestError)
	DeleteRepository(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError
}

type fileRepository struct {
	db *pgxpool.Pool
	s3 aws.S3Service
}

func NewFileRepository(db *pgxpool.Pool, s3 aws.S3Service) FileRepositoryInterface {
	return &fileRepository{db, s3}
}

func (r *fileRepository) InsertRepository(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.File, *rest.RestError) {
	return nil, nil
}

func (r *fileRepository) GetOneRepository(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.File, *rest.RestError) {
	return nil, nil
}

func (r *fileRepository) DeleteRepository(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError {
	return nil
}
