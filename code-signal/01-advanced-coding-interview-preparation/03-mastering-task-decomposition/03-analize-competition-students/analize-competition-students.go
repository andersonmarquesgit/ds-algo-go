package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AnalyzeCompetition(logs string) [][]int {
	// TODO: we need split the logs for the logsList slice
	logsList := strings.Split(logs, ", ")

	// TODO: we need a students metrics with map of map
	// TODO: using a loop for get parts on logsList
	// 1: {"solve": 1, "fail": 1, "total-solve": 50}
	studentsMetrics := make(map[int]map[string]int)

	// TODO: a loop in logsList
	for _, log := range logsList {
		studentDetail := make(map[string]int)
		parts := strings.Split(log, " ")
		studentID, _ := strconv.Atoi(parts[0])
		score := parts[1]
		totalSolve := 0

		if score == "solve" {
			totalSolve, _ = strconv.Atoi(parts[2])
		}

		if _, exist := studentsMetrics[studentID]; exist {
			studentsMetrics[studentID][score]++
			studentsMetrics[studentID]["total-solve"] += totalSolve
		} else {
			studentDetail[score]++
			studentDetail["total-solve"] += totalSolve
			studentsMetrics[studentID] = studentDetail
		}
	}

	// TODO: We need a loop metric for record the student metrics
	result := [][]int{}
	for studentID, details := range studentsMetrics {
		totalSolve := details["total-solve"]
		qttSolve := details["solve"]
		qttFail := details["fail"]
		metric := []int{studentID, totalSolve, qttSolve, qttFail}
		result = append(result, metric)
	}

	sort.Slice(result, func(i, j int) bool {
		totalSolveI := result[i][1]
		totalSolveJ := result[j][1]
		return totalSolveI > totalSolveJ
	})

	return result
}

func main() {
	logs := "1 solve 50, 2 solve 60, 1 fail, 3 solve 40, 2 fail, 3 solve 70"

	// Expected output: [[3, 110, 2, 0], [2, 60, 1, 1], [1, 50, 1, 1]]
	result := AnalyzeCompetition(logs)
	for _, entry := range result {
		fmt.Printf("[%d, %d, %d, %d]\n", entry[0], entry[1], entry[2], entry[3])
	}
}
