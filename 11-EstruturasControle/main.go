package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("É maior que ZERO")
	} else if numero < -10 {
		fmt.Println("É menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
