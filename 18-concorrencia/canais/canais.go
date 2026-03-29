package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

// Canais (chan) são a forma do Go de comunicar goroutines com segurança.
// Em vez de compartilhar memória, você passa valores pelo canal.
//
// Como funciona nesse código:
//
//  main                         goroutine escrever
//    |                                |
//    | make(chan string)               |
//    | go escrever(...)  ─────────────>|
//    | fmt.Println(...)               | canal <- "Olá Mundo!" (envia)
//    | range canal <──────────────────| (main recebe, imprime)
//    |                                | time.Sleep(1s)
//    |                                | canal <- "Olá Mundo!" (envia)
//    | (recebe, imprime) <────────────|
//    |                                | ... (5x no total)
//    |                                | close(canal)
//    | range termina                  |
//    | fmt.Println("Fim!")            |
//
// Pontos principais:
//   - make(chan string)  → cria um canal que transporta valores do tipo string
//   - canal <- texto     → a goroutine ENVIA um valor; fica bloqueada até alguém receber
//   - range canal        → o main RECEBE cada valor; fica bloqueado esperando o próximo
//   - close(canal)       → sinaliza que não haverá mais envios, encerrando o range
//                          sem isso o range esperaria para sempre (deadlock)