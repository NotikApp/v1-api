package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gavrylenkoIvan/gonotes/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	authHeader      = "Authorization"
	userCtx         = "userId"
	incorrectHeader = "Incorrect authorization header!"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "incorrect auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "incorrect auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
