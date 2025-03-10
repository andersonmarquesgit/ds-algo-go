package main

import (
	"fmt"
)

func findPairSum(targetSum int, numbers []int) []int {
	// TODO: We have the condition that time complexity is O(n), so we dont use brutal force because is O(nˆ2)
	// usind two loops for{for{}}
	// TODO: So we can use a data structure map and complement method. This case we have a seach using key:complement that O(1)
	// For example {2, 13, 4, 7, 5, 15} and target is 9
	// => 9-2  ->  7:0;
	// => 9-13 -> -4:1;
	// => 9-4  ->  5:2;
	// => 9-7  ->  2:3;
	// => 9-5  ->  4:4;
	// => 9-15 -> -6:5;

	complements := make(map[int]int)

	for idx, num := range numbers {
		if _, exists := complements[num]; exists {
			return []int{numbers[complements[num]], numbers[idx]}
		}
		complement := targetSum - num
		complements[complement] = idx
	}

	return []int{}
}

func main() {
	numbers := []int{2, 13, 4, 7, 5, 15}
	target := 9
	result := findPairSum(target, numbers)
	fmt.Println(result)
}
