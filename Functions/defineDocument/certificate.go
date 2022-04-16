package defineDocument

import "strings"

func IsCertificate(docName string) bool {
	example := []string{
		"certidões",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
