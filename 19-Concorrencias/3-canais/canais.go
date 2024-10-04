package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá, mundo!", canal)

	fmt.Println("Após a função escrever()")

	/* for {
		mensagem, aberto := <-canal

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	} */

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Depois de esperar o channel receber valor")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
