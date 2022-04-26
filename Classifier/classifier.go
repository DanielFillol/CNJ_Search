package Classifier

import "strings"

type FinalData struct {
	SearchName string `json:"SearchName,omitempty"`
	CNJId      string `json:"CNJId,omitempty"`
	CNJIdUpper string `json:"CNJIdUpper,omitempty"`
	CNJName    string `json:"CNJName,omitempty"`
	CnjReturn  []CNJ  `json:"CnjReturn,omitempty"`
}

type FoundCNJ struct {
	BestCNJ CNJ   `json:"best_cnj,omitempty"`
	CNJs    []CNJ `json:"cnjs,omitempty"`
}

type CNJ struct {
	IdItem      string `json:"IdItem,omitempty"`
	IdItemUpper string `json:"IdItemUpper,omitempty"`
	Name        string `json:"Name,omitempty"`
}

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
