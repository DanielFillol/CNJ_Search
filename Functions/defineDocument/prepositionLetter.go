package defineDocument

import "strings"

func IsPrepositonLetter(docName string) bool {
	example := []string{
		"carta de preposição",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
