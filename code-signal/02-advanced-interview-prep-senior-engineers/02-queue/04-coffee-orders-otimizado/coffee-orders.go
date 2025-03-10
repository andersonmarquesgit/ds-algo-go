package main

import (
	"container/list"
	"fmt"
)

func main() {
	// TODO: Create a slice for coffee shop orders with initial orders 'Latte', 'Espresso', 'Cappuccino'
	orders := list.New()

	orders.PushBack("Latte")
	orders.PushBack("Expresso")
	orders.PushBack("Cappuccino")

	// TODO: A new customer orders a 'Mocha', add it to the end of the slice
	orders.PushBack("Mocha")

	// Exibir antes da modificação
	fmt.Print("Antes: ")
	for e := orders.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// TODO: Move the 'Latte' order to the end of the slice as the customer changed their preference.
	MoveOrderToEnd(orders)

	// TODO: Process (remove and print) the first order in the slice. Which drink is it?
	// Exibir após a modificação
	fmt.Print("Depois: ")
	for e := orders.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

}

// Move o primeiro pedido para o final da lista
func MoveOrderToEnd(orders *list.List) {
	if orders.Len() == 0 {
		return // Nada para mover se a lista estiver vazia
	}

	firstElement := orders.Front()  // O(1) -> Acessa o primeiro elemento
	orders.MoveToBack(firstElement) // O(1) -> Move para o final
}
