package handlers

import (
	"net/http"
	"strconv"

	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllNotes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notes, err := h.services.GetNotesByUser(userId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notes,
	})
}

func (h *Handler) deleteNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteNote(id, userId); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.Status{
		Ok: true,
	})
}

func (h *Handler) updateNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var updatedNote gonotes.UpdateNoteStruct
	if err := c.BindJSON(&updatedNote); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateNote(id, userId, updatedNote); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.Status{
		Ok: true,
	})
}

func (h *Handler) createNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input gonotes.Note
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateNote(userId, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"note": id,
	})
}
