package handlers

import (
	"net/http"
	"strconv"

	"github.com/gavrylenkoIvan/gonotes/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) verifyUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	code := c.Param("token")

	if err := h.services.VerifyUser(userId, code); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "http://localhost:8080/")
}

func (h *Handler) undoUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	code := c.Param("token")

	if err := h.services.DeleteUser(userId, code); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "http://localhost:8080/")
}
