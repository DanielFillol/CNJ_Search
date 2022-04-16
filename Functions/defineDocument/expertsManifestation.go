package defineDocument

import "strings"

func IsExpertsManifestation(docName string) bool {
	example := []string{
		"peticionamento eletrônico  petição peritos",
		"petição peritos",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
