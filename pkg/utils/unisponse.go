package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unisponse(c *gin.Context, name string, value int64) {
	c.JSON(http.StatusOK, gin.H{
		"value": value,
		"name":  name,
	})
}
