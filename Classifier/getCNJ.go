package Classifier

import (
	"github.com/Darklabel91/LegalDoc_Classifier/ClassifierDatabase"
	"strings"
)

//getCNJDocument returns FoundCNJ for a search string. returns error when no CNJ is founded in the search
func getCNJDocument(normalizedSearchName string, databaseType rune) (FoundCNJ, error) {
	cnjFetched, err := ClassifierDatabase.FetchCNJ(normalizedSearchName, databaseType)
	if err != nil {
		return FoundCNJ{}, err
	}

	var foundCNJ []CNJ
	for _, cnj := range cnjFetched {
		foundCNJ = append(foundCNJ, CNJ{IdItem: cnj.IdItem, IdItemUpper: cnj.IdItemUpper, Name: cnj.Name})
	}

	bestCNJ := getBestCNJ(normalizedSearchName, foundCNJ)

	return FoundCNJ{
		BestCNJ: bestCNJ,
		CNJs:    foundCNJ,
	}, nil
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
