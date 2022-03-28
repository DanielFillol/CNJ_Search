package defineDocument

import "strings"

func IsExpertsManifestation(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "peticionamento eletrônico  petição peritos",
		"petição peritos",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
