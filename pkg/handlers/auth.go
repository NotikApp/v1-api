package handlers

import (
	"fmt"
	"net/http"

	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input gonotes.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, invCred)
		return
	}

	temp := tempCode(16)
	id, err := h.services.Authorization.CreateUser(input, temp)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"ok": true,
	})

	sendEmail(input.Email, fmt.Sprintf("http://localhost:8080/auth/%d/verify/%s", id, temp), input.Username, fmt.Sprintf("http://localhost:8080/auth/%d/verify/%s/undo", id, temp))
}

func (h *Handler) signIn(c *gin.Context) {
	var input gonotes.SignInInput

	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, invCred)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"ok":    true,
	})
}
