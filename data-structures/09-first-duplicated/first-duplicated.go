package main

//Google Question
// Given an array = [2,5,1,2,3,5,1,2,4]:
// It should return 2

// Given an array = [2,1,1,2,3,5,1,2,4]: // O primeiro algoritmo não consegue resolver esse caso pois no segungo for ele não consegue encontrar o primeiro duplicado
// It should return 1

// Given an array = [2,3,4,5]:
// It should return undefined

func firstDuplicated(nums []int) int {
	for i := 0; i < len(nums); i++ { // O(n^2)
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

func firstDuplicated2(nums []int) int {
	set := make(map[int]bool)
	// Given an array = [2,5,1,2,3,5,1,2,4]:
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return v
		}
		set[v] = true
	}
	return -1
}

func main() {
	nums := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	println(firstDuplicated(nums))  // 2
	println(firstDuplicated2(nums)) // 2

	nums2 := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	println(firstDuplicated(nums2))  // 1
	println(firstDuplicated2(nums2)) // 1

	nums3 := []int{2, 3, 4, 5}
	println(firstDuplicated(nums3))  // -1
	println(firstDuplicated2(nums3)) // -1
}
