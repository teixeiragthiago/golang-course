package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Concorrência é diferente de paralelismo

		Paralelismo-> Acontece quando você tem duas tarefas
		sendo executadas exatamente ao mesmo tempo, mas isso só é
		possível se você tiver um processador com mais de UM núcleo,
		e com isso ele distribui uma tarefa para cada núcleo.

		Concorrência-> Não necessáriamente executam ao mesmo tempo, embora podem!
		Porém, se você só tiver um processador com um único núcleo; com isso,
		a tarefa A não esperaria a tarefa B terminar para executar, as duas
		ficariam revesando entre si a execução.
	*/

	go escrever("Olá mundo!") //go routine
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
