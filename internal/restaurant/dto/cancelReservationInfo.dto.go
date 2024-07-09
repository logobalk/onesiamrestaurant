package dto

type CancelReservationInfo struct {
	BookingID           string
	NumberofFreedTables int
	RemainingTables     int
}
