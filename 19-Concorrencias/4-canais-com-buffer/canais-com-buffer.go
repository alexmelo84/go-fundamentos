package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá, mundo"
	canal <- "Olá, mundo 2"
	// A partir do terceiro, daria deadlock
	// canal <- "Olá, mundo 3"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
