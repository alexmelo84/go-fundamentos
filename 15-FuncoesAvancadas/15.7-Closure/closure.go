package main

import "fmt"

func main() {
	texto := "Na função main()"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

func closure() func() {
	texto := "Na função closure()"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}
