package main

import "fmt"

func main() {
	idAtributo := map[string]int{
		"Força":        1,
		"Destreza":     2,
		"Constituição": 3,
		"Intelecto":    4,
	}

	for index, atributo := range idAtributo {
		fmt.Printf("O ID do atributo de %s é %v \n", index, atributo)
	}

}
