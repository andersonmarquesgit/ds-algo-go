package main

import "fmt"

// Examplo do livro "Entendendo Algoritmos"
type Graph struct {
	adjacencyList map[string][]string // Map for the index string and the value is a slice of string
	numberOfNodes int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
		numberOfNodes: 0,
	}
}

func (g *Graph) AddVertex(node string) {
	if _, ok := g.adjacencyList[node]; !ok {
		g.adjacencyList[node] = []string{}
		g.numberOfNodes++
	}
}

func (g *Graph) AddEdge(node1, node2 string) {
	if _, ok := g.adjacencyList[node1]; ok {
		g.adjacencyList[node1] = append(g.adjacencyList[node1], node2)
	}

	if _, ok := g.adjacencyList[node2]; ok {
		g.adjacencyList[node2] = append(g.adjacencyList[node2], node1)
	}
}

func (g *Graph) ShowConnections() map[string][]string {
	return g.adjacencyList
}

func main() {
	newGraph := NewGraph()
	newGraph.AddVertex("voce")
	newGraph.AddVertex("alice")
	newGraph.AddVertex("bob")
	newGraph.AddVertex("claire")
	newGraph.AddVertex("anuj")
	newGraph.AddVertex("peggy")
	newGraph.AddVertex("jonny")
	newGraph.AddVertex("thom")

	newGraph.AddEdge("voce", "alice")
	newGraph.AddEdge("voce", "bob")
	newGraph.AddEdge("voce", "claire")
	newGraph.AddEdge("bob", "anuj")
	newGraph.AddEdge("bob", "peggy")
	newGraph.AddEdge("alice", "peggy")
	newGraph.AddEdge("claire", "jonny")
	newGraph.AddEdge("claire", "thom")

	fmt.Println(newGraph.ShowConnections())
	SearchBFSVendedorDeManga(*newGraph)
}

func SearchBFSVendedorDeManga(graph Graph) string {
	var searchQueue []string
	searchQueue = append(searchQueue, graph.adjacencyList["voce"]...)
	verified := make(map[string]bool)

	for len(searchQueue) > 0 {
		fmt.Println("Fila de pesquisa:", searchQueue)
		fmt.Println("Verificados:", verified)
		person := searchQueue[0] // Get the first person from the queue
		searchQueue = searchQueue[1:]
		if _, ok := verified[person]; !ok {
			if personIsSeller(person) {
				fmt.Printf("%s é um vendedor de manga!\n", person)
				return person
			} else {
				searchQueue = append(searchQueue, graph.adjacencyList[person]...)
				verified[person] = true
			}
		}
	}
	return ""
}

func personIsSeller(name string) bool {
	return name[len(name)-1:] == "m"
}
