package main

import "fmt"

func main() {
	// If
	numero1 := 8
	if numero1 > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor ou igual a 10")
	}

	// If init - a variável só fica acessível dentro da condição
	if numero2 := numero1; numero2 > 5 {
		fmt.Println("Maior que 5")
		fmt.Println(numero2)
	} else if numero2 < 10 {
		fmt.Println("Menor que 10")
	} else {
		fmt.Println("Menor ou igual a 10")
	}
}
