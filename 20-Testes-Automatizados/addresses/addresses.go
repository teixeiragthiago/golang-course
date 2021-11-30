package addresses

import "strings"

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
