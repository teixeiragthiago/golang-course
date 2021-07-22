package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Thiago"
	u.idade = 24

	fmt.Println(u)

	//Inferência de tipos
	enderecoEx := endereco{"Rua Santa Cecilia", 0}

	user2 := usuario{"Bruno", 26, enderecoEx}
	fmt.Println(user2)

	user3 := usuario{nome: "Rafael"}
	fmt.Println(user3)

}
