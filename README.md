# CNJ_Search
Projeto nasceu da proposta de encontar uma classificação CNJ para os arquivos em pdf protocolados nos tribunais de justiça de todo o país. O projeto evoluíou e agora é possível realizar a busca por qualquer termo dentro da base de dados do CNJ.

## Instal
``` go get -u github.com/Darklabel91/LegalDoc_Classifier ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string```, ```FinalData``` ou ```CNJ```, essas últimas são compostas por:

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
- SearchName: Nome do arquivo para ser classificado conforme o CNJ
- CNJId: Identificador do Código CNJ específico
- CNJIdUpper: Identificador do Código Pai do CNJ específico
- CNJName: Classificação do CNJ
- CnjReturn: Conjunto de dados com a estrutura ```CNJ``` com todos os códigos CNJ encontrados

### CNJ
- IdItem: Identificador do Código CNJ específico
- IdItemUpper: Identificador do Código Pai do CNJ específico
- Name: Classificação do CNJ


### databaseType
É a chave de pesquisa em formato *rune* entre os 4 bancos CNJ possíveis sendo:
- 'D' acessa o banco de Documentos
- 'M' acessa o banco de Movimentos
- 'C' acessa o banco de Classes
- 'A' acessa o banco de Assuntos

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

	//--//Search movement name on CNJ movement table
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
- SearchCNJ(searchString string, databaseType rune)  ->  retorna uma *FinalData* necessitantando da chave de pesquisa tipo rune *databaseType*

- SearchCNJCSV(rawFilePath string, separator rune, nameResultFolder string, databaseType rune) -> retorna um csv necessitando do caminho do arquivo a ser analisado, o separador de colunas (';' ',' etc..), o nome da pasta em que os resultados devem ser salvos e o *databaseType*. O csv a ser analisado deve conter apenas uma coluna, sem títulos.


## Disclaimer
Esse classificador foi testado e apresenta, até o momento, 92% de assertividade, ainda assim, use com cautela.

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
