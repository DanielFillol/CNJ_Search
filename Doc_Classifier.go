package LegalDoc_Classifier

import (
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
	"github.com/Darklabel91/LegalDoc_Classifier/Functions"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
)

const pathCNJ = "CNJ_doc/DataCNJ.csv"

func DocClassifier(identifier string, fileName string) (Structs.FinalData, string) {
	var dataClassified Structs.FinalData
	var status string

	dataCNJ := CSV.ReadCSVcnj(pathCNJ)
	docName := Functions.SplitName(fileName)
	cnjReturn, cnjId, cnjIdUpper, cnjFirst, flag := Functions.ReturnCNJ(docName, dataCNJ)

	if flag {
		status = "success"
	} else {
		status = "failed"
	}

	dataClassified = Structs.FinalData{
		Id:           identifier,
		Doc_name:     fileName,
		CNJ_id:       cnjId,
		CNJ_id_upper: cnjIdUpper,
		CNJname:      cnjFirst,
		CnjReturn:    cnjReturn,
	}

	return dataClassified, status
}
