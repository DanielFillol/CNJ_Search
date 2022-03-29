package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/LegalDoc_Classifier/Error"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"os"
)

func ReadCsvFile(filePath string, separator rune) []Structs.RawDocument {
	var data2 []Structs.RawDocument

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer func(csvFile *os.File) {
		err0 := csvFile.Close()
		Error.CheckError(err0)
	}(csvFile)

	csvLines := csv.NewReader(csvFile)
	csvLines.Comma = separator
	csvData, err1 := csvLines.ReadAll()
	Error.CheckError(err1)

	for _, line := range csvData {
		emp := Structs.RawDocument{
			IdItem: line[0],
			Name:   line[1],
		}
		data2 = append(data2, emp)
	}
	return data2
}
