package defineDocument

import "strings"

func IsLetterOfAttorney(docName string) bool {
	example := []string{
		"procuração",
		"apresentação de procuração",
		"apresentação de procuração",
		"instrumento de mandato  apresentação de procuração",
		"procuração",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
