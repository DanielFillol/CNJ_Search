package Structs

type DocumentCNJ struct {
	IdItem      string
	IdItemUpper string
	Name        string
}

type FinalData struct {
	Id         string
	DocName    string
	CNJId      string
	CNJIdUpper string
	CNJName    string
	CnjReturn  []CNJArray
}

type CNJArray struct {
	DocName     string
	IdItem      string
	IdItemUpper string
	Name        string
}
