package memory

import (
	"errors"
	"testing"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestParkCar(t *testing.T) {
	h, err := NewLotHandler(3)
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i <= 3; i++ {
		s, err := h.ParkCar("xx", "XX")
		if err != nil {
			t.Fatal(err)
		}
		if s != i {
			t.Logf("%s Expected slot %d Found slot %d", Failed, i, s)
		} else {
			t.Logf("%s Allocated slot number: %d", Success, s)
		}
	}
	_, err = h.ParkCar("xx", "XXX")
	if errors.Is(err, ErrParkingFull) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, ErrParkingFull)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, ErrParkingFull)
	}

	err = h.LeaveCar(2)
	if err != nil {
		t.Fatal(err)
	}
	s, err := h.ParkCar("xx", "XXX")
	if err != nil {
		t.Fatal(err)
	}
	if s == 2 {
		t.Logf("%s Expected slot: 2   Found slot: %d", Success, s)
	} else {
		t.Logf("%s Expected slot: 2   Found slot: %d", Failed, s)
	}
}

func TestLeavCar(t *testing.T) {
	h, err := NewLotHandler(3)
	if err != nil {
		t.Fatal(err)
	}
	err = h.LeaveCar(2)
	if errors.Is(err, ErrEmptySlot) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, ErrEmptySlot)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, ErrHeapEmpty)
	}

	err = h.LeaveCar(-1)
	if errors.Is(err, ErrSlotOutOfRange) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, ErrSlotOutOfRange)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, ErrSlotOutOfRange)
	}

	err = h.LeaveCar(4)
	if errors.Is(err, ErrSlotOutOfRange) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, ErrSlotOutOfRange)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, ErrSlotOutOfRange)
	}
}
