package main

import "fmt"

func addTo80(n int) int {
	fmt.Println("Sempre executo o cálculo!")
	return n + 80
}

var cache = map[int]int{}

func memoizedAddTo80(n int) int {

	// Se o valor já foi calculado, retornamos o valor do cache
	if _, ok := cache[n]; ok {
		return cache[n]
	} else {
		// Se não, calculamos o valor e armazenamos no cache
		fmt.Println("Executo uma vez o cálculo!")
		cache[n] = n + 80
		return cache[n]
	}
}

// Outra forma de fazer sem usar variável global
func memoizedAddTo80V2() func(int) int {
	cache := map[int]int{}

	return func(n int) int {
		if _, ok := cache[n]; ok {
			return cache[n]
		} else {
			fmt.Println("Executo uma vez o cálculo!")
			cache[n] = n + 80
			return cache[n]
		}
	}
}

func main() {
	fmt.Println(addTo80(5))
	fmt.Println(addTo80(5))
	fmt.Println(addTo80(5))

	fmt.Println("")
	fmt.Println(memoizedAddTo80(5))
	fmt.Println(memoizedAddTo80(5))
	fmt.Println(memoizedAddTo80(5))

	fmt.Println("")
	memoizedAddTo80V2 := memoizedAddTo80V2()
	fmt.Println(memoizedAddTo80V2(5))
	fmt.Println(memoizedAddTo80V2(5))
	fmt.Println(memoizedAddTo80V2(5))

}
