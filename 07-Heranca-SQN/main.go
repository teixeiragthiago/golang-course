package main

import "fmt"

/*
	Em GO não existe herança, até porque não é uma linguage Orientada a Objetos
	porém, se eu criar um struct genérico e passar ele dentro de outro struct,
	sem passar o tipo de struct como está no exemplo abaixo, é como se fosse
	uma herança, mas lembrando: NÃO É HERANÇA!
*/

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 186}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciencia da Computacao", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.altura)
	fmt.Println(e1.curso)

}
