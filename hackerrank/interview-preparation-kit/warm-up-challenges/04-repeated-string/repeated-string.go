package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

// Abordagem matemárica
/*
1. **Solução Matemática:**
    - Tempo: O(len(s))
    - Memória: O(1)
    - Executa instantaneamente

*/
func repeatedString(s string, n int64) int64 {
	if len(s) == 0 {
		return 0
	}

	// Conta quantos 'a' tem em uma repetição completa
	asEmUmaRepeticao := int64(strings.Count(s, "a"))

	// Calcula quantas repetições completas da string cabem em n
	repeticoesCompletas := n / int64(len(s))

	// Calcula quantos caracteres sobram
	resto := n % int64(len(s))

	// Calcula o total de 'a's
	total := asEmUmaRepeticao * repeticoesCompletas

	// Adiciona os 'a's do resto
	if resto > 0 {
		total += int64(strings.Count(s[:resto], "a"))
	}

	return total

}

/*
1. **Solução com StringBuilder:**
    - Tempo: O(n)
    - Memória: O(n)
    - Pode demorar alguns segundos

*/
// Abordagem bastante lenta usando de forma literal a string
func repeatedString2(s string, n int64) int64 {
	// Verifica se a string está vazia
	if len(s) == 0 {
		return 0
	}

	// Cria uma string builder para construir a string repetida
	var builder strings.Builder

	// Calcula quantas vezes precisamos repetir a string completa
	repeticoes := n / int64(len(s))
	resto := n % int64(len(s))

	// Adiciona as repetições completas
	for i := int64(0); i < repeticoes; i++ {
		builder.WriteString(s)
	}

	// Adiciona o resto da string
	if resto > 0 {
		builder.WriteString(s[:resto])
	}

	// Conta quantos 'a's tem na string final
	resultado := strings.Count(builder.String(), "a")

	return int64(resultado)
}

func main() {
	s := "aba"
	n := 10

	fmt.Println(repeatedString(s, int64(n)))
}
