package main

import "fmt"

// Brute Force using explicit detection of cycles
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList := make([][]int, numCourses)

	for _, p := range prerequisites {
		from, to := p[1], p[0]
		adjacencyList[from] = append(adjacencyList[from], to)
	}

	for v := 0; v < numCourses; v++ {
		queue := []int{}
		seen := make(map[int]bool)

		for _, next := range adjacencyList[v] {
			queue = append(queue, next)
		}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current == v {
				return false
			}

			if seen[current] {
				continue
			}
			seen[current] = true

			for _, next := range adjacencyList[current] {
				queue = append(queue, next)
			}
		}

	}

	return true
}

// Topological Sort
func canFinishWithTopologicalSort(numCourses int, prerequisites [][]int) bool {
	adjacencyList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		from, to := p[1], p[0]
		adjacencyList[from] = append(adjacencyList[from], to)
		inDegree[to]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited++

		for _, next := range adjacencyList[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return visited == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))

	fmt.Println(canFinishWithTopologicalSort(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinishWithTopologicalSort(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
