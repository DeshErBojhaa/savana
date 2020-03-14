package memory

import (
	"errors"
	"sort"
)

// ErrParkingFull is a package level error to represent a full parking.
var ErrParkingFull = errors.New("Sorry, parking lot is full")

// ErrEmptySlot when slot already empty.
var ErrEmptySlot = errors.New("Slot already empty")

// ErrSlotOutOfRange when slot already empty.
var ErrSlotOutOfRange = errors.New("Slot number out of range")

// ErrInconsistent when internal server error occurs.
var ErrInconsistent = errors.New("Application in inconsistant state")

// LotHandler represents a parking lot. It provides some auxiliary data
// structures for efficient queries.
type LotHandler struct {
	N          int
	emptySlots *Minheap
	regToColor map[string]string
	colorToReg map[string][]string
	regToSlot  map[string]int
	slotToReg  map[int]string
}

// NewLotHandler returns a in memory ephemeral store for a parking lot.
func NewLotHandler(n int) (*LotHandler, error) {
	state := LotHandler{
		N:          n,
		emptySlots: NewMinHeap(n),
		regToColor: map[string]string{},
		colorToReg: map[string][]string{},
		regToSlot:  map[string]int{},
		slotToReg:  map[int]string{},
	}
	return &state, nil
}

// ParkCar adds a new car to the parking lot.
func (s *LotHandler) ParkCar(reg, color string) (int, error) {
	slot, err := s.emptySlots.GetNearestSlot()
	if errors.Is(err, ErrHeapEmpty) {
		return -1, ErrParkingFull
	}
	s.colorToReg[color] = append(s.colorToReg[color], reg)
	s.regToColor[reg] = color
	s.regToSlot[reg] = slot
	s.slotToReg[slot] = reg
	return slot, nil
}

// LeaveCar handels the parking state when a car leaves the parking lot.
func (s *LotHandler) LeaveCar(slot int) error {
	if slot > s.N || slot < 1 {
		return ErrSlotOutOfRange
	}
	reg, ok := s.slotToReg[slot]
	if !ok {
		return ErrEmptySlot
	}
	color, ok := s.regToColor[reg]
	if !ok {
		return ErrInconsistent
	}

	// 1. Remove this car from the reg list or colorToReg map.
	for i := 0; i < len(s.colorToReg[color]); i++ {
		if s.colorToReg[color][i] == reg {
			s.colorToReg[color] = append(s.colorToReg[color], s.colorToReg[color]...)
			break
		}
	}
	// 2. Remove this car from regToColor map.
	delete(s.regToColor, reg)
	// 3. Remove this car from regToSlot map.
	delete(s.regToSlot, reg)
	// 4. Remove this car from slotToReg map.
	delete(s.slotToReg, slot)

	// Add this slot into the empty slot heap
	if err := s.emptySlots.Insert(slot); err != nil {
		return ErrInconsistent
	}
	return nil
}

// RegNoOfCarsOfColor returns all the cars registration number whose color is a match.
func (s *LotHandler) RegNoOfCarsOfColor(color string) []string {
	return s.colorToReg[color]
}

// SlotOfCarsOfColor returns all the slots where a car is parked of color.
func (s *LotHandler) SlotOfCarsOfColor(color string) []int {
	var slots []int
	for _, reg := range s.colorToReg[color] {
		slots = append(slots, s.regToSlot[reg])
	}
	return slots
}

// CarInPark represent a car in the parking.
type CarInPark struct {
	Slot  int
	Reg   string
	Color string
}

// GetStatus returns all the cars detail in the parking in their ascending slot order.
func (s *LotHandler) GetStatus() []CarInPark {
	var status []CarInPark
	for reg, col := range s.regToColor {
		status = append(status, CarInPark{Slot: s.regToSlot[reg], Color: col, Reg: reg})
	}
	sort.SliceStable(status, func(i int, j int) bool {
		return status[i].Slot < status[j].Slot
	})
	return status
}
