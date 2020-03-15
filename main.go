package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DeshErBojhaa/gojeck/parking_lot/app"
	"github.com/DeshErBojhaa/gojeck/parking_lot/handler/memory"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
	log.Println("Application Terminating!")
	os.Exit(0)
}

func run() error {
	logD := log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	logE := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

	logD.Println("Started: Initializing Application")
	defer logD.Println("Execution Complete")

	// Make a channel to listen for interuptions or termination signals from the OS.
	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	// Channel that listens for any ireversable fatal errors cause by the application.
	appFatalError := make(chan error, 1)

	inmemoryHandler, err := memory.NewLotHandler(0)
	if err != nil {
		logE.Printf("Error while creating handler %v", err)
		return err
	}
	// Make parkinglot here
	app := app.App{LogD: logD, LogE: logE, Handler: inmemoryHandler}
	go func() {
		appFatalError <- app.Serve()
	}()

	select {
	case err := <-appFatalError:
		if err != nil {
			return fmt.Errorf("Fatal Error: %w", err)
		}
	case sig := <-shutDown:
		logE.Printf("main: %v, start shutdown ", sig)
		err := app.CleanUp()
		if err != nil {
			logE.Printf("main: graceful shutdown failed %v", err)
			return err
		}
		if sig == syscall.SIGSTOP {
			return errors.New("integrity issue caused shutdown")
		}
	}
	return nil
}
