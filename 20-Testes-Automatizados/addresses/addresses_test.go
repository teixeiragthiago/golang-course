package addresses

import "testing"

type testScenario struct {
	endereco       string
	expectedReturn string
}

//Test + NomeFuncao
func TestAddressTypesMustReturnSuccess(t *testing.T) {

	testScenarios := []testScenario{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Tipo inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"AVENIDA REBOUÇAS ", "Avenida"},
		{"Estrada Qualquer", "Estrada"},
		// {"", "Tipo inválido"},
	}

	for _, scen := range testScenarios {
		receivedReturn := AddressTypes(scen.endereco)

		if receivedReturn != scen.expectedReturn {
			t.Errorf("O Tipo recebido '%s' é diferente do tipo esperado '%s'.", receivedReturn, scen.expectedReturn)
		}
	}
}

func TestAddressTypesMustReturnError(t *testing.T) {

	testScenarios := []testScenario{
		{"Rua ABC", "Avenida"},
		{"Avenida Paulista", "Rua"},
		{"Rodovia dos Imigrantes", "Praça"},
		{"Praça das Rosas", "Rodovia inválido"},
		{"Estrada Qualquer", ""},
	}

	for _, scen := range testScenarios {
		receivedReturn := AddressTypes(scen.endereco)

		if receivedReturn == scen.expectedReturn {
			t.Errorf("O Tipo recebido '%s' é diferente do tipo esperado '%s'.", receivedReturn, scen.expectedReturn)
		}
	}
}
