package defineDocument

import "strings"

func IsBacen(docName string) bool {
	example := []string{
		"prova pericial  art 474 do cpc",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
