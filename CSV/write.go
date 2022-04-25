package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
	"os"
	"path/filepath"
	"strings"
)

//writeCSV exports a csv to a given folder, with a given name from a collection of LegalDocument
func writeCSV(fileName string, folderName string, cnjRows []Classifier.FinalData) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, cnj := range cnjRows {
		rows = append(rows, generateRow(cnj))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Termo Pesquisado",
		"Melhor CNJ ID",
		"Melhor CNJ ID_PAI",
		"Melhor CNJ Nome",
		"Lista CNJ's Encontrados",
	}
}

// returns a []string that compose the row in the csv file
func generateRow(cnj Classifier.FinalData) []string {
	var cnjReturn string
	for _, singleCNJ := range cnj.CnjReturn {
		cnjReturn += "{" + singleCNJ.IdItem + " ; " + singleCNJ.IdItemUpper + " ; " + singleCNJ.Name + "} "
	}

	return []string{
		cnj.SearchName,
		cnj.CNJId,
		cnj.CNJIdUpper,
		cnj.CNJName,
		strings.TrimSpace(cnjReturn),
	}
}
