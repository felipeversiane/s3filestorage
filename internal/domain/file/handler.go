package file

import (
	"net/http"
)

type fileHandler struct {
	service FileServiceInterface
}

type FileHandlerInterface interface {
	InsertHandler(w http.ResponseWriter, r *http.Request)
	GetOneHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
}

func NewFileHandler(service FileServiceInterface) FileHandlerInterface {
	return &fileHandler{service}
}

func (h *fileHandler) InsertHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *fileHandler) GetOneHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *fileHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
