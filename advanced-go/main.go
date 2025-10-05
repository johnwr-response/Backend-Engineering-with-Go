package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"maps"
	"sync"
	"time"
)

type contextKey string

var UserIdKey contextKey = "userID"
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
func processTruck(ctx context.Context, truck Truck) error {
	//fmt.Printf("started processing truck %+v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// simulate a long-running process
	delay := 3 * time.Second
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w\n", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w\n", err)
	}

	//fmt.Printf("finished processing truck %+v\n", truck)
	return ErrTruckNotFound
}

// processFleet demonstrates concurrent processing of multiple trucks
func processFleet(ctx context.Context, trucks []Truck) error {

	var wg sync.WaitGroup
	errorsChan := make(chan error, len(trucks))

	for _, truck := range trucks {
		wg.Add(1)
		go func(truck Truck) {
			if err := processTruck(ctx, truck); err != nil {
				log.Println(err)
				errorsChan <- err
			}
			wg.Done()
		}(truck)
	}
	wg.Wait()
	defer close(errorsChan)

	var errs []error
	for err := range errorsChan {
		log.Printf("error processing truck %v\n", err)
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d error(s)", len(errs))
	}
	return nil

	//select {
	//case err := <-errorsChan:
	//	return err
	//default:
	//	return nil
	//}

	//close(errorsChan) // This is also perfectly valid instead of deferring
}

func main() {
	// Maps
	// Key value store { a: 1, b: 2 } O(1) get(a) > 1
	// [1,2,3,4,5,6] O(n)
	m := make(map[string]int)
	_, exists := m["a"]
	if _, ok := m["a"]; ok {
		log.Println("a")
	}
	if !exists {
		log.Println("no a")
	} else {
		log.Println("a")
	}
	delete(m, "a")
	clear(m)
	maps.Clone(m)
	maps.Equal(m, m)

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIdKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT1", cargo: 0},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Printf("All trucks processed succesfully!")

}
