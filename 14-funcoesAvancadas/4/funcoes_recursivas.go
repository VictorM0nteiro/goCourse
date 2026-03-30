package main

import "fmt"

var chamadas uint

func fibonacci(posicao uint) uint {
	chamadas++
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8 13

	posicao := uint(50)
	for i := uint(1); i <= posicao; i++ {
		chamadas = 0
		resultado := fibonacci(i)
		fmt.Printf("\nFibonacci(%d): %d | chamadas: %d", i, resultado, chamadas)
	}
}

/*
	POR QUE ESSE CÓDIGO FICA TÃO LENTO?

	O problema está na recursão sem memoização (sem cache de resultados).

	Para calcular fibonacci(50), a função chama fibonacci(48) e fibonacci(49).
	Para calcular fibonacci(49), ela chama fibonacci(47) e fibonacci(48).
	Para calcular fibonacci(48), ela chama fibonacci(46) e fibonacci(47).
	... e assim por diante.

	Isso significa que fibonacci(48) é calculado 2 vezes, fibonacci(47) é
	calculado 3 vezes, fibonacci(10) é calculado milhares de vezes — o mesmo
	subproblema é resolvido repetidamente do zero.

	O número total de chamadas cresce de forma exponencial: O(2^n).
	Para fibonacci(50), são feitas aproximadamente 2^50 ≈ 1 quadrilhão de chamadas.

	Além disso, o loop chama fibonacci(i) para cada i de 1 até 50 separadamente,
	sem reaproveitar nenhum resultado anterior — multiplicando ainda mais o trabalho.

	SOLUÇÃO: usar memoização (guardar resultados já calculados em um mapa ou slice)
	ou programação dinâmica (bottom-up), o que reduz a complexidade para O(n).

	Exemplo com memoização:

	var cache = map[uint]uint{}

	func fibonacci(n uint) uint {
		if n <= 1 {
			return n
		}
		if v, ok := cache[n]; ok {
			return v
		}
		cache[n] = fibonacci(n-2) + fibonacci(n-1)
		return cache[n]
	}
*/