package repository

import (
	"time"

	"github.com/vikas-gautam/hotel-booking-app/internal/models"
)

type DatabaseRepo interface {
	InsertReservation(models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(roomID int, start, end time.Time) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
