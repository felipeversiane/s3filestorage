package file

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/felipeversiane/s3filestorage/internal/domain"
	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/google/uuid"
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
	fileID := uuid.New()
	fileExtension := filepath.Ext(fileHeader.Filename)
	fileKey := fmt.Sprintf("%s%s", fileID.String(), fileExtension)
	file, err := fileHeader.Open()
	if err != nil {
		return nil, rest.NewBadRequestError(fmt.Sprintf("unable to open file: %v", err))
	}
	defer file.Close()

	url, err := r.s3.UploadFile(ctx, fileKey, file)
	if err != nil {
		return nil, rest.NewInternalServerError(fmt.Sprintf("unable to upload file to S3: %v", err))
	}

	domainFile := domain.NewFile(url, fileKey)

	query := `INSERT INTO files (id, key, url) VALUES ($1, $2, $3)`
	_, err = r.db.Exec(ctx, query, domainFile.ID, domainFile.Key, domainFile.URL)
	if err != nil {
		return nil, rest.NewInternalServerError(fmt.Sprintf("unable to insert file into database: %v", err))
	}

	return domainFile, nil
}

func (r *fileRepository) GetOneRepository(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.File, *rest.RestError) {
	return nil, nil
}

func (r *fileRepository) DeleteRepository(ctx context.Context, fileHeader *multipart.FileHeader) *rest.RestError {
	return nil
}
