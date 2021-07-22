package main

import (
	"errors"
	"fmt"
)

func main() {

	//Numeros inteiros

	var numero int64 = 1000000000000000
	fmt.Println(numero)

	//uint não pode user negativo

	//alias - 'rune' é um alias para int32
	var numero3 rune = 1234
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//Numeros reais
	var numeroReal1 float32 = 132.43
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 132321321.43
	fmt.Println(numeroReal2)

	numeroReal3 := 132321321.43
	fmt.Println(numeroReal3)

	//Strings

	var str string = "Texto" //aspas duplas para strings, asplas simples para char
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//o tipo CHAR não existe, ele é transformado em um INTEIRO, e retorna o valor do caracter
	//referente da tabela ASCII
	char := 'a'
	fmt.Println(char)

	// o valor zero de string é VAZIO
	var texto string
	fmt.Println(texto)

	//VALOR ZERO - se não atribuir nada na variável
	var valorZero int16
	fmt.Println(valorZero)

	//BOOLEAN

	//o valor ZERO do tipo BOOL é FALSE
	var booleano1 bool
	fmt.Println(booleano1)

	//ERROR - o valor ZERO de Error é NIL
	var erro error = errors.New("ERRO")
	fmt.Println(erro)
}
