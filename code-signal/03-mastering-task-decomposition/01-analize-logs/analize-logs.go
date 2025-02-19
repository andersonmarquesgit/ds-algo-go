/*
Task Statement
Imagine a social networking application that allows users to form groups.
Each group has a unique ID ranging from 1 to n, where n is the total number of groups. The app keeps track of group
creation and deletion events, logging all these actions in a string.

The task before us is to create a Go function named AnalyzeLogs. This function will take as input a string of logs
and output a slice of strings representing the groups with the longest lifetime. Each string in the slice contains
two items separated by a space: the group ID and the group's lifetime. By "lifetime," we mean the duration from when
the group was created until its deletion. If a group has been created and deleted multiple times, the lifetime is the
total sum of those durations.

For example, if we have a log string as follows: "1 create 09:00, 2 create 10:00, 1 delete 12:00, 3 create 13:00,
2 delete 15:00, 3 delete 16:00", the function will return: ["2 05:00"].

If multiple groups have the same longest lifetime, the function should return all such groups in ascending order
of their IDs. Let's say we have the log string "1 create 08:00, 2 create 09:00, 1 delete 10:00, 2 delete 11:00,
3 create 12:00, 3 delete 13:00". In this case, both group 1 and group 2 have the same lifetime of 02:00.
The function should return ["1 02:00", "2 02:00"] since both have the longest lifetime, and group 1 comes before group 2.

Step Overview
To tackle this problem, we will take the following steps:

Split the Log Strings: Divide the input string into individual log entries based on a delimiter using strings.Split.
Parse Log Components: For each log entry, identify the group ID, action type, and timestamp, and use strconv.Atoi for conversions.
Record and Calculate Lifetimes: Track creation times using maps and compute lifetimes whenever a group is deleted.
Identify Longest Lifetimes: Compare the lifetimes of all groups to determine the ones with the longest using Go's
iteration and sorting capabilities.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AnalyzeLogs(logs string) []string {
	logList := strings.Split(logs, ", ")

	timeDict := make(map[int][2]int) // Store the creation time in minutes
	lifeDict := make(map[int]int)    // Store the total lifetime for each group

	for _, log := range logList {
		parts := strings.Split(log, " ")
		groupId, _ := strconv.Atoi(parts[0])
		action := parts[1]
		time := parts[2]

		// Parsing the time from HH:MM format
		hour, _ := strconv.Atoi(time[:2])
		minute, _ := strconv.Atoi(time[3:])
		currentTime := hour*60 + minute // Time in minutes from the start of the day

		if action == "create" {
			timeDict[groupId] = [2]int{hour, minute} // If the group is created, log the creation time.
		} else {
			if creationTime, exists := timeDict[groupId]; exists {
				// If the group is deleted, calculate its lifetime and update the records.
				creationMinutes := creationTime[0]*60 + creationTime[1]
				lifetime := currentTime - creationMinutes
				lifeDict[groupId] += lifetime
				delete(timeDict, groupId) // Remove group from timeDict after calculating lifetime.
			}
		}
	}

	// Find the longest lifetime
	var maxLife int
	for _, life := range lifeDict {
		if life > maxLife {
			maxLife = life
		}
	}

	// Building the result list where each item is a string of "group ID lifetime" if it has the longest lifetime.
	var result []string
	for id, life := range lifeDict {
		if life == maxLife {
			hours := life / 60
			minutes := life % 60
			timeString := fmt.Sprintf("%02d:%02d", hours, minutes)
			result = append(result, fmt.Sprintf("%d %s", id, timeString))
		}
	}

	// Sorting the result in ascending order of the group IDs
	sort.Slice(result, func(i, j int) bool {
		id1, _ := strconv.Atoi(strings.Split(result[i], " ")[0])
		id2, _ := strconv.Atoi(strings.Split(result[j], " ")[0])
		return id1 < id2
	})

	return result
}

func main() {
	logs := "1 create 09:00, 2 create 10:00, 1 delete 12:00, 3 create 13:00, 2 delete 15:00, 3 delete 16:00"
	result := AnalyzeLogs(logs)
	for _, entry := range result {
		fmt.Println("Group " + strings.Split(entry, " ")[0] + " lifetime: " + strings.Split(entry, " ")[1])
	}
	// Outputs:
	// Group 2 lifetime: 05:00
}
