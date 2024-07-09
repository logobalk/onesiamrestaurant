package service

import (
	"context"
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/dto"
	"one-siam-restaurant/internal/restaurant/query"
)

type RestaurantService interface {
	SelfCheck(
		ctx context.Context,
	) (string, error)

	Initialize(
		ctx context.Context,
		numTables int,
	) (dto.ResponseDTO, error)

	ReserveTable(
		ctx context.Context,
		numCustomers int,
	) (dto.ResponseDTO, error)

	CancelReservation(
		ctx context.Context,
		bookingID string,
	) (dto.ResponseDTO, error)
}

type RestaurantServiceImpl struct {
	config          *configs.Config
	restaurantQuery query.RestaurantQuery
}

var _ RestaurantService = (*RestaurantServiceImpl)(nil)

func NewRestaurantService(
	config *configs.Config,
	restaurantQuery query.RestaurantQuery,
) *RestaurantServiceImpl {
	restaurantService := &RestaurantServiceImpl{
		config:          config,
		restaurantQuery: restaurantQuery,
	}
	return restaurantService
}

func (s *RestaurantServiceImpl) SelfCheck(
	ctx context.Context,
) (string, error) {
	return "OK", nil
}

func (s *RestaurantServiceImpl) Initialize(
	ctx context.Context,
	numTables int,
) (dto.ResponseDTO, error) {
	if s.restaurantQuery.InitializeRestaurant(ctx, numTables) {
		return dto.ResponseDTO{
			Status:  "Success",
			Message: "OK",
			Data:    nil,
		}, nil
	}
	return dto.ResponseDTO{
		Status:  "Error",
		Message: "Error this API is called again after initialization",
		Data:    nil,
	}, nil
}

func (s *RestaurantServiceImpl) ReserveTable(
	ctx context.Context,
	numCustomers int,
) (dto.ResponseDTO, error) {
	if !s.restaurantQuery.IsInitialzed() {
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Restaurant not initialized",
			Data:    nil,
		}, nil
	}
	if numCustomers <= 0 {
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Number of customers must be greater than 0",
			Data:    nil,
		}, nil
	}

	reserveInfo := s.restaurantQuery.ReserveTable(ctx, numCustomers)
	if reserveInfo.BookingID == "" {
		//FULL
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Fully Reserved",
			Data:    nil,
		}, nil
	}

	return dto.ResponseDTO{
		Status:  "Success",
		Message: "OK",
		Data:    reserveInfo,
	}, nil
}

func (s *RestaurantServiceImpl) CancelReservation(
	ctx context.Context,
	bookingID string,
) (dto.ResponseDTO, error) {
	if !s.restaurantQuery.IsInitialzed() {
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Restaurant not initialized",
			Data:    nil,
		}, nil
	}
	if bookingID == "" {
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Invalid Booking ID",
			Data:    nil,
		}, nil
	}

	if !s.restaurantQuery.IsReserved(ctx, bookingID) {
		return dto.ResponseDTO{
			Status:  "Error",
			Message: "Booking ID not found",
			Data:    nil,
		}, nil
	}

	cancelInfo := s.restaurantQuery.CancelReservation(ctx, bookingID)
	return dto.ResponseDTO{
		Status:  "Success",
		Message: "OK",
		Data:    cancelInfo,
	}, nil
}
