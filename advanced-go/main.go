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
	return nil
}
func (t Truck) UnLoadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo %s: %w\n", truck.id, err)
	}
	if err := truck.UnLoadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo %s: %w\n", truck.id, err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		Truck{id: "Truck-1"},
		Truck{id: "Truck-2"},
		Truck{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		if err := processTruck(truck); err != nil {
			log.Fatalf("Error processing truck %s : %s\n", truck.id, err)
		}
	}

}
