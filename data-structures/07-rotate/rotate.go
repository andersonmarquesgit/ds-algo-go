package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n                // Se k for maior que o tamanho do array, ajustamos para evitar rotações desnecessárias
	reverse(nums, 0, n-1) // Reverte o array inteiro
	reverse(nums, 0, k-1) // Reverte os primeiros k elementos
	reverse(nums, k, n-1) // Reverte os elementos restantes
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // [5 6 7 1 2 3 4]

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2) // [3 99 -1 -100]
}
