package Functions

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"strings"
)

func ReturnCNJ(docNameSplit string, data []Structs.DocumentCNJ) ([]Structs.CNJArray, string, string, string, bool) {
	var dataCNJ []Structs.CNJArray
	var cnjId string
	var cnjIdUpper string
	var cnjName string
	flag := false

	for i := 0; i < len(data); i++ {
		if strings.Contains(strings.ToLower(docNameSplit), strings.ToLower(data[i].Name)) {

			dataCNJ = append(dataCNJ, Structs.CNJArray{
				DocName:     docNameSplit,
				IdItem:      data[i].IdItem,
				IdItemUpper: data[i].IdItemUpper,
				Name:        data[i].Name,
			})

			flag = true
		}
	}
	if flag == true {
		cnjId, cnjIdUpper, cnjName = findBestCNJ(docNameSplit, dataCNJ)
	} else {
		cnjId, cnjIdUpper, cnjName = "", "", ""
	}

	return dataCNJ, cnjId, cnjIdUpper, cnjName, flag
}
