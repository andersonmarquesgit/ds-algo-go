package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented yet")
	ErrNotFound       = errors.New("truck not found")
)

type Truck interface {
	loadCargo() error
	unloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) loadCargo() error {
	t.cargo += 1
	return nil
}

func (t NormalTruck) unloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *ElectricTruck) loadCargo() error {
	t.cargo += 1
	t.battery = -1

	return nil
}

func (t *ElectricTruck) unloadCargo() error {
	t.cargo = 0
	t.battery = -1

	return nil
}

func processTruck(truck Truck) error {
	fmt.Println("Processing truck", truck)
	err := truck.loadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.unloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nTruck := NormalTruck{
		id:    "NormalTruck-1",
		cargo: 0,
	}

	eTruck := ElectricTruck{
		id:      "ElectricTruck-1",
		battery: 100,
	}

	err := processTruck(&nTruck)
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	err = processTruck(&eTruck)
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}
}
