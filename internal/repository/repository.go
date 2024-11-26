package repository

import "github.com/sinakovs/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRestriction(res models.RoomRestriction) error
}
