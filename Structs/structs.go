package Structs

type DocumentCNJ struct {
	IdItem      string `json:"IdItem,omitempty"`
	IdItemUpper string `json:"IdItemUpper,omitempty"`
	Name        string `json:"Name,omitempty"`
}

type RawDocument struct {
	IdItem string `json:"IdItem,omitempty"`
	Name   string `json:"Name,omitempty"`
}

type FinalData struct {
	Id         string     `json:"Id,omitempty"`
	SearchName string     `json:"SearchName,omitempty"`
	CNJId      string     `json:"CNJId,omitempty"`
	CNJIdUpper string     `json:"CNJIdUpper,omitempty"`
	CNJName    string     `json:"CNJName,omitempty"`
	CnjReturn  []CNJArray `json:"CnjReturn,omitempty"`
}

type CNJArray struct {
	SearchName  string `json:"SearchName,omitempty"`
	IdItem      string `json:"IdItem,omitempty"`
	IdItemUpper string `json:"IdItemUpper,omitempty"`
	Name        string `json:"Name,omitempty"`
}
