package main

import "fmt"

func main() {
	defer funcao1()
	funcao2()

	aprovado(10, 4)
}

func funcao1 () {
	fmt.Println("Na função 1")
}

func funcao2 () {
	fmt.Println("Na função 2")
}

func aprovado (nota1, nota2 float32) bool {
	defer fmt.Println("Resultado será retornado")
	fmt.Println("Verificando se está aprovado")

	media := (nota1 + nota2) / 2

	if media >= 7 {
		return true
	}

	return false
}
