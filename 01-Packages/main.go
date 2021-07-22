package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

//O Arquivo .mod é como se fosse o package.json
// Go build gera um binário

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("thia.com")
	fmt.Println(erro)
}
