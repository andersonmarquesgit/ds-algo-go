package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1) // Por se tratar de um teste apenas 1 é suficiente

	go printSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	outPut := string(result)

	os.Stdout = stdOut

	if !strings.Contains(outPut, "epsilon") {
		t.Errorf("Expected to find epsilon, but is not there")
	}
}
