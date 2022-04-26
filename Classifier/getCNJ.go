package Classifier

import (
	"errors"
	"github.com/Darklabel91/LegalDoc_Classifier/Database"
	"strings"
)

//getCNJDocument returns FoundCNJ for a search string.
//  searchType must be 0 or 1
//	0: searches the document cnj table
//  1: searches the subject cnj table
func getCNJDocument(normalizedSearchName string, searchType int) (FoundCNJ, error) {

	var dataCNJ []Database.CNJ
	if searchType != 1 {
		dataCNJ = Database.DataSetCNJDocument()
	} else {
		dataCNJ = Database.DataSetCNJsubject()
	}

	var foundCNJ []CNJ
	for _, singleCNJ := range dataCNJ {
		if strings.Contains(strings.ToLower(normalizedSearchName), strings.ToLower(singleCNJ.Name)) {
			foundCNJ = append(foundCNJ, CNJ{
				IdItem:      singleCNJ.IdItem,
				IdItemUpper: singleCNJ.IdItemUpper,
				Name:        singleCNJ.Name,
			})
		}
	}

	if len(foundCNJ) != 0 {
		bestCNJ := getBestCNJ(normalizedSearchName, foundCNJ)
		return FoundCNJ{
			BestCNJ: bestCNJ,
			CNJs:    foundCNJ,
		}, nil
	} else {
		return FoundCNJ{}, errors.New("impossible to find a cnj code for this document name")
	}
}

//getBestCNJ returns a single CNJ with the best match for the search string
func getBestCNJ(normalizedSearchName string, foundCNJs []CNJ) CNJ {
	for _, cnj := range foundCNJs {
		if strings.ToLower(cnj.Name) == strings.ToLower(normalizedSearchName) {
			return CNJ{
				IdItem:      cnj.IdItem,
				IdItemUpper: cnj.IdItemUpper,
				Name:        cnj.Name,
			}

		}
	}
	return foundCNJs[0]
}
