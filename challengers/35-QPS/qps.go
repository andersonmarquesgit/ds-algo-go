package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalQueries int64
var startTime = time.Now()

func processQuery() {
	time.Sleep(100 * time.Millisecond) // Simulando latência de 100 milissegundos
	atomic.AddInt64(&totalQueries, 1)
}

func main() {
	for i := 0; i < 100; i++ {
		go processQuery()
	}

	time.Sleep(2 * time.Second) // Espera 2 segundos para que todas as goroutines terminem

	duration := time.Since(startTime)
	qps := float64(atomic.LoadInt64(&totalQueries)) / duration.Seconds()
	fmt.Printf("Total Queries: %d\n", totalQueries)
	fmt.Printf("QPS: %.2f\n", qps)
}
