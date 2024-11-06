package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	println(containsDuplicate(nums)) // true

	nums2 := []int{1, 2, 3, 4}
	println(containsDuplicate(nums2)) // false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	println(containsDuplicate(nums3)) // true
}
