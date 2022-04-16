package defineDocument

import "strings"

func IsQualificationRequest(docName string) bool {
	example := []string{
		"solicitação de habilitação",
		"habilitação",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
