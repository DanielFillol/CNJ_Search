package Functions

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"strings"
)

func ReturnCNJ(docNameSplit string, data []Structs.DocumentCNJ) ([]Structs.CNJArray, string, string, string, bool) {

	var dataCNJ []Structs.CNJArray
	for _, singleData := range data {
		if strings.Contains(strings.ToLower(docNameSplit), strings.ToLower(singleData.Name)) {

			dataCNJ = append(dataCNJ, Structs.CNJArray{
				SearchName:  docNameSplit,
				IdItem:      singleData.IdItem,
				IdItemUpper: singleData.IdItemUpper,
				Name:        singleData.Name,
			})

			cnjId, cnjIdUpper, cnjName := findBestCNJDocument(docNameSplit, dataCNJ)

			return dataCNJ, cnjId, cnjIdUpper, cnjName, true
		}
	}

	return []Structs.CNJArray{}, "", "", "", false
}
