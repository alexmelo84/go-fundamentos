package main

import "fmt"

func main() {
	// Exemplo sem ponteiro, atribuição simples de variável de outra variável
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Exemplo usando ponteiro
	var variavel3 int
	var ponteiro1 *int
	variavel3 = 20
	ponteiro1 = &variavel3
	fmt.Println(variavel3, ponteiro1, *ponteiro1)

	// Se mudar o valor da variável, o ponteiro vai pegar o valor alterado
	variavel3 = 40
	fmt.Println(variavel3, ponteiro1, *ponteiro1)
}
