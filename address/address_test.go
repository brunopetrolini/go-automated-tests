package address

import "testing"

type testingScenario struct {
	inputAddress   string
	expectedReturn string
}

func TestAddressType(t *testing.T) {
	testingScenarios := []testingScenario{
		{"Rua dos Campos", "Rua"},
		{"Avenida Wenceslau Braz", "Avenida"},
		{"Estrada do Mato Dentro", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalid Address Type"},
		{"RUA DOS CAMPOS", "Rua"},
		{"avenida wenceslau braz", "Avenida"},
		{"", "Invalid Address Type"},
	}

	for _, scenario := range testingScenarios {
		returnedAddressType := ValidateAddressType(scenario.inputAddress)
		if returnedAddressType != scenario.expectedReturn {
			t.Errorf("Address type is wrong, got: %s, want: %s.", returnedAddressType, scenario.expectedReturn)
		}
	}
}
