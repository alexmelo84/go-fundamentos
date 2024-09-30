package main

import (
	"errors"
	"fmt"
)

// Exemplos de tipos de dados
func main() {
	// Inteiro
	var numero1 int16 = 100
	fmt.Println(numero1)

	// Inteiro sem sinal (não aceita negativos)
	var numero2 uint32 = 192192
	fmt.Println(numero2)

	// Alias para int32
	var numero3 rune = -1312
	fmt.Println(numero3)

	// Alias para uint8
	var numero4 byte = 131
	fmt.Println(numero4)

	// Números reais
	var numero5 float32 = 232.45
	fmt.Println(numero5)

	var numero6 float64 = 23212.2323
	fmt.Println(numero6)

	// Texto
	var string1 string = "Texto qualquer 1"
	string2 := "Texto qualquer 2"
	fmt.Println(string1)
	fmt.Println(string2)

	// Chars
	char1 := 'C'
	fmt.Println(char1)

	// Valor zero. Usado nos casos que inicializamos uma variável sem valor
	var valorZero1 string
	var valorZero2 int8
	fmt.Println(valorZero1)
	fmt.Println(valorZero2)

	// Booleano
	var boolean1 bool = true
	var boolean2 bool = false
	fmt.Println(boolean1)
	fmt.Println(boolean2)

	// Error
	var error1 error = errors.New("Erro 1")
	fmt.Println(error1)
}
