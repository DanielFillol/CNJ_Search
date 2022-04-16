package defineDocument

import "strings"

func IsExpertClarification(docName string) bool {
	example := []string{
		"solicitação do perito",
		"solicitação do perito (digitalizado)",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
