package Structs

type DocumentJus struct {
	Id       string
	Doc_name string
	URL      string
}

type DocumentCNJ struct {
	Id_item       string
	Id_item_upper string
	Name          string
}

type CNJArray struct {
	Doc_name      string
	Id_item       string
	Id_item_upper string
	Name          string
}

type FinalData struct {
	Id           string
	Doc_name     string
	CNJ_id       string
	CNJ_id_upper string
	CNJname      string
	URL          string
	CnjReturn    []CNJArray
}

type NonClass struct {
	Id          string
	Doc_name    string
	SplitString string
	URL         string
}
