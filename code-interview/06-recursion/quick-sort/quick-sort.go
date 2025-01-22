package main

import "fmt"

func quicksort(nums []int) []int {
	pivotPosition := len(nums) - 1

	if pivotPosition > 0 { // Dá-nos certeza que os arrays da esq e direita tem pelo menos 1 elemento
		pivot := nums[pivotPosition] // Define o Pivot sempre como o elemento mais a direita
		/*
			O pivô é sempre o último elemento do array (nums[len(nums)-1]).
			Essa abordagem funciona, mas pode levar ao pior caso O(n^2) quando a entrada está ordenada ou quase ordenada,
			pois o pivô será sempre o menor ou maior elemento.

			Exemplo de escolha aleatória:
			pivotPosition := rand.Intn(len(nums))
		*/

		i, j := 0, 0
		// i => Rastreia onde o pivô deve ficar
		// j => Percorre o array e valida se o elemento em j é maior ou menor que o pivô, caso seja menor, incrementamos i

		for j <= pivotPosition {
			if nums[j] < pivot {
				swap(nums, i, j)
				i++
			}
			j++
		}

		swap(nums, i, pivotPosition)
		quicksort(nums[:i])   // left
		quicksort(nums[i+1:]) // right
	}

	return nums
}

func swap(nums []int, i, j int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}

func main() {
	nums := []int{5, 3, 1, 6, 4, 2}
	fmt.Println(quicksort(nums))
}
