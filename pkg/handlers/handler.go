package handlers

import (
	"github.com/gavrylenkoIvan/gonotes/pkg/service"
	"github.com/gin-contrib/static"
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

const (
	invCred = "invalid input body! Check your credentials and try again"
	url     = "http://localhost:8080"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("dist", false)))
	router.Use(static.Serve("/assets", static.LocalFile("dist\\assets", false)))
	router.Use(static.Serve("/favicon.ico", static.LocalFile("dist\\favicon.ico", false)))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		verify := auth.Group("/:user_id/verify")
		{
			verify.GET("/:token", h.verifyUser)
			verify.GET("/:token/undo", h.undoUser)
		}
	}

	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes", h.userIdentity)
		{
			notes.GET("/", h.getAllNotes)
			notes.POST("/", h.createNote)
			notes.DELETE("/:id", h.deleteNote)
			notes.PUT("/:id", h.updateNote)
		}
	}

	return router
}
