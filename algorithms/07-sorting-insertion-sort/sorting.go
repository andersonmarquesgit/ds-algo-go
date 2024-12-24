package main

import "fmt"

func insertionSort(nums []int) []int { //O(n^2)
	for i := 1; i < len(nums); i++ {
		j := i
		for j > 0 && nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}

	return nums
}

func main() {
	nums := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(insertionSort(nums)) // [1 2 3 4 5 6 7 8]

}
