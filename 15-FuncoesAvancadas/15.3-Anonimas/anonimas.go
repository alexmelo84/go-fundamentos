package main

import "fmt"

func main() {
	func() {
		fmt.Println("Sou uma função anônima")	
	}()

	func(texto string) {
		fmt.Println(texto)	
	}("Sou um parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Retorno com o parâmetro '%s'", texto)
	}("Sou um parâmetro")
	fmt.Println(retorno)
}
