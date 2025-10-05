package main

import "errors"

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type TruckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() TruckManager {
	return TruckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *TruckManager) AddTruck(id string, cargo int) error {
	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *TruckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := tm.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}
func (tm *TruckManager) RemoveTruck(id string) error {
	_, ok := tm.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}
func (tm *TruckManager) UpdateTruckCargo(id string, cargo int) error {
	_, ok := tm.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}
	tm.trucks[id].Cargo = cargo
	return nil
}
func main() {

}
