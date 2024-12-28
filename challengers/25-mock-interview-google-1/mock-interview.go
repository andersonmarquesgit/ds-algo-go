package main

import "fmt"

// Brutal force: O(nˆ2)
func findBanks(banks []int, cost int) []int {
	var matchBanks []int

	for i := 0; i < len(banks); i++ {
		for j := 0; j < len(banks); j++ {
			if i != j && banks[i]+banks[j] == cost {
				matchBanks = append(matchBanks, i)
			}
		}
	}

	return matchBanks
}

// Optimized: O(n)
func findBanks2(banks []int, cost int) []int {
	bankMap := make(map[int]int)

	for bankIndex := 0; bankIndex < len(banks); bankIndex++ {
		bankComp := cost - banks[bankIndex] // Comp representa o valor que precisamos encontrar para completar o custo
		// Em vez de verificar todos os elementos do array, verificamos se o complemento já existe no map
		// E interrompe imediatamente a execução

		if _, ok := bankMap[bankComp]; ok {
			return []int{bankMap[bankComp], bankIndex}
		}

		bankMap[banks[bankIndex]] = bankIndex
	}

	return []int{}
}

func main() {
	banks := []int{1, 2, 3}
	fmt.Println(findBanks(banks, 4))

	fmt.Println(findBanks2(banks, 4))
}
