package repository

import "github.com/vikas-gautam/hotel-booking-app/internal/models"

type DatabaseRepo interface {
	InsertReservation(models.Reservation) error
}
