package defineDocument

import "strings"

func IsDiverseRequest(docName string) bool {
	example := []string{
		"petições diversas",
		"petição (outras)",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
