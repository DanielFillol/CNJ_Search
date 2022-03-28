package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/LegalDoc_Classifier/Error"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"os"
)

func ReadCSVcnj(filePath string) []Structs.DocumentCNJ {
	var data []Structs.DocumentCNJ

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer func(csvFile *os.File) {
		err0 := csvFile.Close()
		Error.CheckError(err0)
	}(csvFile)

	csvLines, err1 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err1)

	for _, line := range csvLines {
		emp := Structs.DocumentCNJ{
			Id_item:       line[0],
			Id_item_upper: line[1],
			Name:          line[3],
		}
		data = append(data, emp)
	}

	return data
}
