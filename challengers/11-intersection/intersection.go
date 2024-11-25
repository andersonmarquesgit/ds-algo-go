package main

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	println(intersection(nums1, nums2)) // [2]

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	println(intersection(nums3, nums4)) // [9, 4]
}
