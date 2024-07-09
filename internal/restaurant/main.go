package restaurant

import (
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/controller"
	"one-siam-restaurant/internal/restaurant/handler"
	"one-siam-restaurant/internal/restaurant/query"
	"one-siam-restaurant/internal/restaurant/service"

	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, config *configs.Config) {
	restaurantQuery := query.NewRestaurantQuery(config)

	restaurantService := service.NewRestaurantService(config, restaurantQuery)
	restaurantHandler := handler.NewRestaurantHandler(engine, config, restaurantService)
	controller.New(engine, config, restaurantHandler)
}
