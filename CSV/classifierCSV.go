package CSV

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
)

// LegalDocumentCSV creates a csv with every single row with LegalDocument
//from a given raw file and separator rune in a given folder
func LegalDocumentCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := readCsvFile(rawFilePath, separator)
	if err != nil {
		return err
	}

	err = createCSVs(raw, nameResultFolder)
	if err != nil {
		return err
	}

	fmt.Println("Files created")
	return nil
}

//execute the AnalyzeCNJ from a []string
func createCSVs(raw []string, nameResultFolder string) error {
	var final []Classifier.FinalData
	for i := 0; i < len(raw); i++ {
		emp, _ := Classifier.LegalDocument(raw[i])
		final = append(final, emp)
	}

	err := writeCSV("filesOK", nameResultFolder, final)
	if err != nil {
		return err
	}

	return nil
}
