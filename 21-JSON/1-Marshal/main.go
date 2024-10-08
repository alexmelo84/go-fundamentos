package main

import (
	"bytes"
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
	c := cachorro{"Rex", "Pinscher", 2}
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

	c2 := map[string]string {
		"nome": "Pingo",
		"raca": "Corgi",
	}
	cachorro2JSON, erro2 := json.Marshal(c2)
	if erro2 != nil {
		log.Fatal(erro2)
	}
	fmt.Println(cachorro2JSON)
	fmt.Println(bytes.NewBuffer(cachorro2JSON))
}
