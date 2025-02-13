package main

import (
	"fmt"
	"sync"
)

func printSomething(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Finaliza o elemento do grupo
	fmt.Println(text)
}

func main() {
	//go printSomething("Primeiro texto!")
	//time.Sleep(1 * time.Second) // A pior forma de esperar a goroutine
	//printSomething("Segundo texto!")

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"theta",
		"lambda",
		"omega",
	}
	//wg.Add(7) // Geralmente não usamos assim pois podemos ter um número maior do que o WG e ter um deadlock
	wg.Add(len(words)) // 7 representa a quantidade de vezes que vamos fazer a execução e que precisamos esperar do grupo

	for i, text := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, text), &wg) // Precisamos ser cautelosos na passagem de WGs por isso usamos ponteiros (documentação do Go)
	}

	wg.Wait() // Aguardará até que o grupo seja ZERO. O lugar para fazer essa redução é na goroutine

	maisUmTexto := "Mais um texto!"
	wg.Add(1) // Logo precisamos adicionar no grupo
	printSomething(maisUmTexto, &wg)
	// Wait Group é a forma mais fácil de usar goroutines, mas não podem ficar negativos pois haverá erro
	// Várias goroutine em execução não há garantia de que serão executadas em ordem
}
