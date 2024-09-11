package file

type FileService struct {
	repository FileRepository
}

func NewFileService(repository FileRepository) *FileService {
	return &FileService{repository}
}
