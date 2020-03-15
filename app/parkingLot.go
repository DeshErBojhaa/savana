package app

import "github.com/DeshErBojhaa/gojeck/parking_lot/data"

// ParkingLotHandler is a interface that represents all the functionality
// provided by a parking lot system.
type ParkingLotHandler interface {
	SetCapacity(n int)
	ParkCar(reg, color string) (int, error)
	LeaveCar(slot int) error
	RegNoOfCarsOfColor(color string) []string
	SlotOfCarsOfColor(color string) []int
	SlotOfCar(reg string) (int, error)
	GetStatus() []data.CarInPark
}
