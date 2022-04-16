package defineDocument

import "strings"

func IsReplica(docName string) bool {
	example := []string{
		"manifestação sobre a contestação",
		"réplica", "replica",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
