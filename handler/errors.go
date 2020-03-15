package handler

import "errors"

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

// ErrInvalidReg when registration is not present in parking
var ErrInvalidReg = errors.New("Not found")

// ErrInvalidCapacity when capacity is < 1
var ErrInvalidCapacity = errors.New("Capacity must be a positive number")
