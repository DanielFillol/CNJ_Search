package Functions

import "github.com/Darklabel91/LegalDoc_Classifier/Functions/defineDocument"

func newName(docName string) string {
	if defineDocument.IsInitialRequest(docName) {
		return "Petição Inicial"
	} else if defineDocument.IsInitialRequestEmd(docName) {
		return "Emenda à Inicial"
	} else if defineDocument.IsDiverseRequest(docName) {
		return "Petição (Outras)"
	} else if defineDocument.IsDiverseMiddleRequest(docName) {
		return "Petição Intermediária"
	} else if defineDocument.IsDocument(docName) {
		return "Documentos Diversos"
	} else if defineDocument.IsQualificationRequest(docName) {
		return "Pedido de Habilitação"
	} else if defineDocument.IsReplacementRequest(docName) {
		return "Instrumento de Procuração"
	} else if defineDocument.IsLetterOfAttorney(docName) {
		return "Instrumento de Procuração"
	} else if defineDocument.IsProofOfResidence(docName) {
		return "Comprovante de Endereço"
	} else if defineDocument.IsAssetsToPledge(docName) {
		return "Pedido de Penhora"
	} else if defineDocument.IsHearingRequest(docName) {
		return "Pedido de Designação/Redesignação de Audiência"
	} else if defineDocument.IsRequirements(docName) {
		return "Apresentação de Quesitos/Indicação de Assistente Técnico"
	} else if defineDocument.IsExpertFees(docName) {
		return "Solicitação de Honorários de Perito"
	} else if defineDocument.IsExpertClarification(docName) {
		return "Esclarecimento de Perito"
	} else if defineDocument.IsExpertsManifestation(docName) {
		return "Manifestação do Perito"
	} else if defineDocument.IsExpertProof(docName) {
		return "Agendamento de Vistoria - Prova Pericial - Art. 474 do CPC"
	} else if defineDocument.IsBacen(docName) {
		return "Sisbajud"
	} else if defineDocument.IsCertificate(docName) {
		return "Certidão"
	} else if defineDocument.IsProof(docName) {
		return "Elementos de prova"
	} else if defineDocument.IsInformation(docName) {
		return "Informação"
	} else if defineDocument.IsMemorial(docName) {
		return "Memoriais"
	} else if defineDocument.IsCounterarguments(docName) {
		return "Contrarrazões da Apelação"
	} else if defineDocument.IsReplica(docName) {
		return "Réplica"
	} else if defineDocument.IsPrepositonLetter(docName) {
		return "Pedido"
	} else {
		return docName
	}

}
