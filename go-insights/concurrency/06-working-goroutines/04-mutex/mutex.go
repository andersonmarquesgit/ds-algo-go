package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Income struct {
	Amount float64
	Source string
}

func main() {
	// variable for bank balance
	balance := 0.0

	// print out starting values
	fmt.Printf("Starting balance: $%.2f\n", balance)

	// define weekly revenue
	incomes := []Income{
		{Amount: 1000, Source: "Main job"},
		{Amount: 50, Source: "Side job"},
		{Amount: 200, Source: "Freelance project"},
		{Amount: 500, Source: "Investments"},
		{Amount: 10, Source: "Gift"},
	}

	// loop through 52 weeks and print out how much is made. Keep a running total
	for i, income := range incomes {
		wg.Add(1)
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mutex.Lock()
				balance += income.Amount
				fmt.Printf("Week %d - %s: $%.2f\n", week, income.Source, balance)
				mutex.Unlock()
			}
		}(i, income)
	}

	// print out the final balance
	wg.Wait()
	fmt.Printf("Final balance: $%.2f\n", balance)
}
