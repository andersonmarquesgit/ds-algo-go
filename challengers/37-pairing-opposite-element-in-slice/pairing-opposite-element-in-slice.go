package main

import (
	"fmt"
	"math"
)

type Pair struct {
	First  int
	Second float64
}

func solution(numbers []int) []Pair {
	// TODO: Implement the function to compute a new array of pairs, where each pair consists of an element from
	// the input array paired with its geometrical mean with the 'opposite' element. Remember to calculate
	// the geometrical mean to two decimal places and handle edge cases like odd-length arrays.
	result := []Pair{}
	n := len(numbers)

	for i := 0; i < n; i++ {
		newPair := Pair{
			First:  numbers[i],
			Second: math.Round(math.Sqrt(float64(numbers[i])*float64(numbers[n-i-1]))*100) / 100,
		}
		result = append(result, newPair)
	}

	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(solution(numbers))
}
