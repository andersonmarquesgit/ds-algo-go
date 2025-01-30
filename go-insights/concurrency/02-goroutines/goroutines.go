package main

import (
	"fmt"
	"time"
)

func talk(people, text string, qtt int) {
	for i := 0; i < qtt; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", people, text, i+1)
	}
}

func main() {
	//talk("Maria", "Fala comigo!", 3)                // Processo serial
	//talk("João", "Só posso falar depois de vc!", 1) // Processo serial

	//go talk("Maria", "Fala comigo!", 500)
	//go talk("João", "Só posso falar depois de vc!", 500)

	go talk("Maria", "Fala comigo!", 10)
	talk("João", "Só posso falar depois de vc!", 5)
	// A goroutine é apenas até o fim da execução do main, a forma de manter a thread aberta até a sua conclusão é com channel
}
