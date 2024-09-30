package main

import (
	"fmt"
	"modulo1/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Primeiro módulo")
	auxiliar.Escrever()

	fail := checkmail.ValidateFormat("abcdefg")
	success := checkmail.ValidateFormat("success@teste.com")
	fmt.Println(fail)
	fmt.Println(success)
}

// nome do módulo: modulo1
