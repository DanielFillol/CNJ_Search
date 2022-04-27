package ClassifierDatabase

import (
	"errors"
	"strings"
)

//FetchCNJ returns []CNJData searching on docCNJDatabase or subjectCNJDatabase
func FetchCNJ(normalizedSearchName string, databaseType rune) ([]CNJData, error) {
	err := verifyRune(databaseType)
	if err != nil {
		return nil, err
	}

	var limitData []CNJData
	for _, cnj := range cnjDatabase {
		if cnj.Type == databaseType {
			limitData = append(limitData, cnj)
		}
	}

	var cnjMatches []CNJData
	for _, cnj := range limitData {
		if strings.Contains(strings.ToLower(normalizedSearchName), strings.ToLower(cnj.Name)) {
			cnjMatches = append(cnjMatches, cnj)
		}

	}

	if len(cnjMatches) != 0 {
		return cnjMatches, nil
	} else {
		return nil, errors.New("no cnj's where found for this normalizedSearchName: " + normalizedSearchName)
	}
}

//verify return error if the rune is not allowed
func verifyRune(databaseType rune) error {
	allowed := []rune{'D', 'M', 'C', 'A'}

	for _, r := range allowed {
		if databaseType == r {
			return nil
		}
	}

	return errors.New("this rune is invalid: must be 'D', 'M', 'C' or 'A'")

}
