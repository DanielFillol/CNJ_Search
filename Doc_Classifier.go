package LegalDoc_Classifier

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
	"github.com/Darklabel91/LegalDoc_Classifier/Database"
	"github.com/Darklabel91/LegalDoc_Classifier/Functions"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
)

func DocClassifier(identifier string, searchString string, searchType int) (Structs.FinalData, string) {
	var dataClassified Structs.FinalData
	var dataCNJ []Structs.DocumentCNJ
	var docName string
	var status string

	if searchType == 0 {
		dataCNJ = Database.DataSetCNJDocument()
		docName = Functions.SplitName(searchString)
	} else if searchType == 1 {
		dataCNJ = Database.DataSetCNJsubject()
		docName = searchString
	} else {
		panic("searchType must be 1 or 0")
	}

	cnjReturn, cnjId, cnjIdUpper, cnjFirst, flag := Functions.ReturnCNJ(docName, dataCNJ)

	if flag {
		status = "success"
	} else {
		status = "failed"
	}

	dataClassified = Structs.FinalData{
		Id:         identifier,
		SearchName: searchString,
		CNJId:      cnjId,
		CNJIdUpper: cnjIdUpper,
		CNJName:    cnjFirst,
		CnjReturn:  cnjReturn,
	}

	return dataClassified, status
}

func DocClassifierCSV(rawFilePath string, separator rune, nameResultFolder string, searchType int) error {
	raw, err := CSV.ReadCsvFile(rawFilePath, separator)
	if err != nil {
		return err
	}
	err = createCSVs(raw, nameResultFolder, searchType)
	if err != nil {
		return err
	}
	fmt.Println("Files created")
	return nil
}

func createCSVs(raw []Structs.RawDocument, nameResultFolder string, searchType int) error {
	var filesOK []Structs.FinalData
	var filesError []Structs.FinalData

	for i := 0; i < len(raw); i++ {
		emp, status := DocClassifier(raw[i].IdItem, raw[i].Name, searchType)
		if status == "success" {
			filesOK = append(filesOK, emp)
		} else {
			filesError = append(filesError, emp)
		}
	}

	if len(filesOK) != 0 {
		err := CSV.ExportCSV("filesOK", nameResultFolder, filesOK)
		if err != nil {
			return err
		}
	}

	if len(filesError) != 0 {
		err := CSV.ExportCSV("filesError", nameResultFolder, filesError)
		if err != nil {
			return err
		}
	}

	return nil
}
