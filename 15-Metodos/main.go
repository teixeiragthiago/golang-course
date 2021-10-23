package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", user.nome)
}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

//Aqui passo o usuario como ponteiro pq estou fazendo uma alteração no usuário
//qnd estou somente acessando uma atributo como nos métodos acima não é necessário
//já qnd vou altera valores, sim!
func (user *usuario) fazerAniversario() {
	user.idade++
}

func (user *usuario) alterarNome(nome string) {
	user.nome = nome
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"David", 24}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2)

	usuario2.alterarNome("Thiago")
	fmt.Println(usuario2)
}
