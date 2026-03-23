package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "pedro", 20, 171}
	fmt.Println(p1)

	e1 := estudante{p1, "Sistemas", "UFU"}
	fmt.Println(e1)
}