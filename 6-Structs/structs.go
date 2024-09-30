package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	cep string
}

func main() {
	var endereco1 endereco
	endereco1.logradouro = "Rua teste"
	endereco1.cep = "CEP"

	var usuario1 usuario
	// Valor zero
	fmt.Println(usuario1)

	usuario1.nome = "Nome 1"
	usuario1.idade = 20
	fmt.Println(usuario1)
	fmt.Println(usuario1.nome)
	fmt.Println(usuario1.idade)

	// Inferência de tipo
	usuario2 := usuario{"Nome 2", 30, endereco1}
	fmt.Println(usuario2)

	// Inicializando sem ter todos os valores
	usuario3 := usuario{nome: "Nome 3"}
	fmt.Println(usuario3)
}
