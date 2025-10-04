package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
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
	fmt.Printf("started processing truck %+v\n", truck)

	// Simulate some processing time
	time.Sleep(1 * time.Second)

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w\n", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w\n", err)
	}

	fmt.Printf("finished processing truck %+v\n", truck)
	return nil
}

// processFleet demonstrates concurrent processing of multiple trucks
func processFleet(trucks []Truck) error {
	// Running synchronously
	/*
		for _, truck := range trucks {
			_ = processTruck(truck)
		}
	*/

	// Running asynchronously with wait groups
	var wg sync.WaitGroup
	//wg.Add(len(trucks)) // This would also work, instead of adding them one by one, but it's arguably better to do it within the loop
	for _, truck := range trucks {
		wg.Add(1)
		go func(truck Truck) {
			if err := processTruck(truck); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(truck)
	}
	wg.Wait()

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Printf("All trucks processed succesfully!")

	// waiting for all goroutines to finish if not using wait group
	//time.Sleep(5 * time.Second)

}
