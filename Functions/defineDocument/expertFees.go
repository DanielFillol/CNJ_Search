package defineDocument

import "strings"

func IsExpertFees(docName string) bool {
	example := []string{
		"requisição final de honorários periciais",
		"apresentação de proposta de honorários periciais",
		"honorários periciais",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
