package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRENCIA != PARALELISMO
	go escrever("Olá Mundo!") //goroutine
	go escrever("programando em go!")
	escrever("Victor")
}

func escrever(texto string) {
	for range 1000000000{
		fmt.Println(texto)
		time.Sleep(2 * time.Second)
	}
}