package query

import (
	"context"
	"fmt"
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant/dto"
)

type RestaurantQuery interface {
	InitializeRestaurant(
		ctx context.Context,
		tables int,
	) bool

	IsInitialzed() bool

	ReserveTable(
		ctx context.Context,
		numCustomers int,
	) dto.ReserveInfo

	IsReserved(
		ctx context.Context,
		bookingID string,
	) bool

	CancelReservation(
		ctx context.Context,
		bookingID string,
	) dto.CancelReservationInfo
}

type RestaurantQueryImpl struct {
	config          *configs.Config
	availableTables int // default 0
	reservedTables  int // default 0
	initialized     bool
	booking         map[string]int
}

var _ RestaurantQuery = (*RestaurantQueryImpl)(nil)

func NewRestaurantQuery(
	config *configs.Config,
) *RestaurantQueryImpl {
	restaurantQuery := &RestaurantQueryImpl{
		config: config,
	}
	return restaurantQuery
}

func (q *RestaurantQueryImpl) InitializeRestaurant(
	ctx context.Context,
	tables int,
) bool {
	if !q.initialized {
		q.availableTables = tables
		q.initialized = true
		q.booking = make(map[string]int)
		return true
	}
	return false
}

func (q *RestaurantQueryImpl) ReserveTable(
	ctx context.Context,
	numCustomers int,
) dto.ReserveInfo {
	var reserveInfo dto.ReserveInfo
	numberOfNeededTables := numCustomers / q.config.MaxTableCapacity
	if numCustomers%q.config.MaxTableCapacity != 0 {
		numberOfNeededTables++
	}

	if int(numberOfNeededTables) <= q.availableTables {
		bookingID := fmt.Sprintf("ONE-%06d", len(q.booking)+1)
		q.booking[bookingID] = numCustomers
		q.availableTables -= int(numberOfNeededTables)
		q.reservedTables += int(numberOfNeededTables)

		reserveInfo.BookingID = bookingID
		reserveInfo.NumberOfBookedTables = int(numberOfNeededTables)
		reserveInfo.RemainingTables = q.availableTables
	}
	return reserveInfo
}

func (q *RestaurantQueryImpl) IsInitialzed() bool {
	return q.initialized
}

func (q *RestaurantQueryImpl) IsReserved(
	ctx context.Context,
	bookingID string,
) bool {
	_, ok := q.booking[bookingID]
	return ok
}

func (q *RestaurantQueryImpl) CancelReservation(
	ctx context.Context,
	bookingID string,
) dto.CancelReservationInfo {
	var cancelReservationInfo dto.CancelReservationInfo
	freedCustomers := q.booking[bookingID]
	freedTables := 0
	if freedCustomers > 0 {
		freedTables = freedCustomers / q.config.MaxTableCapacity
		if freedCustomers%q.config.MaxTableCapacity != 0 {
			freedTables++
		}
		q.availableTables += freedTables
		q.reservedTables -= freedTables
		delete(q.booking, bookingID)

		cancelReservationInfo.BookingID = bookingID
		cancelReservationInfo.NumberofFreedTables = freedTables
		cancelReservationInfo.RemainingTables = q.availableTables
	}
	return cancelReservationInfo
}
