package main

import "fmt"

func main() {
	numero := 30

	numeroInvertido := inverteSinal(numero)
	fmt.Println("Sem ponteiro", numeroInvertido)

	numeroComPonteiro := 40
	inverteSinalComPonteiro(&numeroComPonteiro)
	fmt.Println("Com ponteiro", numeroComPonteiro)
}

// Sem usar ponteiro
func inverteSinal(numero int) int {
	return numero * -1
}

// Usando ponteiro
func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}
