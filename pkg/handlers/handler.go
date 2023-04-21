package handlers

import (
	"github.com/gavrylenkoIvan/gonotes/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		verify := auth.Group("/:user_id/verify")
		{
			verify.GET("/:token", h.verifyUser)
		}
	}

	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.GET("/", h.getAllNotes)
			notes.POST("/", h.createNote)
			notes.DELETE("/:id", h.deleteNote)
			notes.PUT("/:id", h.updateNote)
		}
	}

	return router
}
