package router

import (
	"net/http"

	"github.com/felipeversiane/s3filestorage/internal/domain/file"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		file.FileRouter(v1)
	}

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
