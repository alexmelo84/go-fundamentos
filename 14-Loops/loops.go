package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Imprimindo i", i)
		time.Sleep(time.Second)
	}

	// Inicializa no for
	for j := 0; j < 10; j += 2 {
		fmt.Println("Imprimindo j", j)
		time.Sleep(time.Second)
	}

	// Range
	nomes := [3]string{"Alex", "João", "Paulo"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	// Loop no map
	usuario := map[string]string {
		"nome": "Alex",
		"sobrenome": "Melo",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	/* for {
		fmt.Println("Infinito")
		time.Sleep(time.Second)
	} */
}
