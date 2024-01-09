package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ValidateAddressType validates the address type
func ValidateAddressType(address string) string {
	addressTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	firstWord := strings.Split(strings.ToLower(address), " ")[0]

	isValidAddressType := false
	for _, addressType := range addressTypes {
		if addressType == firstWord {
			isValidAddressType = true
		}
	}

	if isValidAddressType {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(firstWord)
	}

	return "Invalid Address Type"
}
