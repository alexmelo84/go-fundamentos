package main

import "fmt"

func main() {
	// make(tipo, tamanho, máximo de itens)
	slice := make([]float32, 10, 11)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice)

	// Passando do limite da capacidade
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(slice, len(slice), cap(slice))

	slice2 := make([]float32, 5)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// Passando do limite da capacidade
	slice2 = append(slice2, 1)
	fmt.Println(slice2)
	fmt.Println(slice2, len(slice2), cap(slice2))
}