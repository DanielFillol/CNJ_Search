package Classifier

import (
	"strings"
)

const (
	InitialRequest       = "Petição Inicial"
	InitialRequestEmd    = "Emenda à Inicial"
	DiverseRequest       = "Petição (Outras)"
	DiverseMiddleRequest = "Petição Intermediária"
	Document             = "Documentos Diversos"
	QualificationRequest = "Pedido de Habilitação"
	ReplacementRequest   = "Instrumento de Procuração"
	ProofOfResidence     = "Comprovante de Endereço"
	AssetsToPledge       = "Pedido de Penhora"
	HearingRequest       = "Pedido de Designação/Redesignação de Audiência"
	Requirements         = "Apresentação de Quesitos/Indicação de Assistente Técnico"
	ExpertFees           = "Solicitação de Honorários de Perito"
	ExpertClarification  = "Esclarecimento de Perito"
	ExpertsManifestation = "Manifestação do Perito"
	ExpertProof          = "Agendamento de Vistoria - Prova Pericial - Art. 474 do CPC"
	Bacen                = "Sisbajud"
	Certificate          = "Certidão"
	Proof                = "Elementos de prova"
	Information          = "Informação"
	Memorial             = "Memoriais"
	Counterarguments     = "Contrarrazões da Apelação"
	Replica              = "Réplica"
	PrepositonLetter     = "Pedido"
)

//Main normalizer function using getNormalizeName and pageAndNotInitial
func normalizeDocName(docName string) string {
	nmlzName := getNormalizeName(docName)

	if pageAndNotInitial(nmlzName) {
		splName := strings.Split(docName, "-")
		return getNormalizeName(strings.TrimSpace(splName[0]))
	} else {
		return nmlzName
	}
}

//verify if document name is a Page and not an Initial Requirement
func pageAndNotInitial(spltName string) bool {
	if isPage(spltName) == true && isInitialRequest(spltName) == false {
		return true
	}
	return false
}

//Split the doc name and normalize the file using nameNormalized
func getNormalizeName(name string) string {
	splName := strings.Split(name, "-")

	if len(splName) > 1 {
		splText := strings.Split(strings.TrimSpace(splName[1]), ".")

		if len(splText) >= 1 {
			return strings.TrimSpace(nameNormalized(splText[0]))
		} else {
			return strings.TrimSpace(nameNormalized(splText[1]))
		}
	} else {
		return strings.TrimSpace(nameNormalized(name))
	}
}

//return 1 of 23 possible normalized document names
func nameNormalized(docName string) string {
	if isInitialRequest(docName) {
		return InitialRequest
	}
	if isInitialRequestEmd(docName) {
		return InitialRequestEmd
	}
	if isDiverseRequest(docName) {
		return DiverseRequest
	}
	if isDiverseMiddleRequest(docName) {
		return DiverseMiddleRequest
	}
	if isDocument(docName) {
		return Document
	}
	if isQualificationRequest(docName) {
		return QualificationRequest
	}
	if isReplacementRequest(docName) {
		return ReplacementRequest
	}
	if isLetterOfAttorney(docName) {
		return ReplacementRequest
	}
	if isProofOfResidence(docName) {
		return ProofOfResidence
	}
	if isAssetsToPledge(docName) {
		return AssetsToPledge
	}
	if isHearingRequest(docName) {
		return HearingRequest
	}
	if isRequirements(docName) {
		return Requirements
	}
	if isExpertFees(docName) {
		return ExpertFees
	}
	if isExpertClarification(docName) {
		return ExpertClarification
	}
	if isExpertsManifestation(docName) {
		return ExpertsManifestation
	}
	if isExpertProof(docName) {
		return ExpertProof
	}
	if isBacen(docName) {
		return Bacen
	}
	if isCertificate(docName) {
		return Certificate
	}
	if isProof(docName) {
		return Proof
	}
	if isInformation(docName) {
		return Information
	}
	if isMemorial(docName) {
		return Memorial
	}
	if isCounterarguments(docName) {
		return Counterarguments
	}
	if isReplica(docName) {
		return Replica
	}
	if isPrepositonLetter(docName) {
		return PrepositonLetter
	}

	return docName
}
