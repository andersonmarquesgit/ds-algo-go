package main

import (
	"fmt"
)

func main() {
	// TODO: Create a slice for coffee shop orders with initial orders 'Latte', 'Espresso', 'Cappuccino'
	orders := []string{"Latte", "Expresso", "Cappuccino"}

	// TODO: A new customer orders a 'Mocha', add it to the end of the slice
	orders = append(orders, "Mocha")

	// TODO: Move the 'Latte' order to the end of the slice as the customer changed their preference.
	preference := orders[0]
	orders = orders[1:] // Essa operação com slice nos custa O(n), em Go existe uma forma otimizada de fazer isso usando deque "container/list"
	orders = append(orders, preference)
	fmt.Println(orders)

	// TODO: Process (remove and print) the first order in the slice. Which drink is it?
	orders = orders[1:]
	fmt.Println(orders[0])

}
