package main

import "fmt"

type usuario struct {
	nome string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o usuário %s\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	if u.idade >= 18 {
		return true
	}
	
	return false
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {
	escrever()
}

func escrever() {
	fmt.Println("Escrevendo da função")

	usuario1 := usuario{"Alex Melo", 39}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorIdade())
	usuario1.fazAniversario()
	fmt.Println(usuario1.idade)

	usuario2 := usuario{"Alex Melo 2", 16}
	fmt.Println(usuario2)
	usuario2.salvar()
	fmt.Println(usuario2.maiorIdade())
}
