package main

import (
	"fmt"
)

func Solution(array1 []string, array2 []string) []bool {
	// TODO: First we can use a map for convert one this arrays
	set := make(map[string]bool)

	// TODO: We can store key celestial body using this structure
	for _, cb2 := range array2 {
		set[cb2] = true
	}

	// TODO: We need a commom slice []bool starting with all position false
	result := make([]bool, len(array1))

	// TODO: We can iterate over array1 for verify celestial body commons
	for i, value := range array1 {
		_, exists := set[value]
		result[i] = exists
	}

	return result
}

func main() {
	array1 := []string{"mars", "jupiter", "venus", "earth"}
	array2 := []string{"earth", "mars", "neptune"}
	output := Solution(array1, array2)
	fmt.Println(output) // Should print: [true false false true]
}
