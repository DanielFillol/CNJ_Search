package defineDocument

import "strings"

func IsHearingRequest(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "requerimento de adiamento de audiência",
		"audiência  requerimento de designação",
		"audiência  requerimento de adiamento",
		"audiência  requerimento de antecipação",
		"audiência realizada exitosa",
		"audiência",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
