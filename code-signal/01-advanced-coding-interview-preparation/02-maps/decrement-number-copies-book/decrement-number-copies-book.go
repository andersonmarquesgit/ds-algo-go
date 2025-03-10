package main

import "fmt"

func main() {
	// Map to hold the library catalog system
	libraryCatalog := make(map[string]map[string]int)

	// Example: Adding a book
	bookInfo := make(map[string]int)
	bookInfo["copies"] = 3
	libraryCatalog["978-3-16-148410-0"] = bookInfo

	// Example: Borrowing a book (decrementing the number of copies)
	numberOfCopies := libraryCatalog["978-3-16-148410-0"]["copies"]
	if numberOfCopies > 0 {
		for i := 0; i < numberOfCopies; i++ {
			libraryCatalog["978-3-16-148410-0"]["copies"]--
		}
	}

	// Print the number of copies left
	fmt.Println("Copies left:", libraryCatalog["978-3-16-148410-0"]["copies"])
}
