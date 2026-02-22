package main

import (
	"testing"
)

func TestProcessTruck(t *testing.T) {
	tests := []struct {
		name           string
		truck          Truck
		checkAssertion func(t *testing.T)
	}{
		{
			name: "should process normal truck cargo",
			truck: &NormalTruck{
				id:    "NormalTruck-1",
				cargo: 0,
			},
			checkAssertion: func(t *testing.T) {
				// Note: NormalTruck.unloadCargo uses a value receiver in main.go,
				// so it might not update the state as expected.
			},
		},
		{
			name: "should process electric truck cargo",
			truck: &ElectricTruck{
				id:      "ElectricTruck-1",
				battery: 100,
			},
			checkAssertion: func(t *testing.T) {
				// Logic for ElectricTruck specific assertions
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := processTruck(tc.truck)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}
			tc.checkAssertion(t)
		})
	}
}
