package app

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/DeshErBojhaa/gojeck/parking_lot/handler/memory"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestExecutionInstruction(t *testing.T) {
	//t.SkipNow()

	instructions := []string{
		"create_parking_lot 6",
		"park KA-01-HH-1234 White",
		"park KA-01-HH-9999 White",
		"park KA-01-BB-0001 Black",
		"park KA-01-HH-7777 Red",
		"park KA-01-HH-2701 Blue",
		"park KA-01-HH-3141 Black",
		"leave 4",
		"park KA-01-P-333 White",
		"park DL-12-AA-9999 White",
		"registration_numbers_for_cars_with_colour White",
		"slot_numbers_for_cars_with_colour White",
		"slot_number_for_registration_number KA-01-HH-3141",
		"slot_number_for_registration_number MH-04-AY-1111",
	}
	expectedRes := []string{
		"Created a parking lot with 6 slots",
		"Allocated slot number: 1",
		"Allocated slot number: 2",
		"Allocated slot number: 3",
		"Allocated slot number: 4",
		"Allocated slot number: 5",
		"Allocated slot number: 6",
		"Slot number 4 is free",
		"Allocated slot number: 4",
		"Sorry, parking lot is full",
		"KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333",
		"1, 2, 4",
		"6",
		"Not found",
	}

	inmemoryHandler, err := memory.NewLotHandler(0)
	if err != nil {
		t.Fatal(err)
	}

	app := App{LogD: &log.Logger{}, Handler: inmemoryHandler}

	for i, ins := range instructions {
		r, w, _ := os.Pipe()
		os.Stdout = w
		app.ExecInstruction(ins)
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		var result = strings.TrimSuffix(buf.String(), "\n")
		expected := expectedRes[i]
		if result != expected {
			t.Logf("%s\nWant: \n%v\nGot: \n%v\n", Failed, expected, result)
			t.Fail()
		} else {
			t.Logf("%v Instruction : %v", Success, ins)
		}
	}
}
