package main

import "fmt"

func main() {
	canal := make(chan string, 200)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	canal <- "Programando em Go De Novo!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}

/*
	CANAIS COM BUFFER

	Por padrão, canais em Go são sem buffer (unbuffered): o envio bloqueia
	até que outra goroutine faça a leitura, e vice-versa.

	Um canal com buffer aceita um número fixo de valores sem precisar de um
	receptor imediato. O envio só bloqueia quando o buffer está cheio,
	e a leitura só bloqueia quando o buffer está vazio.

	Sintaxe:
		canal := make(chan tipo, capacidade)

	Exemplo deste arquivo:
		canal := make(chan string, 200)  // buffer de 200 posições

	Com isso, é possível enviar até 200 mensagens sem que nenhuma goroutine
	esteja lendo ao mesmo tempo — útil para desacoplar produtor e consumidor.

	Boas práticas:
	- Prefira canais sem buffer quando quiser sincronização explícita entre goroutines.
	- Use buffer quando o produtor puder ser mais rápido que o consumidor por um curto período.
	- Evite buffers muito grandes para esconder problemas de design de concorrência.
*/