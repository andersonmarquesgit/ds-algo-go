package main

import (
	"fmt"
	"reflect"
)

type KeyValuePair struct {
	key   interface{}
	value interface{}
}

type HashTable struct {
	size    int
	buckets [][]KeyValuePair
}

func NewHashTable(size int) *HashTable {
	buckets := make([][]KeyValuePair, size)
	return &HashTable{
		size:    size,
		buckets: buckets,
	}
}

// Função hash para gerar um índice para a chave
func (h *HashTable) hash(key interface{}) int {
	hash := 0

	switch k := key.(type) {
	case string:
		for i := 0; i < len(k); i++ {
			charCode := int(k[i])
			hash = (hash + charCode*i) % h.size
		}
	case int:
		hash = k % h.size
	default:
		panic(fmt.Sprintf("Tipo de chave não suportado: %v", reflect.TypeOf(key)))
	}

	return hash
}

// Método para definir um par chave-valor na tabela hash
func (h *HashTable) Set(key, value interface{}) {
	address := h.hash(key)
	currentBucket := &h.buckets[address]

	// Se não há bucket no índice, cria um bucket novo
	if *currentBucket == nil {
		*currentBucket = []KeyValuePair{}
	}

	// Verifica se a chave já existe, se sim, atualiza o valor
	for i, pair := range *currentBucket {
		if pair.key == key {
			(*currentBucket)[i].value = value
			return
		}
	}

	// Caso contrário, adiciona um novo par chave-valor
	*currentBucket = append(*currentBucket, KeyValuePair{key: key, value: value})
}

// Método para obter o valor associado a uma chave na tabela hash
func (h *HashTable) Get(key interface{}) interface{} {
	address := h.hash(key)
	currentBucket := h.buckets[address]

	// Procura a chave no bucket atual
	if currentBucket != nil {
		for _, pair := range currentBucket {
			if pair.key == key {
				return pair.value
			}
		}
	}
	return nil
}

func (h *HashTable) Contains(key interface{}) bool {
	address := h.hash(key)
	currentBucket := h.buckets[address]

	// Procura a chave no bucket atual
	if currentBucket != nil {
		for _, pair := range currentBucket {
			if pair.key == key {
				return true
			}
		}
	}
	return false
}

func main() {
	myHashTable := NewHashTable(50)
	myHashTable.Set("grapes", 10000)
	fmt.Println("Value for 'grapes':", myHashTable.Get("grapes"))

	myHashTable.Set("apples", 9)
	fmt.Println("Value for 'apples':", myHashTable.Get("apples"))

	// Testando colisões
	myHashTable.Set("pears", 500)
	fmt.Println("Value for 'pears':", myHashTable.Get("pears"))

	myHashTable.Set("orange", 1000)
	fmt.Println("Value for 'orange':", myHashTable.Get("orange"))
	fmt.Println("Contains 'orange'?", myHashTable.Contains("orange"))
}
