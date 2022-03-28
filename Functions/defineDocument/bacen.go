package defineDocument

import "strings"

func IsBacen(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "prova pericial  art 474 do cpc",
		"474 do cpc",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag
}
