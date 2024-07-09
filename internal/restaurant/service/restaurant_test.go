package service

import (
	"context"
	"one-siam-restaurant/configs"
	"testing"

	"one-siam-restaurant/internal/restaurant/dto"
	query_mocks "one-siam-restaurant/internal/restaurant/query/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRestaurantService_Initialize(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery)
		act     func(t *testing.T, restaurant RestaurantService) (dto.ResponseDTO, error)
		assert  func(t *testing.T, response dto.ResponseDTO, err error)
	}{
		{
			name: "should initial success when call InitializeRestaurant first time",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("InitializeRestaurant", mock.Anything, 10).Return(true)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.Initialize(context.TODO(), 10)
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Success", response.Status)
			},
		},
		{
			name: "should initial fail when call InitializeRestaurant second time",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("InitializeRestaurant", mock.Anything, 10).Return(false)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.Initialize(context.TODO(), 10)
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Error", response.Status)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restaurantQuery := query_mocks.NewRestaurantQuery(t)
			mockConfigs := &configs.Config{
				MaxTableCapacity: 4,
			}
			restaurantService := NewRestaurantService(mockConfigs, restaurantQuery)
			tt.arrange(t, restaurantQuery)
			response, err := tt.act(t, restaurantService)
			tt.assert(t, response, err)
		})
	}
}

func TestRestaurantService_ReserveTable(t *testing.T) {
	// errTesting := errors.New("Testing")
	tests := []struct {
		name    string
		arrange func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery)
		act     func(t *testing.T, restaurant RestaurantService) (dto.ResponseDTO, error)
		assert  func(t *testing.T, response dto.ResponseDTO, err error)
	}{
		{
			name: "should success when call after initialize",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(true)
				var reserveInfo dto.ReserveInfo
				reserveInfo.BookingID = "1234"
				reserveInfo.NumberOfBookedTables = 1
				reserveInfo.RemainingTables = 1
				restaurantQuery.On("ReserveTable", mock.Anything, 4).Return(reserveInfo)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.ReserveTable(context.TODO(), 4)
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Success", response.Status)
				assert.Equal(t, "1234", response.Data.(dto.ReserveInfo).BookingID)
				assert.Equal(t, 1, response.Data.(dto.ReserveInfo).NumberOfBookedTables)
				assert.Equal(t, 1, response.Data.(dto.ReserveInfo).RemainingTables)
			},
		},
		{
			name: "should error when call before initialize",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(false)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.ReserveTable(context.TODO(), 10)
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Error", response.Status)
			},
		},
		{
			name: "should error when customer is less than 1",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(true)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.ReserveTable(context.TODO(), 0)
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Error", response.Status)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restaurantQuery := query_mocks.NewRestaurantQuery(t)
			mockConfigs := &configs.Config{
				MaxTableCapacity: 4,
			}
			restaurantService := NewRestaurantService(mockConfigs, restaurantQuery)
			tt.arrange(t, restaurantQuery)
			response, err := tt.act(t, restaurantService)
			tt.assert(t, response, err)
		})
	}
}

func TestRestaurantService_CancelReservation(t *testing.T) {
	// errTesting := errors.New("Testing")
	tests := []struct {
		name    string
		arrange func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery)
		act     func(t *testing.T, restaurant RestaurantService) (dto.ResponseDTO, error)
		assert  func(t *testing.T, response dto.ResponseDTO, err error)
	}{
		{
			name: "should success when call after initialize",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(true)
				restaurantQuery.On("IsReserved", mock.Anything, "1234").Return(true)

				var cancelInfo dto.CancelReservationInfo
				cancelInfo.BookingID = "1234"
				cancelInfo.NumberofFreedTables = 1
				cancelInfo.RemainingTables = 1
				restaurantQuery.On("CancelReservation", mock.Anything, "1234").Return(cancelInfo)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.CancelReservation(context.TODO(), "1234")
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Success", response.Status)
				assert.Equal(t, "1234", response.Data.(dto.CancelReservationInfo).BookingID)
				assert.Equal(t, 1, response.Data.(dto.CancelReservationInfo).NumberofFreedTables)
				assert.Equal(t, 1, response.Data.(dto.CancelReservationInfo).RemainingTables)
			},
		},
		{
			name: "should error when call before initialize",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(false)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.CancelReservation(context.TODO(), "1234")
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Error", response.Status)
			},
		},
		{
			name: "should error when Boooking not found",
			arrange: func(t *testing.T, restaurantQuery *query_mocks.RestaurantQuery) {
				restaurantQuery.On("IsInitialzed").Return(true)
				restaurantQuery.On("IsReserved", mock.Anything, "1234").Return(false)
			},
			act: func(t *testing.T, restaurantService RestaurantService) (dto.ResponseDTO, error) {
				return restaurantService.CancelReservation(context.TODO(), "1234")
			},
			assert: func(t *testing.T, response dto.ResponseDTO, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Error", response.Status)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restaurantQuery := query_mocks.NewRestaurantQuery(t)
			mockConfigs := &configs.Config{
				MaxTableCapacity: 4,
			}
			restaurantService := NewRestaurantService(mockConfigs, restaurantQuery)
			tt.arrange(t, restaurantQuery)
			response, err := tt.act(t, restaurantService)
			tt.assert(t, response, err)
		})
	}
}
