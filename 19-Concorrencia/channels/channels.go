package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go escrever("Olá mundo!", channel)

	fmt.Println("Depois da função escrever começar a ser executada!")

	// for {
	// 	mensagem, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range channel {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
