package main

import "fmt"

func main() {
	tarefas := make(chan int, 50)
	resultados := make(chan int, 50)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 50; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 50; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

var chamadas uint

func fibonacci(posicao int) int {
	chamadas++
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
