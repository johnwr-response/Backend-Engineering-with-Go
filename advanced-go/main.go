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
	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}

	person := make(map[string]any)
	person["name"] = "John Doe"
	person["age"] = 42
	age, exists := person["age"].(int)
	if !exists {
		log.Fatal("age is not an integer")
		return
	}
	log.Println(age)

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error processing truck %s\n", err)
	}
	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error processing truck %s\n", err)
	}
	log.Println(nt.cargo)
	log.Println(et.battery)

}
