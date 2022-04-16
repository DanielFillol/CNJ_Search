package CSV

import (
	"encoding/csv"
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

func ExportCSV(nameFile string, nameFolder string, result []Structs.FinalData) error {
	var classDocs [][]string

	classDocs = append(classDocs, []string{"ID", "DocName", "CNJid", "CNJidUpper", "CNJname", "CNJall"})

	for i := 0; i < len(result); i++ {
		var cnjReturn string
		dataCNJ := result[i].CnjReturn

		for j := 0; j < len(dataCNJ); j++ {
			cnjReturn += "{" + dataCNJ[j].IdItem + " ; " + dataCNJ[j].IdItemUpper + " ; " + dataCNJ[j].Name + "} "
		}

		classDoc := []string{result[i].Id, result[i].SearchName, result[i].CNJId, result[i].CNJIdUpper, result[i].CNJName, cnjReturn}
		classDocs = append(classDocs, classDoc)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range classDocs {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()

	return nil
}
