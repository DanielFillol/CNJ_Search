package CSV

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
)

// LegalDocumentCSV creates a csv with every single row with SearchCNJ
//from a given raw file and separator rune in a given folder
func LegalDocumentCSV(rawFilePath string, separator rune, nameResultFolder string, searchType int) error {
	raw, err := readCsvFile(rawFilePath, separator)
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

//execute the SearchCNJ from a []string
func createCSVs(raw []string, nameResultFolder string, searchType int) error {
	var final []Classifier.FinalData
	for _, search := range raw {
		data, err := Classifier.SearchCNJ(search, searchType)
		if err != nil {
			final = append(final, Classifier.FinalData{
				SearchName: search,
				CNJId:      err.Error(),
				CNJIdUpper: err.Error(),
				CNJName:    err.Error(),
				CnjReturn:  nil,
			})
		} else {
			final = append(final, data)
		}
	}

	err := writeCSV("filesOK", nameResultFolder, final)
	if err != nil {
		return err
	}

	return nil
}
