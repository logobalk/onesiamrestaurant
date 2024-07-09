package service

import (
	"one-siam-restaurant/configs"
)

type RestaurantService interface {
}

type RestaurantServiceImpl struct {
	config *configs.Config
}

var _ RestaurantService = (*RestaurantServiceImpl)(nil)

func NewRestaurantService(
	config *configs.Config,
) *RestaurantServiceImpl {
	restaurantService := &RestaurantServiceImpl{
		config: config,
	}
	return restaurantService
}
