package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Thiago",
		"sobrenome": "Teixeira",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Thiago",
			"segundo":  "Teixeira",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso") //função nativa do GO para deletar chave do MAP
	fmt.Println(usuario2)

	//Adicionando nova chave no MAP
	usuario2["signo"] = map[string]string{
		"nome": "Áries",
	}

	fmt.Println(usuario2)

}
