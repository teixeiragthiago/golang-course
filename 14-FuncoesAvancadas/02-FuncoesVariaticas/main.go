package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 200, 400, 140, 10)
	fmt.Println(totalSoma)

	escrever("Olá mundo!", 3, 231, 34, 13, 4, 7, 3, 9)
}
