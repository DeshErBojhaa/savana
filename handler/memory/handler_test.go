package memory

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/DeshErBojhaa/gojeck/parking_lot/data"
	"github.com/DeshErBojhaa/gojeck/parking_lot/handler"
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
		s, err := h.ParkCar("xx"+strconv.Itoa(i), "XX")
		if err != nil {
			t.Fatal(err)
		}
		if s != i {
			t.Logf("%s Expected slot %d Found slot %d", Failed, i, s)
		} else {
			t.Logf("%s Allocated slot number: %d", Success, s)
		}
	}
	_, err = h.ParkCar("xxXX", "XXX")
	if errors.Is(err, handler.ErrParkingFull) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, handler.ErrParkingFull)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, handler.ErrParkingFull)
		t.Fail()
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
		t.Fail()
	}
}

func TestCarAlreadyExists(t *testing.T) {
	h, err := NewLotHandler(3)
	if err != nil {
		t.Fatal(err)
	}
	_, err = h.ParkCar("Reg", "Black")
	if err != nil {
		t.Fatal(err)
	}

	s, err := h.ParkCar("Reg", "Black")
	if s != -1 {
		t.Logf("%s Expected slot: -1   Found slot: %d", Failed, s)
		t.Fail()
	}
	if errors.Is(err, handler.ErrAlreadyExistsInParking) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, handler.ErrAlreadyExistsInParking)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, handler.ErrAlreadyExistsInParking)
		t.Fail()
	}
}

func TestLeavCar(t *testing.T) {
	h, err := NewLotHandler(3)
	if err != nil {
		t.Fatal(err)
	}
	err = h.LeaveCar(2)
	if errors.Is(err, handler.ErrEmptySlot) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, handler.ErrEmptySlot)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, ErrHeapEmpty)
		t.Fail()
	}

	err = h.LeaveCar(-1)
	if errors.Is(err, handler.ErrSlotOutOfRange) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, handler.ErrSlotOutOfRange)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, handler.ErrSlotOutOfRange)
		t.Fail()
	}

	err = h.LeaveCar(4)
	if errors.Is(err, handler.ErrSlotOutOfRange) {
		t.Logf("%s Expected err: %v    Got err: %v", Success, err, handler.ErrSlotOutOfRange)
	} else {
		t.Logf("%s Expected err: %v    Got err: %v", Failed, err, handler.ErrSlotOutOfRange)
		t.Fail()
	}
}

func TestGetStatus(t *testing.T) {
	h, err := NewLotHandler(3)
	if err != nil {
		t.Fatal(err)
	}
	_, err = h.ParkCar("ABC", "Red")
	if err != nil {
		t.Fatal(err)
	}

	_, err = h.ParkCar("DEF", "Red")
	if err != nil {
		t.Fatal(err)
	}

	_, err = h.ParkCar("GHI", "Blue")
	if err != nil {
		t.Fatal(err)
	}

	want := []data.CarInPark{
		data.CarInPark{Slot: 1, Color: "Red", Reg: "ABC"},
		data.CarInPark{Slot: 2, Color: "Red", Reg: "DEF"},
		data.CarInPark{Slot: 3, Color: "Blue", Reg: "GHI"},
	}
	status := h.GetStatus()
	for i, s := range status {
		if reflect.DeepEqual(s, want[i]) {
			t.Logf("%s Expected Car: %#v    Got Car: %#v", Success, want[i], s)
		} else {
			t.Logf("%s Expected Car: %#v    Got Car: %#v", Failed, want[i], s)
			t.Fail()
		}
	}
}
