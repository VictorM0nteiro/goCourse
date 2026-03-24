package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)


	// chamada depois
	somar := func(a, b int) int {
    return a + b
	}
	x := somar(2, 3) // chamada depois

	fmt.Println(x)


	// Passada como argumento (higher-order function):
	executar(func(n int) int { return n * 2 }, 5) // retorna 10
	fmt.Println(executar(func(n int) int { return n * 2 }, 5))

	contador := 0
	incrementar := func() {
		contador++ // acessa 'contador' de fora
	}
	for i := 0; i < 10; i++{
		incrementar()
		incrementar()
	}
	incrementar()
	incrementar()
	fmt.Println(contador) // 2

}

// Passada como argumento (higher-order function):
	func executar(f func(int) int, valor int) int {
    return f(valor)
	}