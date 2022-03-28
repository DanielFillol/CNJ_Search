package defineDocument

import "strings"

func IsAssetsToPledge(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "indicação de bens à penhora",
		"penhora indicação de bens",
		"penhora  indicação de bens",
		"outros bens ou direitos",
		"nomeação de bens à penhora",
		"bens",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
