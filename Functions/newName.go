package Functions

import "github.com/Darklabel91/LegalDoc_Classifier/Functions/defineDocument"

func newName(docName string) string {
	var text string

	if defineDocument.IsInitialRequest(docName) {
		text = "Petição Inicial"
	} else if defineDocument.IsInitialRequestEmd(docName) {
		text = "Emenda à Inicial"
	} else if defineDocument.IsDiverseRequest(docName) {
		text = "Petição (Outras)"
	} else if defineDocument.IsDiverseMiddleRequest(docName) {
		text = "Petição Intermediária"
	} else if defineDocument.IsDocument(docName) {
		text = "Documentos Diversos"
	} else if defineDocument.IsQualificationRequest(docName) {
		text = "Pedido de Habilitação"
	} else if defineDocument.IsReplacementRequest(docName) {
		text = "Instrumento de Procuração"
	} else if defineDocument.IsLetterOfAttorney(docName) {
		text = "Instrumento de Procuração"
	} else if defineDocument.IsProofOfResidence(docName) {
		text = "Comprovante de Endereço"
	} else if defineDocument.IsAssetsToPledge(docName) {
		text = "Pedido de Penhora"
	} else if defineDocument.IsHearingRequest(docName) {
		text = "Pedido de Designação/Redesignação de Audiência"
	} else if defineDocument.IsRequirements(docName) {
		text = "Apresentação de Quesitos/Indicação de Assistente Técnico"
	} else if defineDocument.IsExpertFees(docName) {
		text = "Solicitação de Honorários de Perito"
	} else if defineDocument.IsExpertClarification(docName) {
		text = "Esclarecimento de Perito"
	} else if defineDocument.IsExpertsManifestation(docName) {
		text = "Manifestação do Perito"
	} else if defineDocument.IsExpertProof(docName) {
		text = "Agendamento de Vistoria - Prova Pericial - Art. 474 do CPC"
	} else if defineDocument.IsBacen(docName) {
		text = "Sisbajud"
	} else if defineDocument.IsCertificate(docName) {
		text = "Certidão"
	} else if defineDocument.IsProof(docName) {
		text = "Elementos de prova"
	} else if defineDocument.IsInformation(docName) {
		text = "Informação"
	} else if defineDocument.IsMemorial(docName) {
		text = "Memoriais"
	} else if defineDocument.IsCounterarguments(docName) {
		text = "Contrarrazões da Apelação"
	} else if defineDocument.IsReplica(docName) {
		text = "Réplica"
	} else if defineDocument.IsPrepositonLetter(docName) {
		text = "Pedido"
	} else {
		text = docName
	}
	return text
}
