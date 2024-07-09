package query

import (
	"context"
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestaurantQuery_ReserveTable(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(t *testing.T, rq *RestaurantQueryImpl)
		act     func(t *testing.T, restaurantQuery RestaurantQuery) dto.ReserveInfo
		assert  func(t *testing.T, response dto.ReserveInfo)
	}{
		{
			name: "should success when table is available and customer is less than table capacity",
			arrange: func(t *testing.T, rq *RestaurantQueryImpl) {
				rq.availableTables = 3
				rq.reservedTables = 1
				rq.booking = make(map[string]int)
			},
			act: func(t *testing.T, restaurantQuery RestaurantQuery) dto.ReserveInfo {
				return restaurantQuery.ReserveTable(context.TODO(), 3)
			},
			assert: func(t *testing.T, response dto.ReserveInfo) {
				assert.Equal(t, "ONE-000001", response.BookingID)
				assert.Equal(t, 1, response.NumberOfBookedTables)
				assert.Equal(t, 2, response.RemainingTables)
			},
		},
		{
			name: "should fully booking when table is available and customer is more than table capacity",
			arrange: func(t *testing.T, rq *RestaurantQueryImpl) {
				rq.availableTables = 3
				rq.reservedTables = 1
				rq.booking = make(map[string]int)
			},
			act: func(t *testing.T, restaurantQuery RestaurantQuery) dto.ReserveInfo {
				return restaurantQuery.ReserveTable(context.TODO(), 13)
			},
			assert: func(t *testing.T, response dto.ReserveInfo) {
				assert.Equal(t, "", response.BookingID)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockConfigs := &configs.Config{
				MaxTableCapacity: 4,
			}
			RestaurantQuery := NewRestaurantQuery(mockConfigs)
			tt.arrange(t, RestaurantQuery)
			response := tt.act(t, RestaurantQuery)
			tt.assert(t, response)
		})
	}
}

func TestRestaurantQuery_CancelReservation(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(t *testing.T, rq *RestaurantQueryImpl)
		act     func(t *testing.T, restaurantQuery RestaurantQuery) dto.CancelReservationInfo
		assert  func(t *testing.T, response dto.CancelReservationInfo)
	}{
		{
			name: "should success when booking is found",
			arrange: func(t *testing.T, rq *RestaurantQueryImpl) {
				rq.availableTables = 3
				rq.reservedTables = 1
				rq.booking = make(map[string]int)
				rq.booking["ONE-000001"] = 3
			},
			act: func(t *testing.T, restaurantQuery RestaurantQuery) dto.CancelReservationInfo {
				return restaurantQuery.CancelReservation(context.TODO(), "ONE-000001")
			},
			assert: func(t *testing.T, response dto.CancelReservationInfo) {
				assert.Equal(t, "ONE-000001", response.BookingID)
				assert.Equal(t, 1, response.NumberofFreedTables)
				assert.Equal(t, 4, response.RemainingTables)
			},
		},
		{
			name: "should not cancel when booking is not found",
			arrange: func(t *testing.T, rq *RestaurantQueryImpl) {
				rq.availableTables = 3
				rq.reservedTables = 1
				rq.booking = make(map[string]int)
			},
			act: func(t *testing.T, restaurantQuery RestaurantQuery) dto.CancelReservationInfo {
				return restaurantQuery.CancelReservation(context.TODO(), "ONE-000001")
			},
			assert: func(t *testing.T, response dto.CancelReservationInfo) {
				assert.Equal(t, "", response.BookingID)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockConfigs := &configs.Config{
				MaxTableCapacity: 4,
			}
			RestaurantQuery := NewRestaurantQuery(mockConfigs)
			tt.arrange(t, RestaurantQuery)
			response := tt.act(t, RestaurantQuery)
			tt.assert(t, response)
		})
	}
}
