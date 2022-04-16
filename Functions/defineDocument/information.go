package defineDocument

import "strings"

func IsInformation(docName string) bool {
	example := []string{
		"informações",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
