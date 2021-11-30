package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case mensagegemCanal1 := <-canal1:
			fmt.Println(mensagegemCanal1)
		case mensagegemCanal2 := <-canal2:
			fmt.Println(mensagegemCanal2)
		}

	}
}
