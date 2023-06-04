package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(message string) gin.H {
	return gin.H{"error": message}
}