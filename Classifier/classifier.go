package Classifier

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

func LegalDocument(docName string) (FinalData, error) {

	nlzDocName := normalizeDocName(docName)

	foundCnj, err := getCNJ(nlzDocName)
	if err != nil {
		return FinalData{}, err
	}

	return FinalData{
		SearchName: docName,
		CNJId:      foundCnj.BestCNJ.IdItem,
		CNJIdUpper: foundCnj.BestCNJ.IdItemUpper,
		CNJName:    foundCnj.BestCNJ.Name,
		CnjReturn:  foundCnj.CNJs,
	}, nil

}
