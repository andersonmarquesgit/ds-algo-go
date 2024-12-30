package main

import (
	"fmt"
	"reflect"
)

type HashTable struct {
	buckets map[int]int
}

func NewHashTable() *HashTable {
	return &HashTable{
		buckets: make(map[int]int),
	}
}

func (h *HashTable) hash(key interface{}) int {
	const prime = 7 // Constante primo menor para reduzir a magnitude do hash
	hash := 0

	switch k := key.(type) {
	case string:
		for i, char := range k {
			hash = (hash*prime + int(char) + i) & 0xFFFF // Limitar a 16 bits (0 a 65535)
		}
	case int:
		hash = (k * prime) & 0xFFFF // Limitar a 16 bits
	default:
		panic(fmt.Sprintf("Tipo de chave não suportado: %v", reflect.TypeOf(key)))
	}

	return hash
}

func (h *HashTable) Insertion(key interface{}, value int) { // O(1)
	address := h.hash(key)

	h.buckets[address] = value
}

func (h *HashTable) Search(key interface{}) interface{} { // O(1)
	address := h.hash(key)
	currentBucket, exists := h.buckets[address]

	if !exists {
		return nil
	}

	return currentBucket
}

func (h *HashTable) Contains(key interface{}) bool { // O(1)
	address := h.hash(key)
	_, exists := h.buckets[address]

	return exists
}

func (h *HashTable) Deletion(key interface{}) { // O(1)
	address := h.hash(key)
	delete(h.buckets, address)
}

func main() {
	myHashTable := NewHashTable()
	myHashTable.Insertion("grapes", 10000)
	fmt.Println("Value for 'grapes':", myHashTable.Search("grapes"))
	fmt.Println("")

	myHashTable.Insertion("apples", 9)
	fmt.Println("Value for 'apples':", myHashTable.Search("apples"))
	fmt.Println("")

	fmt.Println("Keys:", myHashTable.buckets)
	fmt.Println("Size:", len(myHashTable.buckets))
	fmt.Println("")

	myHashTable.Insertion("apples", 20000)
	fmt.Println("Value for 'apples':", myHashTable.Search("apples"))
	fmt.Println("")

	myHashTable.Insertion("pears", 500)
	fmt.Println("Value for 'pears':", myHashTable.Search("pears"))
	fmt.Println("")

	myHashTable.Insertion("orange", 1000)
	fmt.Println("Value for 'orange':", myHashTable.Search("orange"))
	fmt.Println("Contains 'orange'?", myHashTable.Contains("orange"))
	fmt.Println("")

	myHashTable.Insertion("lemon", 5000)
	fmt.Println("Value for 'lemon':", myHashTable.Search("lemon"))
	fmt.Println("Keys:", myHashTable.buckets)
	fmt.Println("Size:", len(myHashTable.buckets))
	fmt.Println("")
	myHashTable.Deletion("lemon")
	fmt.Println("Contains 'lemon'?", myHashTable.Contains("lemon"))
	fmt.Println("")

	fmt.Println("Keys:", myHashTable.buckets)
	fmt.Println("Size:", len(myHashTable.buckets))
}
