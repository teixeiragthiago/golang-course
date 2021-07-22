package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" //declaração explícita
	variavel2 := "Variável 2"           //declaração implícita (inferência de tipo)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3"
		variavel4 string = "var 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//trocar valores das variáveis sem precisar de uma auxiliar
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
