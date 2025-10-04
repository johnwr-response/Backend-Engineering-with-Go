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

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1
	return nil
}
func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

/*
// UnloadCargo : Note, this is the same as the one above
func UnloadCargo(e *ElectricTruck) error {
	e.cargo = 0
	e.battery += -1
	return nil
}
*/

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)
	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("Error loading cargo: %w\n", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %w\n", err)
	}

	return nil
}

func main() {
	// first example
	truckID := 42
	anotherTruckID := &truckID
	log.Println(truckID)
	log.Println(&truckID)
	log.Println(anotherTruckID)
	truckID = 0
	log.Println(*anotherTruckID)

	// second example
	t := NormalTruck{cargo: 0}
	fillTruckCargo(&t)
	log.Println(t)

	// third example
	var userID1 int
	log.Println(userID1)
	var userID2 *int
	log.Println(userID2)
	userID2 = &userID1
	log.Println(*userID2)

	// fourth example
	t = NormalTruck{cargo: 0}
	log.Printf("Address of t:%p\n", &t)
	fillTruckCargo2(t)
}

func fillTruckCargo2(t NormalTruck) {
	t.cargo = 100
	log.Printf("Address of t:%p\n", &t)
}
func fillTruckCargo(t *NormalTruck) {
	t.cargo = 100
}
