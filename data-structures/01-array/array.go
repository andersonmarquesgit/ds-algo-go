package main

import "fmt"

type MyArray struct {
	len  int
	data []interface{}
}

func NewArray() *MyArray {
	return &MyArray{
		len:  0,
		data: []interface{}{},
	}
}

func (a *MyArray) Get(index int) interface{} { // O(1)
	if index < 0 || index >= a.len {
		return nil
	}
	return a.data[index]
}

func (a *MyArray) Push(item interface{}) { // O(n)
	// Se a capacidade do array não for suficiente para adicionar um novo elemento
	if a.len == len(a.data) {
		// Aumenta a capacidade do array apenas o suficiente para o novo elemento
		newCapacity := a.len + 1
		newData := make([]interface{}, newCapacity)
		copy(newData, a.data) //O(n)
		a.data = newData
	}

	// Adiciona o item e incrementa o tamanho do array
	a.data[a.len] = item
	a.len++
}

func (a *MyArray) Pop() interface{} { // O(1)
	if a.len == 0 {
		return nil // Nenhum elemento para remover
	}

	// Armazena o último elemento a ser removido
	lastItem := a.data[a.len-1]

	// Decrementa o tamanho do array
	a.len--

	// Redefine o slice para conter apenas os elementos válidos
	a.data = a.data[:a.len] // O(1)

	return lastItem
}

func (a *MyArray) Delete(index int) { // O(n)
	if index < 0 || index >= a.len {
		return
	}

	// Deletion o item no índice especificado
	for i := index; i < a.len-1; i++ { // O(n)
		a.data[2] = a.data[i+1]
	}
	a.len--

	// Redefine o slice para conter apenas os elementos válidos
	a.data = a.data[:a.len] // O(1)
}

func main() {
	myArray := NewArray()
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(1)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(2)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(3)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(4)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Pop()
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Pop()
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Pop()
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Pop()
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(1)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(2)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(3)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Push(4)
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)

	myArray.Delete(2) // Deletion o 3
	fmt.Println("Len: ", myArray.len, "Items: ", myArray.data)
}
