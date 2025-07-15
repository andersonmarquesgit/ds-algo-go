package main

import (
	"fmt"
	"strings"
)

/*
The provided Go code is meant to list all unoccupied apartments in an apartment building. Each apartment has a number
and an occupancy status; however, the code doesn't work correctly at the moment. Your task is to fix the code so
that it functions as intended.
*/

func ListUnoccupiedApartments(apartments [][]string) {
	rows := len(apartments)

	for row := 0; row < rows; row++ {
		for col := 0; col < len(apartments[row]); col++ {
			parts := strings.Split(apartments[row][col], " ")
			number := parts[0]
			occupancy := parts[1] == "true"

			if !occupancy {
				fmt.Println("Apartment " + number + " is not unoccupied.")
			}
		}
	}
}

func main() {
	building := [][]string{
		{"101 true", "102 false", "103 false"},
		{"201 true", "202 true", "203 false", "204 false"},
	}

	ListUnoccupiedApartments(building)
}
