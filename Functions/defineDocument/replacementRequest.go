package defineDocument

import "strings"

func IsReplacementRequest(docName string) bool {
	example := []string{
		"substabelecimento com reserva de poderes",
		"substabelecimento sem reserva de poderes",
		"substabelecimento com reserva de poderes (inativo)",
		"apresentação de substabelecimento sem reserva de poderes",
		"apresentação de substabelecimento com reserva de poderes",
		"instrumento de mandato  apresentação de substabelecimento com reservas de poderes",
		"instrumento de mandato  apresentação de substabelecimento sem reservas de poderes",
		"substabelecimento",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
