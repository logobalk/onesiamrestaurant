package controller

import (
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/handler"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	route             *gin.Engine
	config            *configs.Config
	restaurantHandler handler.RestaurantHandler
}

func New(
	route *gin.Engine,
	config *configs.Config,
	restaurentHandler handler.RestaurantHandler,
) *Controller {
	controller := &Controller{
		route:             route,
		config:            config,
		restaurantHandler: restaurentHandler,
	}
	controller.SetupRoute()
	return controller
}

func (c *Controller) SetupRoute() {
}
