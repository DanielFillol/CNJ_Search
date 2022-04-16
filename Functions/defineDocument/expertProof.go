package defineDocument

import "strings"

func IsExpertProof(docName string) bool {
	example := []string{
		"sistema bacenjud",
		"bacen",
		"bacenjud",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
