package app

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/DeshErBojhaa/gojeck/parking_lot/handler"
)

// App is the central
type App struct {
	LogD    *log.Logger
	Handler ParkingLotHandler
}

// Serve ...
func (a *App) Serve() error {
	if len(os.Args) > 1 && os.Args[1] != "" {
		a.LogD.Println("File Mode Enabled")

		file, err := os.Open(os.Args[1])
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			err := a.ExecInstruction(scanner.Text())
			if err != nil {
				return err
			}
		}
		if scanner.Err(); err != nil {
			return err
		}
	} else {
		a.LogD.Println("Interactive Mode Enabled")
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			instruction := scanner.Text()

			if instruction == "exit" {
				a.LogD.Println("Processing Exis command")
				return a.CleanUp()
			}
			if err := a.ExecInstruction(instruction); err != nil {
				return nil
			}
		}
	}
	return nil
}

// CleanUp purges any reuseable resources and prevents memory leak.
func (a *App) CleanUp() error {
	defer a.LogD.Println("All resources purged")
	return nil
}

func isValidateInstruction(instructions []string, reqParts int) bool {
	if len(instructions) != reqParts {
		return false
	}
	for _, s := range instructions {
		if s == "" {
			return false
		}
	}
	return true
}

// ExecInstruction executes the current instruction.
func (a *App) ExecInstruction(ins string) error {
	insParts := strings.Split(ins, " ")
	switch insParts[0] {
	case "create_parking_lot":
		if !isValidateInstruction(insParts, 2) {
			fmt.Println("Malformed instruction")
			return nil
		}

		capacity, err := strconv.Atoi(insParts[1])
		if err != nil {
			fmt.Printf("%v is not an integer\n", insParts[1])
			return nil
		}
		if err := a.Handler.SetCapacity(capacity); err != nil {
			fmt.Printf("%v\n", err)
			return nil
		}
		fmt.Printf("Created a parking lot with %d slots\n", capacity)

	case "park":
		if !isValidateInstruction(insParts, 3) {
			fmt.Println("Malformed instruction")
			return nil
		}

		slot, err := a.Handler.ParkCar(insParts[1], insParts[2])
		if err != nil {
			if errors.Is(err, handler.ErrAlreadyExistsInParking) || errors.Is(err, handler.ErrParkingFull) {
				fmt.Printf("%v\n", err)
				return nil
			}
			return err
		}
		fmt.Printf("Allocated slot number: %d\n", slot)

	case "leave":
		if !isValidateInstruction(insParts, 2) {
			fmt.Println("Malformed instruction")
			return nil
		}

		slot, err := strconv.Atoi(insParts[1])
		if err != nil {
			return err
		}
		err = a.Handler.LeaveCar(slot)
		if err != nil {
			if errors.Is(err, handler.ErrSlotOutOfRange) || errors.Is(err, handler.ErrEmptySlot) {
				fmt.Printf("%v", err)
				return nil
			}
			return err
		}
		fmt.Printf("Slot number %d is free\n", slot)

	case "status":
		fmt.Println("Slot No.    Registration No    Colour")
		cars := a.Handler.GetStatus()
		for _, c := range cars {
			// fmt.Printf("%-12d%-19s%-6s\n", c.Slot, c.Reg, c.Color)
			fmt.Printf("%d           %s      %s\n", c.Slot, c.Reg, c.Color)
		}

	case "registration_numbers_for_cars_with_colour":
		if !isValidateInstruction(insParts, 2) {
			fmt.Println("Malformed instruction")
			return nil
		}

		regNums := a.Handler.RegNoOfCarsOfColor(insParts[1])
		fmt.Println(strings.Join(regNums[:], ", "))

	case "slot_numbers_for_cars_with_colour":
		if !isValidateInstruction(insParts, 2) {
			fmt.Println("Malformed instruction")
			return nil
		}

		slots := a.Handler.SlotOfCarsOfColor(insParts[1])
		slotsStr := []string{}
		for _, s := range slots {
			slotsStr = append(slotsStr, strconv.Itoa(s))
		}
		fmt.Println(strings.Join(slotsStr[:], ", "))

	case "slot_number_for_registration_number":
		if !isValidateInstruction(insParts, 2) {
			fmt.Println("Malformed instruction")
			return nil
		}

		slot, err := a.Handler.SlotOfCar(insParts[1])
		if err != nil {
			if errors.Is(err, handler.ErrInvalidReg) {
				fmt.Printf("%v\n", err)
				return nil
			}
			return err
		}
		fmt.Println(slot)

	default:
		fmt.Println("Invalid instruction")
	}
	return nil
}
