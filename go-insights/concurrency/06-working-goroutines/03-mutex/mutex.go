package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
