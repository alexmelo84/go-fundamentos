package main

import "fmt"

func main() {
	posicao := uint(20)
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}
