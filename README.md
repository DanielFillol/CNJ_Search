# CNJ_Search
Package for searching matching string's in [CNJ](https://www.cnj.jus.br/sgt/consulta_publica_assuntos.php) database.

## Instal
``` go get -u github.com/Darklabel91/LegalDoc_Classifier ```

## Data Struct
Return data can be: ```bool```, ```string```, ```FinalData``` or ```CNJ```, the last two ar compose as:
``` 
type FinalData struct {
	SearchName string `json:"SearchName,omitempty"`
	CNJId      string `json:"CNJId,omitempty"`
	CNJIdUpper string `json:"CNJIdUpper,omitempty"`
	CNJName    string `json:"CNJName,omitempty"`
	CnjReturn  []CNJ  `json:"CnjReturn,omitempty"`
}


type CNJ struct {
	IdItem      string `json:"IdItem,omitempty"`
	IdItemUpper string `json:"IdItemUpper,omitempty"`
	Name        string `json:"Name,omitempty"`
}

```
### FinalData
- SearchName: search string to be searched in CNJ database
- CNJId: matched id on CNJ database (the best the package could find or the only)
- CNJIdUpper: upper id, if exists, of the initial id give above (the best the package could find or the only)
- CNJName: matched CNJ data (the best the package could find or the only)
- CnjReturn: all matched CNJ data found [CNJId,CNJIdUpper, CNJName]

### CNJ
- IdItem: id on CNJ database
- IdItemUpper: upper id, if exists, of the initial id give above 
- Name: matched CNJ data


### DatabaseType
Key *rune* for identifing one of four CNJ databases:
- 'D' documents database (documentos)
- 'M' movements database (movimentos)
- 'C' classes database (classes)
- 'A' subjects database (assuntos)

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
	"github.com/Darklabel91/LegalDoc_Classifier/ClassifierCSV"
)

func main() {

	//Search document name on CNJ document table
	search := "7b00616 - Recurso de Revista.pdf"

	found, err := Classifier.SearchCNJ(search, 'D')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(found)

	//Search movement name on CNJ movement table
	search = "prisão"

	found, err = Classifier.SearchCNJ(search, 'M')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(found)

	//Search Class name on CNJ movement table
	search = "Embargos à Arrematação"

	found, err = Classifier.SearchCNJ(search, 'C')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(found)

	//Search subject name on CNJ subject table
	search = "plano de saúde"

	found, err = Classifier.SearchCNJ(search, 'A')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(found)

	//CSV search document example
	raw := "CSV/testDocuments.csv"
	separator := ','
	resultFolder := "Test"

	err = ClassifierCSV.LegalDocumentCSV(raw, separator, resultFolder, 0)
	if err != nil {
		fmt.Println(err)
	}
}

 ```
Retorno
``` 
{7b00616 - Recurso de Revista.pdf 233 48 Recurso de Revista [{48 3 Recurso} {233 48 Recurso de Revista}]}

{prisão 128 157 Prisão [{128 157 Prisão}]}

{Embargos à Arrematação 171 169 Embargos à Arrematação [{207 197 Embargos} {1006 1071 Embargos} {169 158 Embargos} {171 169 Embargos à Arrematação}]}

{plano de saúde 2364 2581 Plano de Saúde [{2364 2581 Plano de Saúde} {10064 10028 Saúde} {13853 13831 Plano de Saúde}]}

files created.

 ```

## Functions

### Main Function:
- SearchCNJ(searchString string, databaseType rune)  ->  return *FinalData* for a given search in *searchString* on a specific database given in *databaseType* as rune.

- SearchCNJCSV(rawFilePath string, separator rune, nameResultFolder string, databaseType rune) -> return CSV file in given folder on *nameResultFolder*. The .csv passed in *rawFilePath* must contain a single column.



## Disclaimer
The precision was verifyed as 92%
