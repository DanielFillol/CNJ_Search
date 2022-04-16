package defineDocument

import "strings"

func IsInitialRequest(docName string) bool {
	example := []string{
		"páginas 1 ",
		"página 1 ",
		"paginas 1 ",
		"pagina 1 ",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
