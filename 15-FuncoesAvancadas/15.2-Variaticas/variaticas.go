package main

import "fmt"

func main() {
	resultado := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(resultado)

	escrever("Texto do parâmetro", 1, 2, 3, 4, 5)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
