package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"os"
)

func ReadCsvFile(filePath string, separator rune) ([]Structs.RawDocument, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return []Structs.RawDocument{}, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return []Structs.RawDocument{}, err
	}

	var data []Structs.RawDocument

	for _, line := range csvData {
		newData := Structs.RawDocument{
			IdItem: line[0],
			Name:   line[1],
		}
		data = append(data, newData)
	}

	return data, nil
}
