package main

import (
	"fmt"
	"strconv"
	"strings"
)

// struct Transaction
type Transaction struct {
	ContaOrigem  string
	ContaDestino string
	Valor        float64
	Data         string
}

// dividir a string e converter os valores para a struct
func ParseTransaction(s string) (*Transaction, error) {
	parts := strings.Fields(s)

	if len(parts) != 4 {
		return nil, fmt.Errorf("formato inválido")
	}

	amount, _ := strconv.ParseFloat(parts[2], 64)
	t := Transaction{
		ContaOrigem:  parts[0],
		ContaDestino: parts[1],
		Valor:        amount,
		Data:         parts[3],
	}

	return &t, nil
}

var accountHistory = make(map[string][]Transaction) // Mapa de histórico
var transactionSet = make(map[string]struct{})      // Mapa para evitar duplicidade

func ingest(transaction string) {
	if _, exists := transactionSet[transaction]; exists { // Verifica duplicidade
		fmt.Println("rejeitada!")
		return // Não insere se já existe
	}

	t, err := ParseTransaction(transaction)
	if err != nil {
		return
	}

	fmt.Println(*t)
	accountHistory[t.ContaOrigem] = append(accountHistory[t.ContaOrigem], *t)
	accountHistory[t.ContaDestino] = append(accountHistory[t.ContaDestino], *t)
	transactionSet[transaction] = struct{}{} // Marca transação como inserida
}

func fetchHistory(accountID string) []Transaction {
	return accountHistory[accountID]
}

func main() {
	ingest("conta1 conta2 200.00 2023-01-31")
	ingest("conta1 conta2 200.00 2023-01-31") // Tenta inserir duplicada
	ingest("conta2 conta3 100.00 2023-01-31") // Insere outra transação
	
	fmt.Println()
	fmt.Println("History")
	history := fetchHistory("conta1") // Busca histórico da conta1
	for _, t := range history {       // Itera sobre histórico
		fmt.Println(t) // Imprime transação
	}
}
