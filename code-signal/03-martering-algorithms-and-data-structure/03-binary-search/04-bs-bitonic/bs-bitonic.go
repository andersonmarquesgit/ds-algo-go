package main

import (
	"fmt"
)

func FindIndex(arr []int, x int) int {
	// 1. Encontramos o pico
	peak := FindPeak(arr)

	// 2. Buscamos na parte crescente (esquerda do pico)
	result := BinarySearch(arr, x, 0, peak, true)
	if result != -1 {
		return result
	}

	// 3. Buscamos na parte decrescente (direita do pico)
	return BinarySearch(arr, x, peak+1, len(arr)-1, false)
}

func FindPeak(arr []int) int {
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[mid+1] {
			// Se `arr[mid]` é maior que `arr[mid+1]`, o pico está na esquerda (ou no meio)
			high = mid
		} else {
			// Caso contrário, está na direita
			low = mid + 1
		}
	}
	return low
}

func BinarySearch(arr []int, x, low, high int, ascending bool) int {
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		}

		if ascending {
			// Busca normal (array crescente)
			if arr[mid] < x {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// Busca invertida (array decrescente)
			if arr[mid] > x {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1 // Elemento não encontrado
}

func main() {
	arr := []int{-3, -2, 4, 6, 10, 8, 7, 1}
	x := 7
	position := FindIndex(arr, x)
	if position == -1 {
		fmt.Println("Element Not Present")
	} else {
		fmt.Printf("Element Present at Index %d\n", position)
	}
}
