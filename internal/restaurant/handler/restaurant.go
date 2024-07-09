package handler

import (
	"net/http"
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/service"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler interface {
	SelfCheck(
		ctx *gin.Context,
	)
}

type RestaurantHandlerImpl struct {
	route             *gin.Engine
	config            *configs.Config
	restaurantService service.RestaurantService
}

var _ RestaurantHandler = (*RestaurantHandlerImpl)(nil)

func NewRestaurantHandler(
	route *gin.Engine,
	config *configs.Config,
	restaurantService service.RestaurantService,
) *RestaurantHandlerImpl {
	restaurantHandler := &RestaurantHandlerImpl{
		route:             route,
		config:            config,
		restaurantService: restaurantService,
	}
	return restaurantHandler
}

func (h *RestaurantHandlerImpl) SelfCheck(
	ctx *gin.Context,
) {
	result, err := h.restaurantService.SelfCheck(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
