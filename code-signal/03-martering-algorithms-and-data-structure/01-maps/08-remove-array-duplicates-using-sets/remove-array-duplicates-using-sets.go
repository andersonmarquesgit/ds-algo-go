package main

func RemoveDuplicates(arr []int) []int { // O(n^2)
	uniqueList := []int{}
	for _, num := range arr {
		contains := false
		for _, uniqueNum := range uniqueList {
			if num == uniqueNum {
				contains = true
				break
			}
		}
		if !contains {
			uniqueList = append(uniqueList, num) // Add number if it's not already in the list
		}
	}
	return uniqueList
}

func RemoveDuplicatesOptimized(arr []int) []int { // O(n)
	nums := make(map[int]struct{})
	for _, num := range arr {
		nums[num] = struct{}{} // Add the number if it's not already present, ignoring duplicates
	}

	result := make([]int, 0, len(nums))
	for num := range nums {
		result = append(result, num) // Each unique element is added to the result slice
	}
	return result
}
