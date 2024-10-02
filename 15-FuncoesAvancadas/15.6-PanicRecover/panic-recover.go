package main

import "fmt"

func main() {
	fmt.Println(aprovado(7, 7))
	fmt.Println("Última execução")
}

func aprovado(nota1, nota2 float32) bool {
	defer recuperaExecucao()

	media := (nota1 + nota2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}

	panic("Média é exatamente 7!")
}

func recuperaExecucao() {
	fmt.Println("Tentando recuperar execução")

	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso!")
	}
}
