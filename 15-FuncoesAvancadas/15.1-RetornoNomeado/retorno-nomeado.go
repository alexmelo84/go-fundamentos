package main

import "fmt"

func main() {
	soma, subtracao := calculos(20, 5)
	fmt.Println(soma, subtracao)
}

func calculos(numero1, numero2 int) (soma int, subtracao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2

	return
}
