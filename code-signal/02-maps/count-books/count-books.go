package main

import (
	"fmt"
)

func main() {
	// Create and initialize a slice of strings 'books' with some duplicated book titles of your choice
	books := []string{"Clean Architecture", "No Rules Rule", "Clean Architecture", "Clean Code", "Refactoring", "Refactoring", "Refactoring", "Refactoring", "System Design Interview", "System Design Interview", "Building Evolutionary Architecture", "No Rules Rule", "No Rules Rule"}

	// Create an empty map 'bookCount' to store the count of each book
	bookCount := map[string]int{}

	// Loop through each book in the 'books' slice and count the occurrences
	for _, book := range books {
		bookCount[book]++
	}

	// Finally, print the 'bookCount' map
	for book, qtt := range bookCount {
		fmt.Printf("%s: %d\n", book, qtt)
	}
}
