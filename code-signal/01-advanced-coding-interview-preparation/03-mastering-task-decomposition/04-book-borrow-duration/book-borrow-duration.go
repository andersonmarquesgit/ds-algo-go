package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solution(logs string) []string {
	// TODO: Split the logs in logsList slice
	logsList := strings.Split(logs, ", ")

	// TODO: Get parts of each logsList [0]bookId, [1]borrowed/return [2]time
	// TODO: Calculate duration borrowed, we need two maps, first control borrow and return (we remove if return). Other map for the store duration for book id
	bookBorrowAndReturn := make(map[int]map[string]int)
	bookBorrowedDuration := make(map[int]int)

	// TODO: We need convert the time in minutes using split for calculate. Ex. 09:00 -> 09 00 -> 9 * 60 + 0 = 540 minutes
	for _, log := range logsList {
		parts := strings.Split(log, " ")
		bookID, _ := strconv.Atoi(parts[0])
		action := parts[1]

		time := strings.Split(parts[2], ":")
		hour, _ := strconv.Atoi(time[0])
		minutes, _ := strconv.Atoi(time[1])
		timestamp := hour*60 + minutes

		if _, exists := bookBorrowAndReturn[bookID]; exists {
			if action == "return" {
				duration := timestamp - bookBorrowAndReturn[bookID]["borrow"]
				bookBorrowedDuration[bookID] = duration
				delete(bookBorrowAndReturn, bookID)
			}
		} else {
			if action == "borrow" {
				detailBorrow := make(map[string]int)
				detailBorrow[action] = timestamp
				bookBorrowAndReturn[bookID] = detailBorrow
			}
		}

	}

	// TODO: We need revert the time for format hh:mm. For example 540 minutes -> ()540/60)%0 -> 09:00
	result := []string{}

	for bookID, duration := range bookBorrowedDuration {
		hours := duration / 60
		minutes := duration % 60
		durationString := fmt.Sprintf("%02d:%02d", hours, minutes)
		durationDetail := fmt.Sprintf("%d %s", bookID, durationString)
		result = append(result, durationDetail)
	}

	// TODO: We need a result slice of string and sort then
	sort.Slice(result, func(i, j int) bool {
		IDi, _ := strconv.Atoi(strings.Split(result[i], " ")[0])
		IDj, _ := strconv.Atoi(strings.Split(result[j], " ")[0])
		return IDi < IDj
	})

	return result
}

func main() {
	logs := "1 borrow 09:00, 2 borrow 10:00, 1 return 12:00, 3 borrow 13:00, 2 return 15:00, 3 return 16:00"
	result := Solution(logs)
	for _, entry := range result {
		fmt.Println("Book " + strings.Split(entry, " ")[0] + " duration: " + strings.Split(entry, " ")[1])
	}
	// Expected output: "Book 2 duration: 05:00"
}
