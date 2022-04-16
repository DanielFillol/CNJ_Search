package defineDocument

import "strings"

func IsPage(docName string) bool {
	example := []string{
		"páginas",
		"página",
		"paginas",
		"pagina",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
