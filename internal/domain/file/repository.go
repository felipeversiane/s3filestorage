package file

import (
	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepository struct {
	db *pgxpool.Pool
	s3 aws.S3Service
}

func NewFileRepository(db *pgxpool.Pool, s3 aws.S3Service) *FileRepository {
	return &FileRepository{db, s3}
}
