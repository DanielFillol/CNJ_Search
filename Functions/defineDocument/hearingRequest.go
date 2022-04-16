package defineDocument

import "strings"

func IsHearingRequest(docName string) bool {
	example := []string{
		"requerimento de adiamento de audiência",
		"audiência  requerimento de designação",
		"audiência  requerimento de adiamento",
		"audiência  requerimento de antecipação",
		"audiência realizada exitosa",
		"audiência",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
