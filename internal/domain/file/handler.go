package file

import (
	"github.com/gin-gonic/gin"
)

type fileHandler struct {
	service FileServiceInterface
}

type FileHandlerInterface interface {
	InsertHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
}

func NewFileHandler(service FileServiceInterface) FileHandlerInterface {
	return &fileHandler{service}
}

func (h *fileHandler) InsertHandler(c *gin.Context) {
}

func (h *fileHandler) GetOneHandler(c *gin.Context) {
}

func (h *fileHandler) DeleteHandler(c *gin.Context) {
}
