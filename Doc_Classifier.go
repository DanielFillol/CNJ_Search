package main

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Database"
	"github.com/Darklabel91/LegalDoc_Classifier/Functions"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
)

func DocClassifier(identifier string, fileName string) (Structs.FinalData, string) {
	var dataClassified Structs.FinalData
	var status string

	dataCNJ := Database.DataSetCNJ()
	docName := Functions.SplitName(fileName)
	cnjReturn, cnjId, cnjIdUpper, cnjFirst, flag := Functions.ReturnCNJ(docName, dataCNJ)

	if flag {
		status = "success"
	} else {
		status = "failed"
	}

	dataClassified = Structs.FinalData{
		Id:         identifier,
		DocName:    fileName,
		CNJId:      cnjId,
		CNJIdUpper: cnjIdUpper,
		CNJName:    cnjFirst,
		CnjReturn:  cnjReturn,
	}

	return dataClassified, status
}
