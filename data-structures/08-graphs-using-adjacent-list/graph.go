package main

import "fmt"

type Graph struct {
	adjacencyList map[int][]int // Map for the index int and the value is a slice of int
	numberOfNodes int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
		numberOfNodes: 0,
	}
}

func (g *Graph) AddVertex(node int) {
	if _, ok := g.adjacencyList[node]; !ok {
		g.adjacencyList[node] = []int{}
		g.numberOfNodes++
	}
}

func (g *Graph) AddEdge(node1, node2 int) {
	if _, ok := g.adjacencyList[node1]; ok {
		g.adjacencyList[node1] = append(g.adjacencyList[node1], node2)
	}

	if _, ok := g.adjacencyList[node2]; ok {
		g.adjacencyList[node2] = append(g.adjacencyList[node2], node1)
	}
}

func (g *Graph) ShowConnections() map[int][]int {
	return g.adjacencyList
}

func main() {
	graph := NewGraph()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)

	graph.AddEdge(3, 1)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 2)
	graph.AddEdge(4, 5)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 0)
	graph.AddEdge(0, 2)
	graph.AddEdge(6, 5)

	fmt.Println(graph.ShowConnections())
}
