package Classifier

import (
	"errors"
	"github.com/Darklabel91/LegalDoc_Classifier/Database"
	"strings"
)

func getCNJ(normalizedDocName string) (FoundCNJ, error) {
	dataCNJ := Database.DataSetCNJDocument()

	var foundCNJ []CNJ
	for _, singleCNJ := range dataCNJ {
		if strings.Contains(strings.ToLower(normalizedDocName), strings.ToLower(singleCNJ.Name)) {
			foundCNJ = append(foundCNJ, CNJ{
				IdItem:      singleCNJ.IdItem,
				IdItemUpper: singleCNJ.IdItemUpper,
				Name:        singleCNJ.Name,
			})
		}
	}

	if len(foundCNJ) != 0 {
		bestCNJ := getBestCNJ(normalizedDocName, foundCNJ)
		return FoundCNJ{
			BestCNJ: bestCNJ,
			CNJs:    foundCNJ,
		}, nil
	} else {
		return FoundCNJ{}, errors.New("impossible to find a cnj code for this document name")
	}
}

func getBestCNJ(normalizedDocName string, foundCNJs []CNJ) CNJ {
	for _, cnj := range foundCNJs {
		if strings.ToLower(cnj.Name) == strings.ToLower(normalizedDocName) {
			return CNJ{
				IdItem:      cnj.IdItem,
				IdItemUpper: cnj.IdItemUpper,
				Name:        cnj.Name,
			}

		}
	}
	return foundCNJs[0]
}
