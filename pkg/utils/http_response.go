package utils

import "github.com/gin-gonic/gin"

type errorMessages struct {
	SignUpError, LoginError string
}

var ErrorMessages = errorMessages{
	SignUpError: "Email already exists",
	LoginError:  "Invalid email or password",
}

func ErrorResponse(message string) gin.H {
	return gin.H{"error": message}
}
