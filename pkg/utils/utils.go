package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type Status struct {
	Ok bool `json:"ok"`
}

func NewErrorResponse(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(code, errorResponse{
		Ok:      false,
		Message: err,
	})
	log.Println(err)
}
