package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	campus string
}

func main () {
	pessoa1 := pessoa{"Alex", "Melo", 39, 175}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Computação", "Campus 1"}
	fmt.Println(estudante1)
}
