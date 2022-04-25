package Classifier

import "strings"

func isAssetsToPledge(docName string) bool {
	example := []string{"indicação de bens à penhora", "penhora indicação de bens", "penhora  indicação de bens", "outros bens ou direitos", "nomeação de bens à penhora", "bens"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isBacen(docName string) bool {
	example := []string{"prova pericial  art 474 do cpc"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isCertificate(docName string) bool {
	example := []string{"certidões"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isCounterarguments(docName string) bool {
	example := []string{"contrarrazões de apelação"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isDiverseMiddleRequest(docName string) bool {
	example := []string{"petições intermediárias diversas"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isDiverseRequest(docName string) bool {
	example := []string{"petições diversas", "petição (outras)"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isDocument(docName string) bool {
	example := []string{"documento diverso", "documento"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isExpertClarification(docName string) bool {
	example := []string{"solicitação do perito", "solicitação do perito (digitalizado)"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isExpertFees(docName string) bool {
	example := []string{"requisição final de honorários periciais", "apresentação de proposta de honorários periciais", "honorários periciais"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isExpertProof(docName string) bool {
	example := []string{"sistema bacenjud", "bacen", "bacenjud"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isExpertsManifestation(docName string) bool {
	example := []string{"peticionamento eletrônico  petição peritos", "petição peritos"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isHearingRequest(docName string) bool {
	example := []string{"requerimento de adiamento de audiência", "audiência  requerimento de designação", "audiência  requerimento de adiamento", "audiência  requerimento de antecipação", "audiência realizada exitosa", "audiência"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isInformation(docName string) bool {
	example := []string{"informações"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isInitialRequest(docName string) bool {
	example := []string{"páginas 1 ", "página 1 ", "paginas 1 ", "pagina 1 "}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isInitialRequestEmd(docName string) bool {
	example := []string{"emenda à inicial", "emenda", "emenda a inicial"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isLetterOfAttorney(docName string) bool {
	example := []string{"procuração", "apresentação de procuração", "apresentação de procuração", "instrumento de mandato  apresentação de procuração", "procuração"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isMemorial(docName string) bool {
	example := []string{"memorial"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isPrepositonLetter(docName string) bool {
	example := []string{"carta de preposição"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isProof(docName string) bool {
	example := []string{"indicação de provas"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isProofOfResidence(docName string) bool {
	example := []string{"endereço localizado", "endereço apresentação", "endereço"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isQualificationRequest(docName string) bool {
	example := []string{"solicitação de habilitação", "habilitação"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isReplacementRequest(docName string) bool {
	example := []string{"substabelecimento com reserva de poderes", "substabelecimento sem reserva de poderes", "substabelecimento com reserva de poderes (inativo)", "apresentação de substabelecimento sem reserva de poderes", "apresentação de substabelecimento com reserva de poderes", "instrumento de mandato  apresentação de substabelecimento com reservas de poderes", "instrumento de mandato  apresentação de substabelecimento sem reservas de poderes", "substabelecimento"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isReplica(docName string) bool {
	example := []string{"manifestação sobre a contestação", "réplica", "replica"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isRequirements(docName string) bool {
	example := []string{"quesitos", "quesitos complementares", "indicação de assistente técnico"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}

func isPage(docName string) bool {
	example := []string{"páginas", "página", "paginas", "pagina"}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false
}
