package Functions

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"strings"
)

func findBestCNJ(docNameSplit string, dataCNJ []Structs.CNJArray) (string, string, string) {
	var cnjId string
	var cnjIdUpper string
	var cnjName string

	if len(dataCNJ) > 1 {
		for i := 0; i < len(dataCNJ); i++ {
			if strings.ToLower(dataCNJ[i].Name) == strings.ToLower(docNameSplit) {

				cnjId = dataCNJ[i].Id_item
				cnjIdUpper = dataCNJ[i].Id_item_upper
				cnjName = dataCNJ[i].Name

			} else {

				if dataCNJ[i].Id_item_upper == "1" || dataCNJ[i].Id_item_upper == "3" || dataCNJ[i].Id_item_upper == "4" || dataCNJ[i].Id_item_upper == "5" || dataCNJ[i].Id_item_upper == "6" {

					cnjId = dataCNJ[i].Id_item
					cnjIdUpper = dataCNJ[i].Id_item_upper
					cnjName = dataCNJ[i].Name

				} else {

					cnjId = dataCNJ[0].Id_item
					cnjIdUpper = dataCNJ[0].Id_item_upper
					cnjName = dataCNJ[0].Name

				}
			}
		}

	} else {

		cnjId = dataCNJ[0].Id_item
		cnjIdUpper = dataCNJ[0].Id_item_upper
		cnjName = dataCNJ[0].Name

	}

	return cnjId, cnjIdUpper, cnjName
}
