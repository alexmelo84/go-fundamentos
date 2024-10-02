package main

import "fmt"

func main() {
	generica("ok")
	generica(2)
	generica(true)
}

func generica(i interface{}) {
	fmt.Println(i)
}
