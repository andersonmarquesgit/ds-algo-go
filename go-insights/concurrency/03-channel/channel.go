package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1     // adicionando dados no canal
	num := <-ch // recebendo dados do canal

	ch <- 2

	fmt.Println(<-ch) // recebendo dados do canal
	fmt.Println(num)
}
