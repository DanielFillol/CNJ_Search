package defineDocument

import "strings"

func IsExpertFees(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "requisição final de honorários periciais",
		"apresentação de proposta de honorários periciais",
		"honorários periciais",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
