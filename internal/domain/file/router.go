package file

import (
	"github.com/felipeversiane/s3filestorage/internal/infra/services/aws"
	"github.com/felipeversiane/s3filestorage/internal/infra/services/database"
	"github.com/gin-gonic/gin"
)

func FileRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	handler := NewFileHandler(NewFileService(NewFileRepository(database.Connection, aws.S3Client)))

	file := v1.Group("/file")
	{
		file.POST("/", handler.InsertHandler)
		file.GET("/:id", handler.GetOneHandler)
		file.DELETE("/:id", handler.DeleteHandler)

	}

	return file
}
