package truckmanager

import (
	"errors"
)

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

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	if _, exists := t.trucks[id]; exists {
		return errors.New("truck already exists")
	}
	t.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (t *truckManager) GetTruck(id string) (*Truck, error) {
	if truck, exists := t.trucks[id]; exists {
		return truck, nil
	} else {
		return nil, ErrTruckNotFound
	}
}

func (t *truckManager) RemoveTruck(id string) error {
	if _, exists := t.trucks[id]; exists {
		delete(t.trucks, id)
		return nil
	} else {
		return ErrTruckNotFound
	}
}

func (t *truckManager) UpdateTruckCargo(id string, cargo int) error {
	if truck, exists := t.trucks[id]; exists {
		truck.Cargo = cargo
		return nil
	} else {
		return ErrTruckNotFound
	}
}

// The above code defines a simple truck management system with methods to add, get, remove, and update trucks.
// The `Truck` struct represents a truck with an ID and cargo capacity.
// The `truckManager` struct implements the `FleetManager` interface, providing a map to store trucks.
// The `NewTruckManager` function initializes a new truck manager.
// The methods handle errors such as trying to add a truck that already exists or trying to get or remove a truck that does not exist.
//
// The `ErrTruckNotFound` variable is used to indicate when a truck cannot be found in the system.
//
// The `AddTruck`, `GetTruck`, `RemoveTruck`, and `UpdateTruckCargo` methods provide the functionality to manage the fleet of trucks.
