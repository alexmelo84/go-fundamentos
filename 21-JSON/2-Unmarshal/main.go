package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome":"Rex","raca":"Pinscher","idade":2}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2JSON := `{"nome":"Pingo","raca":"Corgi"}`
	c2 := make(map[string]string)
	if erro2 := json.Unmarshal([]byte(cachorro2JSON), &c2); erro2 != nil {
		log.Fatal(erro2)
	}

	fmt.Println(c2)
}
