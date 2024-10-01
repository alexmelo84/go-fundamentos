package main

import "fmt"

func main() {
	usuario := map[string]string {
		"nome": "Alex",
		"sobrenome": "Melo",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	// Map aninhado
	usuario2 := map[string]map[string]string {
		"usuario1": {
			"nome": "Alex",
			"sobrenome": "Melo",
		},
		"usuario2": {
			"nome": "Alex2",
			"sobrenome": "Melo2",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["usuario1"])
	fmt.Println(usuario2["usuario2"])

	usuario2["usuario3"] = map[string]string {
		"nome": "Alex3",
		"sobrenome": "Melo3",
	}
	fmt.Println(usuario2)

	delete(usuario, "sobrenome")
	fmt.Println(usuario)
}
