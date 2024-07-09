package handler

import (
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/service"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler interface {
}

type RestaurantHandlerImpl struct {
	route             *gin.Engine
	config            *configs.Config
	restaurantService service.RestaurantService
}

func NewRestaurantHandler(
	route *gin.Engine,
	config *configs.Config,
	restaurantService *service.RestaurantService,
) *RestaurantHandlerImpl {
	restaurantHandler := &RestaurantHandlerImpl{
		route:             route,
		config:            config,
		restaurantService: restaurantService,
	}
	return restaurantHandler
}
