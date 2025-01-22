package main

import (
	"fmt"
	"runtime"
	"time"
)

func recFactorial(x int) int {
	if x <= 1 {
		return 1
	} else {
		return x * recFactorial(x-1)
	}
}

func tailFactorial(x int) int {
	totalSoFar := 1
	return factorial(x, totalSoFar)
}

func factorial(x, totalSoFar int) int {
	if x == 0 {
		return totalSoFar
	} else {
		return factorial(x-1, totalSoFar*x)
	}
}

func measureTimeAndMemory(f func(int) int, x int) {
	// Medindo o tempo
	start := time.Now()
	result := f(x)
	duration := time.Since(start)

	// Medindo o uso da memória
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Resultado: %d, Tempo: %v, Memória Alocada: %d bytes\n",
		result, duration, memStats.Alloc)
}

func main() {
	fmt.Println("Recursivo:")
	measureTimeAndMemory(recFactorial, 10)

	fmt.Println("Tail Recursivo:")
	measureTimeAndMemory(tailFactorial, 10)
}
