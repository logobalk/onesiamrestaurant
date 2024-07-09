package service

import (
	"context"
	"one-siam-restaurant/configs"
)

type RestaurantService interface {
	SelfCheck(
		ctx context.Context,
	) (string, error)
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

func (s *RestaurantServiceImpl) SelfCheck(
	ctx context.Context,
) (string, error) {
	return "OK", nil
}
