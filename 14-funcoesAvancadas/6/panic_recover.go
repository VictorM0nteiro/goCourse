package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) (bool, float64) {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true, media
	} else if media < 6 {
		return false, media
	}

	panic("A MEDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(10, 9))
	fmt.Println("Pós execucao")
}