package data

// EmptySlotHandler ...
type EmptySlotHandler interface {
	Insert(item int) error
	GetNearestSlot() (int, error)
}

// ParkingLot represents a parking lot. It provides some auxiliary data
// structures for efficient queries.
type ParkingLot struct {
	N          int
	EmptySlots EmptySlotHandler
	RegToColor map[string]string
	ColorToReg map[string][]string
	RegToSlot  map[string]int
	SlotToReg  map[int]string
}

// CarInPark represent a car in the parking.
type CarInPark struct {
	Slot  int
	Reg   string
	Color string
}
