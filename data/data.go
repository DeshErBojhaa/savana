package data

// ParkingLot represents a parking lot. It provides some auxiliary data
// structures for efficient queries.
type ParkingLot struct {
	N          int
	EmptySlots *Minheap
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
