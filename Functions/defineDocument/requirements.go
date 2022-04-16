package defineDocument

import "strings"

func IsRequirements(docName string) bool {
	example := []string{
		"quesitos",
		"quesitos complementares",
		"indicação de assistente técnico",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
