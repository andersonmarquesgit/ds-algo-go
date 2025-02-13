package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func func1(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}

func func2(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(time.Millisecond * 30)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start")
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func1(&wg)
	go func2(&wg)

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	wg.Wait()

	//fmt.Scanln()
	fmt.Println("Done")
}
