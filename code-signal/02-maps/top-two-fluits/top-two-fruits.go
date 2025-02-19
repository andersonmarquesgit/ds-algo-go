package main

import (
	"fmt"
)

func main() {
	// TODO: Create a map representing the inventory with fruits as keys and their quantities as values
	inventory := map[string]int{
		"apples":   25,
		"bananas":  20,
		"cherries": 30,
	}

	// TODO: Find the counts of the top two most abundant fruits
	firstMax, secondMax := 0, 0
	for _, qtt := range inventory {
		if qtt > firstMax {
			secondMax = firstMax
			firstMax = qtt
		} else if qtt > secondMax {
			secondMax = qtt
		}
	}

	// TODO: Calculate the total count of the top two most abundant fruits
	totalCount := firstMax + secondMax

	// TODO: Print out the total count of the top two most abundant fruits in the inventory
	fmt.Printf("First Max is %d and Second Max is %d: Total is %d", firstMax, secondMax, totalCount)
}
