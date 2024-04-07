package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func MinioMiddleware(minioClient *minio.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Minio", minioClient)
		c.Next()
	}
}
