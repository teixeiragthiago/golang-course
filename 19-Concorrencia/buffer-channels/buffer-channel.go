package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Olá mundo!"
	channel <- "Programando em go!"

	mensagem := <-channel
	mensagem2 := <-channel
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
