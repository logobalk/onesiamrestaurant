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

	Initialize(
		ctx *gin.Context,
	)

	ReserveTable(
		ctx *gin.Context,
	)

	CancelReservation(
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

func (h *RestaurantHandlerImpl) Initialize(
	ctx *gin.Context,
) {
	var params map[string]interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	numTablesFloat, ok := params["number_of_tables"].(float64)
	numTables := int(numTablesFloat)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing number_of_tables"})
		return
	}
	result, err := h.restaurantService.Initialize(ctx, numTables)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (h *RestaurantHandlerImpl) ReserveTable(
	ctx *gin.Context,
) {
	var params map[string]interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	numCustomersFloat, ok := params["number_of_customer"].(float64)
	numCustomer := int(numCustomersFloat)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing number_of_customer"})
		return
	}

	result, err := h.restaurantService.ReserveTable(ctx, numCustomer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *RestaurantHandlerImpl) CancelReservation(
	ctx *gin.Context,
) {
	var params map[string]interface{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookingID, ok := params["booking_id"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing booking_id"})
		return
	}
	result, err := h.restaurantService.CancelReservation(ctx, bookingID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
