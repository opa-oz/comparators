package middlewares

import (
	"github.com/gin-gonic/gin"
	"comparators/pkg/config"
)

func CfgMiddleware(cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Config", cfg)
		c.Next()
	}
}
