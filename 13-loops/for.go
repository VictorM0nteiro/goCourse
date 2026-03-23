package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 50; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	count := 0
	countSecond := 0
	for {
		fmt.Println("Executando infinitamente")
		countSecond = time.Now().Second()
		time.Sleep(time.Nanosecond)
		count++
		fmt.Println(count)
		if count == 10000{
			fmt.Println(countSecond)
			break
		}
	}
}