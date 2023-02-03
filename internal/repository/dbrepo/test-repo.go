package dbrepo

import (
	"time"

	"github.com/vikas-gautam/hotel-booking-app/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil

}

func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID and false if no availablity.
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(roomID int, start, end time.Time) (bool, error) {

	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms if any for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	return []models.Room{}, nil

}

func (m *testDBRepo) GetRoomByID(roomId int) (models.Room, error) {

	return models.Room{}, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {

	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {

	return nil
}

// Authenticate authenticate users
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {

	var id int
	var hashedPassword string

	return id, hashedPassword, nil
}
