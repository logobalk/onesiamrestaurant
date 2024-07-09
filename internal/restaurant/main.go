package restaurant

import (
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/controller"
	"one-siam-restaurant/internal/restaurant/handler"
	"one-siam-restaurant/internal/restaurant/service"

	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, config *configs.Config) {
	restaurantService := service.NewRestaurantService(config)
	restaurantHandler := handler.NewRestaurantHandler(engine, config, restaurantService)
	controller.New(engine, config, restaurantHandler)
}
