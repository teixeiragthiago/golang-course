package main

import "fmt"

func main() {
	fmt.Println("Função MAIN")
	fmt.Println(n)
}

var n int

func init() {

	n = 10
	fmt.Println("A Função INIT é sempre executada ANTES da função MAIN")
	fmt.Println("Eu posso ter UMA função INIT por ARQUIVO, e não por PACOTE")
}
