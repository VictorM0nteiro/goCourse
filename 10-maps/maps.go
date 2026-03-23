package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ //tipo da chave e tipo do valor
		"nome": "Pedro",
		"sobrenome": "Silva",
	} 		

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Victor",
			"ultimo" : "Monteiro",
		},
		"curso": {
			"nome": "Sistemas",
			"Universidade": "UFU",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "gemeos",
	}

	fmt.Println(usuario2)
}