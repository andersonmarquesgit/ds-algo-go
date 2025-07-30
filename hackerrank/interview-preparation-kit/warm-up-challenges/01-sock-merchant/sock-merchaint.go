package main

import "fmt"

// sockMerchant conta o número de pares de meias de mesma cor em um array.
// Cada meia é representada por um número inteiro que indica sua cor.
// Retorna o número total de pares encontrados.
func sockMerchant(n int32, ar []int32) int32 {
	colorCount := make(map[int32]int)
	var pairs int32

	for _, color := range ar {
		colorCount[color]++
		if colorCount[color] == 2 {
			pairs++
			colorCount[color] = 0
		}
	}

	return pairs

}

func main() {
	n := 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	fmt.Println(sockMerchant(int32(n), ar))
}

/*
Para o problema do (contagem de pares de meias), aqui está uma boa abordagem de perguntas iniciais para uma entrevista de código: `sockMerchant`
1. **Esclarecer o Problema:**
    - "Posso assumir que o array de entrada nunca será nulo?"
    - "Como devo tratar um array vazio?"
    - "Existe algum limite para os valores que representam as cores?"
    - "Uma meia pode ser usada em mais de um par?"

2. **Explorar Casos de Uso:**
    - "Vamos revisar alguns exemplos juntos para garantir que entendi corretamente?"
    - "Como devemos tratar o caso onde temos 3 meias da mesma cor? Conta como 1 par ou nenhum?"
    - "Se tivermos 4 meias da mesma cor, isso conta como 2 pares?"

3. **Restrições:**
    - "Existe alguma restrição de memória que eu deva considerar?"
    - "Precisamos otimizar para algum caso específico?"
    - "Qual é o tamanho máximo do array de entrada?"

   // Propor casos de teste como:
   ar := []int32{} // array vazio
   ar := []int32{1, 1, 1} // 3 meias iguais
   ar := []int32{1, 2, 3} // nenhum par
   ar := []int32{1, 1, 1, 1} // 2 pares
   ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20} // caso complexo

1. **Abordagem Inicial:**
    - "Antes de começar a codificar, posso explicar minha abordagem?"
    - "Estou pensando em usar um mapa para contar as ocorrências. O que você acha?"
    - "Vejo outras possíveis soluções, como ordenar o array primeiro. Gostaria de discutir os prós e contras?"

2. **Complexidade:**
    - "Podemos discutir a complexidade de tempo e espaço antes de implementar?"
    - "Existe alguma expectativa específica em termos de complexidade?"

3. **Validações:**
    - "Devo incluir validações para o parâmetro n?"
    - "Como devo tratar valores negativos no array, se houver?"

Esta abordagem demonstra:
- Pensamento analítico
- Atenção aos detalhes
- Consideração por casos de borda

*/
