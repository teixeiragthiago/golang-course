package main

import (
	"fmt"
)

func main() {

	//Funcão anônima
	func(texto string) {
		fmt.Println(texto)
	}("Olá mundo!") //chamando a func anonima

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
}
