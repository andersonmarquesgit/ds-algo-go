package main

import "fmt"

/*
Retorne o elemento com maior número de repetições
Por exemplo: [1, 2, 2, 3, 4, 5] = 2
Caso seja mais de um número, retorne qualquer um deles
Por exemplo: [1, 2, 2, 3, 3, 5] = 2 ou 3
[] = nil
*/

func countMaxFrequency(arr []int) int { //O(n^2)
	maxFrequency := 0
	element := -1

	for i := 0; i < len(arr); i++ {
		localFrequency := 0

		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				localFrequency++
			}
		}

		if localFrequency > maxFrequency {
			maxFrequency = localFrequency
			element = arr[i]
		}
	}

	return element
}

func countMaxFrequencyOptimization(arr []int) int {
	maxFrequency := 0
	element := -1
	frequencies := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		frequencies[arr[i]]++
	}

	for elem, frequency := range frequencies {
		if frequency > maxFrequency {
			maxFrequency = frequency
			element = elem
		}
	}

	return element
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5}
	fmt.Println(countMaxFrequency(arr))

	arr2 := []int{1, 2, 2, 3, 4, 5}
	fmt.Println(countMaxFrequencyOptimization(arr2))

	arr3 := []int{1, 2, 2, 3, 3, 5}
	fmt.Println(countMaxFrequency(arr3))

	arr4 := []int{1, 2, 2, 3, 3, 5}
	fmt.Println(countMaxFrequencyOptimization(arr4))
}
