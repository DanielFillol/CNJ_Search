package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/LegalDoc_Classifier/Error"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportCSV(nameFile string, nameFolder string, result []Structs.FinalData) {
	empData := [][]string{}

	empData = append(empData, []string{"ID", "DocName", "CNJid", "CNJidUpper", "CNJname", "CNJall"})

	for i := 0; i < len(result); i++ {
		var cnjReturn string
		dataCNJ := result[i].CnjReturn

		for j := 0; j < len(dataCNJ); j++ {
			cnjReturn += "{" + dataCNJ[j].IdItem + " ; " + dataCNJ[j].IdItemUpper + " ; " + dataCNJ[j].Name + "} "
		}

		final := []string{result[i].Id, result[i].SearchName, result[i].CNJId, result[i].CNJIdUpper, result[i].CNJName, cnjReturn}
		empData = append(empData, final)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")
	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}
