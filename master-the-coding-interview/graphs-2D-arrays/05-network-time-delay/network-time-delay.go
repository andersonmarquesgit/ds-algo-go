package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node int
	dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

/*
Time: O(E log V)
Space: O(V + E)
Usa Dijkstra com min-heap
*/
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]Edge, n)

	for _, t := range times {
		u, v, w := t[0]-1, t[1]-1, t[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k-1] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{k - 1, 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Item)

		if curr.dist > dist[curr.node] {
			continue
		}

		for _, edge := range graph[curr.node] {
			newDist := curr.dist + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(pq, Item{edge.to, newDist})
			}
		}
	}

	maxTime := 0
	for _, d := range dist {
		if d == math.MaxInt32 {
			return -1
		}
		if d > maxTime {
			maxTime = d
		}
	}

	return maxTime
}

func networkDelayTimeWithoutPriorityQueue(times [][]int, n int, k int) int {
	const INF = math.MaxInt32
	dist := make([]int, n)

	for i := range dist {
		dist[i] = INF
	}
	dist[k-1] = 0

	for i := 0; i < n-1; i++ {
		updated := false
		for _, t := range times {
			u, v, w := t[0]-1, t[1]-1, t[2]
			if dist[u] != INF && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	maxTime := 0
	for _, d := range dist {
		if d == INF {
			return -1
		}
		if d > maxTime {
			maxTime = d
		}
	}

	return maxTime
}

func main() {
	times := [][]int{
		{1, 2, 9},
		{1, 4, 2},
		{2, 5, 1},
		{4, 2, 4},
		{4, 5, 6},
		{3, 2, 3},
		{5, 3, 7},
		{3, 1, 5},
	}
	n := 5
	k := 1

	fmt.Println(networkDelayTimeWithoutPriorityQueue(times, n, k))
	fmt.Println(networkDelayTime(times, n, k))
}
