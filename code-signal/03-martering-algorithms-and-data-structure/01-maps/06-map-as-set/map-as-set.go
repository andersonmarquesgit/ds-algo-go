package main

import "fmt"

func main() {
	uniqueVisitors := make(map[string]struct{})

	// Simulated visits to the website
	uniqueVisitors["visitor1@example.com"] = struct{}{}
	uniqueVisitors["visitor2@example.com"] = struct{}{}
	uniqueVisitors["visitor1@example.com"] = struct{}{}

	// TODO: Check presence of a visitor and print the result
	if _, visited := uniqueVisitors["visitor3@example.com"]; visited {
		fmt.Println("visitor3@example.com is visited")
	} else {
		fmt.Println("visitor3@example.com no visited")
	}
}
