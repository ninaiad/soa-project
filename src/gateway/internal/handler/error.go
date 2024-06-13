package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println("error response:", message)
	c.AbortWithStatusJSON(statusCode, map[string]interface{}{
		"error": message,
	})
}
