package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrganizeInbox(inboxString string) []string {
	// TODO: Your solution
	emailList := strings.Split(inboxString, "; ")
	senderCount := make(map[string]int)
	result := []string{}

	for _, email := range emailList {
		parts := strings.Split(email, ", ")
		sender := parts[0]
		//subject := parts[1]
		//time := parts[2]
		senderCount[sender]++
	}

	for sender, qtt := range senderCount {
		result = append(result, fmt.Sprintf("%s %d", sender, qtt))
	}

	sort.Slice(result, func(i, j int) bool {
		senderCountsI, _ := strconv.Atoi(strings.Split(result[i], " ")[1])
		senderCountsJ, _ := strconv.Atoi(strings.Split(result[j], " ")[1])
		return senderCountsI > senderCountsJ
	})

	return result
}

func main() {
	inboxString := "sender1@example.com, Subject1, 09:00; sender2@example.com, Subject2, 10:00; sender1@example.com, Subject3, 12:00"

	// Expected output: ["sender1@example.com 2", "sender2@example.com 1"]
	result := OrganizeInbox(inboxString)
	for _, entry := range result {
		fmt.Println(entry)
	}
}
