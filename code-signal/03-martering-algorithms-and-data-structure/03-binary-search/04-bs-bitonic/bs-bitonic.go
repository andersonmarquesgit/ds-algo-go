package main

import "fmt"

func BinarySearch(temperatures []int, low int, high int, targetTemp int, ascending bool) int {
	for low <= high {
		mid := low + (high-low)/2 // Calculate the middle index
		if temperatures[mid] == targetTemp {
			return mid // Target found at mid
		}
		// Check the direction of the binary search using the ascending flag
		if (ascending && temperatures[mid] < targetTemp) || (!ascending && temperatures[mid] > targetTemp) {
			low = mid + 1 // Move to the right half
		} else {
			high = mid - 1 // Move to the left half
		}
	}
	return -1 // Target not found in the array
}

func FindPeak(temperatures []int) int {
	low, high := 0, len(temperatures)-1 // Initialize the search range
	for low < high {
		mid := low + (high-low)/2 // Calculate the middle index
		// Check if the mid element is greater than the next element
		if temperatures[mid] > temperatures[mid+1] {
			high = mid // Peak is in the left half including mid
		} else {
			low = mid + 1 // Peak is in the right half excluding mid
		}
	}
	return low // This is the index of the peak temperature.
}

func SearchBitonicArray(temperatures []int, targetTemp int) int {
	if len(temperatures) == 0 {
		return -1 // Return -1 if the array has no elements
	}
	peakIndex := FindPeak(temperatures) // Find peak in the bitonic array

	// Search in the ascending part of the bitonic array
	searchResult := BinarySearch(temperatures, 0, peakIndex, targetTemp, true)
	if searchResult != -1 {
		return searchResult // Target found in ascending part
	} else {
		// Search in the descending part of the bitonic array
		return BinarySearch(temperatures, peakIndex+1, len(temperatures)-1, targetTemp, false)
	}
}

func main() {
	temperatures := []int{1, 3, 8, 12, 4, 2}
	targetTemp := 12

	fmt.Println(BinarySearch(temperatures, 0, len(temperatures)-1, targetTemp, true))
}
