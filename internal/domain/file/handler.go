package file

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
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
	fileHeader, formErr := c.FormFile("file")
	if formErr != nil {
		validationError := rest.NewBadRequestError("file is required")
		c.JSON(validationError.Code, validationError)
		return
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := h.service.InsertService(ctxTimeout, fileHeader)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *fileHandler) GetOneHandler(c *gin.Context) {
}

func (h *fileHandler) DeleteHandler(c *gin.Context) {
}
