package main

import "fmt"

func main() {
	// Função sinmples
	soma := somar(8, 20)
	fmt.Println(soma)

	// Função anônima
	var funcao1 = func() {
		fmt.Println("Dentro da função 1")
	}
	funcao1()

	var funcao2 = func(string1 string) {
		fmt.Println(string1)
	}
	funcao2("Texto passado por parâmetro")

	var funcao3 = func(string2 string) string {
		return (string2)
	}
	retornoFuncao3 := funcao3("Texto que será retornado")
	fmt.Println(retornoFuncao3)

	// Função com múltiplo retorno
	multiplosRetornos1, multiplosRetornos2 := calculos(8, 2)
	fmt.Println(multiplosRetornos1)
	fmt.Println(multiplosRetornos2)

	// Função com múltiplo retorno descartando algum valor retornado
	multiplosRetornos3, _ := calculos(12, 9)
	fmt.Println(multiplosRetornos3)
}

// Função de soma de dois números inteiros
func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

// Executa uma série de cálculos
func calculos(numero1, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2

	return soma, subtracao
}
