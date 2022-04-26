package Classifier

//All functions here utilizes a [] of string that indicates 1 of 23
//most common document names that are often misspelled by lawyers

//return true for Assets to Pledge document (AssetsToPledge)
func isAssetsToPledge(docName string) bool {
	example := []string{"indicação de bens à penhora", "penhora indicação de bens", "penhora  indicação de bens", "outros bens ou direitos", "nomeação de bens à penhora", "bens"}
	return getMatchWord(docName, example)
}

//return true for Brazil Nation Bank Legal System (Bacen)
func isBacen(docName string) bool {
	example := []string{"prova pericial  art 474 do cpc"}
	return getMatchWord(docName, example)
}

//return true for a legal document emitted by a notary (Certificate)
func isCertificate(docName string) bool {
	example := []string{"certidões"}
	return getMatchWord(docName, example)
}

//return true for Counter Arguments (Counterarguments)
func isCounterarguments(docName string) bool {
	example := []string{"contrarrazões de apelação"}
	return getMatchWord(docName, example)
}

//return true for Diverse Middle Request (DiverseMiddleRequest)
func isDiverseMiddleRequest(docName string) bool {
	example := []string{"petições intermediárias diversas"}
	return getMatchWord(docName, example)
}

//return true for common Request (DiverseRequest)
func isDiverseRequest(docName string) bool {
	example := []string{"petições diversas", "petição (outras)"}
	return getMatchWord(docName, example)
}

//return true for common legal document (Document)
func isDocument(docName string) bool {
	example := []string{"documento diverso", "documento", "diverso"}
	return getMatchWord(docName, example)
}

//return true for Expert Clarification (ExpertClarification)
func isExpertClarification(docName string) bool {
	example := []string{"solicitação do perito", "solicitação do perito (digitalizado)"}
	return getMatchWord(docName, example)
}

//return true for Expert Fees (ExpertFees)
func isExpertFees(docName string) bool {
	example := []string{"requisição final de honorários periciais", "apresentação de proposta de honorários periciais", "honorários periciais"}
	return getMatchWord(docName, example)
}

//return true for Expert proof (ExpertProof)
func isExpertProof(docName string) bool {
	example := []string{"sistema bacenjud", "bacen", "bacenjud"}
	return getMatchWord(docName, example)
}

//return true for Experts Manifestation (ExpertsManifestation)
func isExpertsManifestation(docName string) bool {
	example := []string{"peticionamento eletrônico  petição peritos", "petição peritos"}
	return getMatchWord(docName, example)
}

//return true for Hearing Request (HearingRequest)
func isHearingRequest(docName string) bool {
	example := []string{"requerimento de adiamento de audiência", "audiência  requerimento de designação", "audiência  requerimento de adiamento", "audiência  requerimento de antecipação", "audiência realizada exitosa", "audiência"}
	return getMatchWord(docName, example)
}

//return true for legal request of information (Information)
func isInformation(docName string) bool {
	example := []string{"informações"}
	return getMatchWord(docName, example)
}

//return true for Initial Request (InitialRequest)
func isInitialRequest(docName string) bool {
	example := []string{"páginas 1 ", "página 1 ", "paginas 1 ", "pagina 1 ", "inicial"}
	return getMatchWord(docName, example)
}

//return true for Amendment of Initial Request (InitialRequestEmd)
func isInitialRequestEmd(docName string) bool {
	example := []string{"emenda à inicial", "emenda", "emenda a inicial"}
	return getMatchWord(docName, example)
}

//return true for Letter of Attorney (ReplacementRequest)
func isLetterOfAttorney(docName string) bool {
	example := []string{"procuração", "apresentação de procuração", "apresentação de procuração", "instrumento de mandato  apresentação de procuração", "procuração"}
	return getMatchWord(docName, example)
}

//return true for memorial request (Memorial)
func isMemorial(docName string) bool {
	example := []string{"memorial"}
	return getMatchWord(docName, example)
}

//return true for Prepositon Letter (PrepositonLetter)
func isPrepositonLetter(docName string) bool {
	example := []string{"carta de preposição"}
	return getMatchWord(docName, example)
}

//return true for proof request (Proof)
func isProof(docName string) bool {
	example := []string{"indicação de provas"}
	return getMatchWord(docName, example)
}

//return true for proof of Residence (ProofOfResidence)
func isProofOfResidence(docName string) bool {
	example := []string{"endereço localizado", "endereço apresentação", "endereço"}
	return getMatchWord(docName, example)
}

//return true for Qualification Request (QualificationRequest)
func isQualificationRequest(docName string) bool {
	example := []string{"solicitação de habilitação", "habilitação"}
	return getMatchWord(docName, example)
}

//return true for Replacement Request (ReplacementRequest)
func isReplacementRequest(docName string) bool {
	example := []string{"substabelecimento com reserva de poderes", "substabelecimento sem reserva de poderes", "substabelecimento com reserva de poderes (inativo)", "apresentação de substabelecimento sem reserva de poderes", "apresentação de substabelecimento com reserva de poderes", "instrumento de mandato  apresentação de substabelecimento com reservas de poderes", "instrumento de mandato  apresentação de substabelecimento sem reservas de poderes", "substabelecimento"}
	return getMatchWord(docName, example)
}

//return true for replica (Replica)
func isReplica(docName string) bool {
	example := []string{"manifestação sobre a contestação", "réplica", "replica"}
	return getMatchWord(docName, example)
}

//return true for requirements (Requirements)
func isRequirements(docName string) bool {
	example := []string{"quesitos", "quesitos complementares", "indicação de assistente técnico"}
	return getMatchWord(docName, example)
}

//return true for docName containing "page"
func isPage(docName string) bool {
	example := []string{"páginas", "página", "paginas", "pagina"}
	return getMatchWord(docName, example)
}
