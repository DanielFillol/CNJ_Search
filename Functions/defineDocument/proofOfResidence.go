package defineDocument

import "strings"

func IsProofOfResidence(docName string) bool {
	example := []string{
		"endereço localizado",
		"endereço apresentação",
		"endereço",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
