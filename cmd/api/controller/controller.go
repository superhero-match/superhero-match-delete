package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/superhero-match-delete/cmd/api/service"
	"github.com/superhero-match-delete/internal/config"
)

// Controller holds the Controller data.
type Controller struct {
	Service *service.Service
}

// NewController returns new controller.
func NewController(cfg *config.Config) (*Controller, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &Controller{
		Service: srv,
	}, nil
}

// RegisterRoutes registers all the superhero_suggestions API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/match")

	// Middleware.
	// sr.Use(c.Authorize)

	// Routes.
	sr.POST("/delete", ctl.Delete)

	return router
}