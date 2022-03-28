package defineDocument

import "strings"

func IsLetterOfAttorney(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "procuração",
		"apresentação de procuração",
		"apresentação de procuração",
		"instrumento de mandato  apresentação de procuração",
		"procuração",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
