package main

import "fmt"

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Thiago", "Pedro", "Lucas"}

	// for i, nome := range nomes {
	// 	fmt.Println(i, nome)
	// }

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//Código da ASCII se não chamar string()
	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Thiago",
		"sobrenome": "Teixeira",
	}

	fmt.Println(usuario)

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	//loop infinito
	/*
		for {

		}
	*/

}
