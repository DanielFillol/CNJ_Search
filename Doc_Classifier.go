package LegalDoc_Classifier

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
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

func DocClassifierCSV(rawFilePath string, separator rune, nameResultFolder string) {
	raw := CSV.ReadCsvFile(rawFilePath, separator)
	createCSVs(raw, nameResultFolder)
	fmt.Println("Files created")
}

func createCSVs(raw []Structs.RawDocument, nameResultFolder string) {
	var filesOK []Structs.FinalData
	var filesError []Structs.FinalData

	for i := 0; i < len(raw); i++ {
		emp, status := DocClassifier(raw[i].IdItem, raw[i].Name)
		if status == "success" {
			filesOK = append(filesOK, emp)
		} else {
			filesError = append(filesError, emp)
		}
	}

	if len(filesOK) != 0 {
		CSV.ExportCSV("filesOK", nameResultFolder, filesOK)
	}

	if len(filesError) != 0 {
		CSV.ExportCSV("filesError", nameResultFolder, filesError)
	}
}
