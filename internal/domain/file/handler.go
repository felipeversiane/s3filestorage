package file

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/s3filestorage/internal/infra/config/rest"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type fileHandler struct {
	service FileServiceInterface
}

type FileHandlerInterface interface {
	InsertHandler(c *gin.Context)
	GetOneHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
	ListAllHandler(c *gin.Context)
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
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		errorMessage := rest.NewBadRequestError("the ID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	file, err := h.service.GetOneService(ctxTimeout, id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, file)
}

func (h *fileHandler) DeleteHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		errorMessage := rest.NewBadRequestError("the ID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := h.service.DeleteService(ctxTimeout, id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *fileHandler) ListAllHandler(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	files, err := h.service.ListAllService(ctxTimeout)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, files)
}
