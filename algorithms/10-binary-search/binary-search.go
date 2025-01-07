package main

import "fmt"

func binarySearch(arr []int, target int) bool { // O(log n)
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		fmt.Println("Target:", target)
		fmt.Println("Arr:", arr[low:high+1])
		fmt.Println("Mid:", arr[mid])
		fmt.Println("Low:", arr[low])
		fmt.Println("High:", arr[high])
		fmt.Println("")

		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1 // Update low, so we can search the right side, in other words, half the array size - O(log n)
		} else {
			high = mid - 1 // Update high, so we can search the left side, in other words, half the array size - O(log n)
		}

	}

	return false
}

func binarySearchRecursion(arr []int, target, low, high int) bool { // O(log n)
	if low > high {                                                 // Base case
		return false
	}

	mid := low + (high-low)/2

	fmt.Println("Mid:", mid)
	fmt.Println("Low:", low)
	fmt.Println("High:", high)
	fmt.Println("")

	if arr[mid] == target { // Base case
		return true
	} else if arr[mid] < target {
		return binarySearchRecursion(arr, target, mid+1, high)
	} else {
		return binarySearchRecursion(arr, target, low, mid-1)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6
	fmt.Println("Exist:", binarySearch(arr, target))
	fmt.Println("")

	arr2 := []int{1, 2, 3, 4}
	target2 := 6
	fmt.Println("Exist:", binarySearch(arr2, target2))
	fmt.Println("")
	//
	//fmt.Println("Exist Recursion:", binarySearchRecursion(arr, target, 0, len(arr)-1))
	//fmt.Println("")
	//
	//fmt.Println("Exist Recursion:", binarySearchRecursion(arr2, target2, 0, len(arr2)-1))
	//fmt.Println("")
}
