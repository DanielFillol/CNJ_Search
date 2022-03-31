package Structs

type DocumentCNJ struct {
	IdItem      string
	IdItemUpper string
	Name        string
}

type RawDocument struct {
	IdItem string
	Name   string
}

type FinalData struct {
	Id         string
	SearchName string
	CNJId      string
	CNJIdUpper string
	CNJName    string
	CnjReturn  []CNJArray
}

type CNJArray struct {
	SearchName  string
	IdItem      string
	IdItemUpper string
	Name        string
}
