package main

import "fmt"

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {
	fmt.Println("Funções Recursivas")
	// 1 1 2 3 5 8 13 21

	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))
}
