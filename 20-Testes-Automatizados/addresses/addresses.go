package addresses

import "strings"

/*
	go test --cover -> para checar a porctg da cobertura de testes
	go test --coverprofile 'filename.txt' -> gerar relatorio
	go tool cover --func=filename.txt -> pra ler o relatório gerado
	go tool cover --html=filename.txt -> jeito mais legal do mundo
	pra ler os relatórios
*/

//AddressTypes validates input address is valid
func AddressTypes(endereco string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	addressToLower := strings.ToLower(endereco)
	firstWord := strings.Split(addressToLower, " ")[0]

	addressHasValidType := false

	for _, tipo := range validTypes {
		if tipo == firstWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}

	return "Tipo inválido"
}
