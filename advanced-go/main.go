package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

func (t Truck) LoadCargo() error {
	return ErrTruckNotFound
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo %s: %w\n", truck.id, err)
	}

	return ErrNotImplemented
}

func main() {
	trucks := []Truck{
		Truck{id: "Truck-1"},
		Truck{id: "Truck-2"},
		Truck{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck %s : %s\n", truck.id, err)
		}
		switch err {
		case ErrTruckNotFound:
			return
		case ErrNotImplemented:
			return
		default:
			log.Fatal(err)
		}

		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrNotImplemented) {
				// we do this
			}
			if errors.Is(err, ErrTruckNotFound) {
				// we do this
			}

			if errors.Is(err, ErrTruckNotFound) {
				log.Fatal("TRUE")
			}

			log.Fatalf("Error processing truck %s : %s\n", truck.id, err)
		}
	}

}
