package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2

}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multplicacao, divisao
}

func main() {

	//se eu passar o underline não preciso usar a variável do retorno da funcao
	resultCalcSoma, resultCalcSub, resultCalcMult, _ := calculosMatematicos(10, 15)
	fmt.Println(resultCalcSoma, resultCalcSub, resultCalcMult)

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)
}
