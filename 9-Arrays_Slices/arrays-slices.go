package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1[5] int
	fmt.Println(array1)

	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"Primeiro", "Segundo", "Terceiro", "Quarto", "Quinto"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{10, 20, 30}
	fmt.Println(slice1)

	slice1 = append(slice1, 40)
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice1))

	slice2 := array2[0:2]
	fmt.Println(slice2)

	array2[0] = "Primeiro editado"
	fmt.Println(slice2)
}
