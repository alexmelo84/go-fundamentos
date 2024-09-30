package main

import "fmt"

// Formas de inicializar uma variável inferindo tipos
func main() {
	// Declaração explícita
	var string1 string = "string 1"
	// Declaração implícita
	string2 := "string 2"
	fmt.Println(string1)
	fmt.Println(string2)

	// Declaração explícita múltipla de variáveis
	var (
		string3 string = "string 3"
		string4 string = "string 4"
	)
	fmt.Println(string3)
	fmt.Println(string4)

	// Declaração implícita múltipla de variáveis
	string5, string6 := "string 5", "string 6"
	fmt.Println(string5)
	fmt.Println(string6)

	// Constante
	const const1 string = "constante 1"
	fmt.Println(const1)

	// Inverter valores de variáveis
	string5, string6 = string6, string5
	fmt.Println(string5)
	fmt.Println(string6)
}
