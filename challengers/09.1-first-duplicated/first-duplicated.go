package main

import "fmt"

func firstDuplicated(nums []int) {
	duplicated := make(map[int]bool)
	for numIndex := 0; numIndex < len(nums); numIndex++ {
		if _, ok := duplicated[nums[numIndex]]; ok {
			fmt.Println(nums[numIndex])
			return
		}

		duplicated[nums[numIndex]] = true
	}

	fmt.Println("Its should return undefined!")
}

func main() {
	nums := []int{2, 3, 1, 2, 3, 5, 1, 2, 4}
	firstDuplicated(nums)

	nums2 := []int{2, 3, 4, 5}
	firstDuplicated(nums2)
}
