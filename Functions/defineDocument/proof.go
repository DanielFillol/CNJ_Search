package defineDocument

import "strings"

func IsProof(docName string) bool {
	example := []string{
		"indicação de provas",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
