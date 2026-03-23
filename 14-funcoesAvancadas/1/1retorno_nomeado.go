package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return soma, subtracao
}

func main() {
	x, y := 0,0
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	varOp1, varOp2 := calculosMatematicos(x, y)
	fmt.Println(varOp1, varOp2)
}