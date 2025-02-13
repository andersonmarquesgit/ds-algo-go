package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()

	if msg != "Hello, universe!" {
		t.Errorf("Expected 'Hello, universe!', but got %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	outPut := string(result)

	os.Stdout = stdOut

	if !strings.Contains(outPut, "Hello, universe!") {
		t.Error("Expected 'Hello, universe!'")
	}

	if !strings.Contains(outPut, "Hello, cosmos!") {
		t.Error("Expected 'Hello, cosmos!'")
	}

	if !strings.Contains(outPut, "Hello, world!") {
		t.Error("Expected 'Hello, world!'")
	}
}
