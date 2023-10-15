package handler

import (
	"github.com/atadzan/more-tech/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	app.RedirectTrailingSlash = true
	app.RedirectFixedPath = true

	app.GET("/atms", h.getAtms)
	app.GET("/atms/:id", h.getAtmById)
	app.GET("/offices", h.getOffices)
	app.GET("/offices/:id", h.getOfficeById)

	return app
}
