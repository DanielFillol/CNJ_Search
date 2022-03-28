package defineDocument

import "strings"

func IsInitialRequest(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "páginas 1 ", "página 1 ", "paginas 1 ", "pagina 1 ")

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) || text == "página 1" {
			flag = true
		}
	}

	return flag
}
