package defineDocument

import "strings"

func IsMemorial(docName string) bool {
	example := []string{
		"memorial",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
