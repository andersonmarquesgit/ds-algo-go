package main

func AreDisjoint(arr1, arr2 []int) bool { // O(n * m)
	for _, num1 := range arr1 {
		for _, num2 := range arr2 {
			if num1 == num2 {
				return false // An overlap is found.
			}
		}
	}
	return true // No overlaps found, sets are disjoint.
}

func AreDisjointOptimized(arr1, arr2 []int) bool { // O(n + m)
	set1 := make(map[int]struct{})
	for _, num := range arr1 {
		set1[num] = struct{}{} // Populate the map using struct{} to save space
	}

	for _, num := range arr2 {
		if _, exists := set1[num]; exists {
			return false // If found, the sets are not disjoint.
		}
	}
	return true
}
