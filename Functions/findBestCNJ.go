package Functions

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"strings"
)

func findBestCNJDocument(searchString string, dataCNJ []Structs.CNJArray) (string, string, string) {
	var cnjId string
	var cnjIdUpper string
	var cnjName string

	if len(dataCNJ) > 1 {
		for i := 0; i < len(dataCNJ); i++ {
			if strings.ToLower(dataCNJ[i].Name) == strings.ToLower(searchString) {
				cnjId = dataCNJ[i].IdItem
				cnjIdUpper = dataCNJ[i].IdItemUpper
				cnjName = dataCNJ[i].Name
			} else {
				cnjId = dataCNJ[0].IdItem
				cnjIdUpper = dataCNJ[0].IdItemUpper
				cnjName = dataCNJ[0].Name
			}
		}

	} else {

		cnjId = dataCNJ[0].IdItem
		cnjIdUpper = dataCNJ[0].IdItemUpper
		cnjName = dataCNJ[0].Name

	}

	return cnjId, cnjIdUpper, cnjName
}
