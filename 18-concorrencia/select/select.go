package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}

/*
	SELECT

	O select é parecido com o switch, mas voltado para canais.
	Ele aguarda até que um dos cases esteja pronto (ou seja, um canal tenha
	uma mensagem disponível para leitura ou escrita) e executa esse case.

	- Se mais de um case estiver pronto ao mesmo tempo, o Go escolhe um aleatoriamente.
	- Sem um case default, o select bloqueia até que algum canal esteja pronto.
	- Com um case default, o select nunca bloqueia: se nenhum canal estiver pronto,
	  o default é executado imediatamente.

	Exemplo com default:

	select {
	case msg := <-canal1:
		fmt.Println(msg)
	case msg := <-canal2:
		fmt.Println(msg)
	default:
		fmt.Println("nenhum canal pronto")
	}

	O select é a forma idiomática em Go de multiplexar múltiplos canais
	sem precisar de goroutines extras ou polling manual.
*/