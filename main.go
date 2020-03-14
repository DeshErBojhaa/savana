package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}

}

// ParkingLot ...
type ParkingLot struct{}

func (p *ParkingLot) shutDown() error {
	return nil
}

func run() error {
	logD := log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	logE := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

	logD.Println("Started: Initializing parking lot")
	defer logD.Println("Application terminated")

	// Make a channel to listen for interuptions or termination signals from the OS.
	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	// Make parkinglot here
	parkingLot := ParkingLot{}

	// Channel that listens for any ireversable fatal errors cause by the application.
	appFatalError := make(chan error, 1)

	select {
	case err := <-appFatalError:
		return fmt.Errorf("Fatal Error: %w", err)
	case sig := <-shutDown:
		logE.Printf("main: %v, start shutdown ", sig)
		err := parkingLot.shutDown()
		if err != nil {
			logE.Printf("main: graceful shutdown failed %v", err)
		}
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return fmt.Errorf("could not stop application gracefully: %w", err)
		}
	}
	return nil
}
