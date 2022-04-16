package defineDocument

import "strings"

func IsDiverseMiddleRequest(docName string) bool {
	example := []string{
		"petições intermediárias diversas",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
