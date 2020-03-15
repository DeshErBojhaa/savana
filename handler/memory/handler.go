package memory

import (
	"errors"
	"sort"

	"github.com/DeshErBojhaa/gojeck/parking_lot/data"
)

// ErrParkingFull is a package level error to represent a full parking.
var ErrParkingFull = errors.New("Sorry, parking lot is full")

// ErrEmptySlot when slot already empty.
var ErrEmptySlot = errors.New("Slot already empty")

// ErrSlotOutOfRange when slot already empty.
var ErrSlotOutOfRange = errors.New("Slot number out of range")

// ErrInconsistent when internal server error occurs.
var ErrInconsistent = errors.New("Application in inconsistant state")

// ErrAlreadyExistsInParking when internal server error occurs.
var ErrAlreadyExistsInParking = errors.New("Car with same reg already exists in parking")

// InMemoryHandler is ephimeral handler for parking lot.
type InMemoryHandler struct {
	parkingLot data.ParkingLot
}

// NewLotHandler returns a in memory ephemeral store for a parking lot.
func NewLotHandler(n int) (*InMemoryHandler, error) {
	handler := InMemoryHandler{
		parkingLot: data.ParkingLot{
			N:          n,
			EmptySlots: data.NewMinHeap(n),
			RegToColor: map[string]string{},
			ColorToReg: map[string][]string{},
			RegToSlot:  map[string]int{},
			SlotToReg:  map[int]string{},
		}}
	return &handler, nil
}

// ParkCar adds a new car to the parking lot.
func (h *InMemoryHandler) ParkCar(reg, color string) (int, error) {
	if _, ok := h.parkingLot.RegToSlot[reg]; ok {
		return -1, ErrAlreadyExistsInParking
	}
	slot, err := h.parkingLot.EmptySlots.GetNearestSlot()
	if errors.Is(err, data.ErrHeapEmpty) {
		return -1, ErrParkingFull
	}
	h.parkingLot.ColorToReg[color] = append(h.parkingLot.ColorToReg[color], reg)
	h.parkingLot.RegToColor[reg] = color
	h.parkingLot.RegToSlot[reg] = slot
	h.parkingLot.SlotToReg[slot] = reg
	return slot, nil
}

// LeaveCar handels the parking state when a car leaves the parking lot.
func (h *InMemoryHandler) LeaveCar(slot int) error {
	if slot > h.parkingLot.N || slot < 1 {
		return ErrSlotOutOfRange
	}
	reg, ok := h.parkingLot.SlotToReg[slot]
	if !ok {
		return ErrEmptySlot
	}
	color, ok := h.parkingLot.RegToColor[reg]
	if !ok {
		return ErrInconsistent
	}

	// 1. Remove this car from the reg list or colorToReg map.
	for i := 0; i < len(h.parkingLot.ColorToReg[color]); i++ {
		if h.parkingLot.ColorToReg[color][i] == reg {
			h.parkingLot.ColorToReg[color] = append(h.parkingLot.ColorToReg[color], h.parkingLot.ColorToReg[color]...)
			break
		}
	}
	// 2. Remove this car from regToColor map.
	delete(h.parkingLot.RegToColor, reg)
	// 3. Remove this car from regToSlot map.
	delete(h.parkingLot.RegToSlot, reg)
	// 4. Remove this car from slotToReg map.
	delete(h.parkingLot.SlotToReg, slot)

	// Add this slot into the empty slot heap
	if err := h.parkingLot.EmptySlots.Insert(slot); err != nil {
		return ErrInconsistent
	}
	return nil
}

// RegNoOfCarsOfColor returns all the cars registration number whose color is a match.
func (h *InMemoryHandler) RegNoOfCarsOfColor(color string) []string {
	return h.parkingLot.ColorToReg[color]
}

// SlotOfCarsOfColor returns all the slots where a car is parked of color.
func (h *InMemoryHandler) SlotOfCarsOfColor(color string) []int {
	var slots []int
	for _, reg := range h.parkingLot.ColorToReg[color] {
		slots = append(slots, h.parkingLot.RegToSlot[reg])
	}
	return slots
}

// GetStatus returns all the cars detail in the parking in their ascending slot order.
func (h *InMemoryHandler) GetStatus() []data.CarInPark {
	var status []data.CarInPark
	for reg, col := range h.parkingLot.RegToColor {
		status = append(status, data.CarInPark{Slot: h.parkingLot.RegToSlot[reg], Color: col, Reg: reg})
	}
	sort.SliceStable(status, func(i int, j int) bool {
		return status[i].Slot < status[j].Slot
	})
	return status
}
