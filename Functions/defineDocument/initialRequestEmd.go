package defineDocument

import "strings"

func IsInitialRequestEmd(docName string) bool {
	example := []string{
		"emenda à inicial",
		"emenda",
		"emenda a inicial",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
