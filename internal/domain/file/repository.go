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
	GetOneRepository(ctx context.Context, id uuid.UUID) (*domain.File, *rest.RestError)
	DeleteRepository(ctx context.Context, id uuid.UUID) *rest.RestError
	ListAllRepository(ctx context.Context) ([]*domain.File, *rest.RestError)
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

func (r *fileRepository) GetOneRepository(ctx context.Context, id uuid.UUID) (*domain.File, *rest.RestError) {
	query := `SELECT key, url FROM files WHERE id = $1`

	var file domain.File
	err := r.db.QueryRow(ctx, query, id).Scan(&file.Key, &file.URL)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, rest.NewNotFoundError(fmt.Sprintf("file with id %s not found", id))
		}
		return nil, rest.NewInternalServerError(fmt.Sprintf("unable to retrieve file: %v", err))
	}

	file.ID = id
	return &file, nil
}

func (r *fileRepository) DeleteRepository(ctx context.Context, id uuid.UUID) *rest.RestError {
	query := `SELECT key FROM files WHERE id = $1`

	var fileKey string
	err := r.db.QueryRow(ctx, query, id).Scan(&fileKey)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return rest.NewNotFoundError(fmt.Sprintf("file with id %s not found", id))
		}
		return rest.NewInternalServerError(fmt.Sprintf("unable to retrieve file key: %v", err))
	}

	err = r.s3.DeleteFile(ctx, fileKey)
	if err != nil {
		return rest.NewInternalServerError(fmt.Sprintf("unable to delete file from S3: %v", err))
	}

	deleteQuery := `DELETE FROM files WHERE id = $1`
	_, err = r.db.Exec(ctx, deleteQuery, id)
	if err != nil {
		return rest.NewInternalServerError(fmt.Sprintf("unable to delete file from database: %v", err))
	}

	return nil
}

func (r *fileRepository) ListAllRepository(ctx context.Context) ([]*domain.File, *rest.RestError) {
	query := `SELECT id, key, url FROM files`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, rest.NewInternalServerError(fmt.Sprintf("unable to retrieve files from database: %v", err))
	}
	defer rows.Close()

	var files []*domain.File
	for rows.Next() {
		var file domain.File
		if err := rows.Scan(&file.ID, &file.Key, &file.URL); err != nil {
			return nil, rest.NewInternalServerError(fmt.Sprintf("unable to scan file: %v", err))
		}
		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, rest.NewInternalServerError(fmt.Sprintf("row iteration error: %v", err))
	}

	return files, nil
}
