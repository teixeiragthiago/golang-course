package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multip := 10 * 4
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multip, restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2

	fmt.Println(soma2)
	//Fim dos aritméticos

	//Atribuição

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Fim atribuição

	//Operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Fim operadores relacionais

	//Operadores Lógicos

	fmt.Println("----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Fim dos operadores Lógicos

	//Operadores unários
	fmt.Println("----------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 5
	fmt.Println(numero)

	numero *= 3
	numero /= 4
	numero %= 2

	fmt.Println(numero)

	//Fim dos operadores unários

	//GO NÃO TEM OPERADORES TERNÁRIOS :/
}
