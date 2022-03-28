package defineDocument

import "strings"

func IsReplacementRequest(docName string) bool {
	var exemple []string

	flag := false
	text := strings.ToLower(docName)
	exemple = append(exemple, "substabelecimento com reserva de poderes",
		"substabelecimento sem reserva de poderes",
		"substabelecimento com reserva de poderes (inativo)",
		"apresentação de substabelecimento sem reserva de poderes",
		"apresentação de substabelecimento com reserva de poderes",
		"instrumento de mandato  apresentação de substabelecimento com reservas de poderes",
		"instrumento de mandato  apresentação de substabelecimento sem reservas de poderes",
		"substabelecimento",
	)

	total := len(exemple)

	for i := 0; i < total; i++ {
		if strings.Contains(text, exemple[i]) {
			flag = true
		}
	}

	return flag

}
