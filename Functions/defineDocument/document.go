package defineDocument

import "strings"

func IsDocument(docName string) bool {
	example := []string{
		"documento diverso",
		"documento",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
