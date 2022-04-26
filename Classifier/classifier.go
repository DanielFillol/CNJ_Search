package Classifier

import "strings"

//FinalData the main project struct, contains the search string and FoundCNJ.
type FinalData struct {
	SearchName string `json:"SearchName,omitempty"`
	CNJId      string `json:"CNJId,omitempty"`
	CNJIdUpper string `json:"CNJIdUpper,omitempty"`
	CNJName    string `json:"CNJName,omitempty"`
	CnjReturn  []CNJ  `json:"CnjReturn,omitempty"`
}

//FoundCNJ contains the best found CNJ and the [] of all found CNJ
type FoundCNJ struct {
	BestCNJ CNJ   `json:"best_cnj,omitempty"`
	CNJs    []CNJ `json:"cnjs,omitempty"`
}

//CNJ struct that reflects the CNJ document and subject database
type CNJ struct {
	IdItem      string `json:"IdItem,omitempty"`
	IdItemUpper string `json:"IdItemUpper,omitempty"`
	Name        string `json:"Name,omitempty"`
}

//SearchCNJ mains project function, returns FinalData,
//need's string to search and int (0 or 1):
//  0: search's the cnj document database
//  1: search's the cnj subject database
func SearchCNJ(searchString string, searchType int) (FinalData, error) {
	nlzDocName := normalizeDocName(strings.TrimSpace(searchString))

	foundCnj, err := getCNJDocument(nlzDocName, searchType)
	if err != nil {
		return FinalData{}, err
	}

	return FinalData{
		SearchName: searchString,
		CNJId:      foundCnj.BestCNJ.IdItem,
		CNJIdUpper: foundCnj.BestCNJ.IdItemUpper,
		CNJName:    foundCnj.BestCNJ.Name,
		CnjReturn:  foundCnj.CNJs,
	}, nil

}
