package defineDocument

import "strings"

func IsReplica(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "manifestação sobre a contestação", "réplica", "replica")

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag
}
