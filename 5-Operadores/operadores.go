package main

import "fmt"

func main () {
	// Aritméticos
	soma1 := 8 + 4
	subtracao1 := 8 - 4
	divisao1 := 8 / 4
	multiplicacao1 := 8 * 4
	restoDivisao1 := 8 % 4
	fmt.Println(soma1)
	fmt.Println(subtracao1)
	fmt.Println(divisao1)
	fmt.Println(multiplicacao1)
	fmt.Println(restoDivisao1)

	// Atribuição explícita
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Atribuição implícita
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println("--------------------------------------------")

	// Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro && falso)
	fmt.Println(verdadeiro || !falso)

	// Unários
	numero1 := 1
	fmt.Println(numero1)
	numero1++
	fmt.Println(numero1)
	numero1 += 15
	fmt.Println(numero1)
	numero1--
	fmt.Println(numero1)
	numero1 *= 3
	fmt.Println(numero1)
	numero1 -= 10
	fmt.Println(numero1)
	numero1 /= 2
	fmt.Println(numero1)
	numero1 %= 4
	fmt.Println(numero1)

	// Ternários - não há ternários no Go, usa-se if/else nesses casos
	var string1 string
	if numero1 > 2 {
		string1 = "Maior que 2"
	} else {
		string1 = "Menor ou igual a 2"
	}
	fmt.Println(string1)
}
