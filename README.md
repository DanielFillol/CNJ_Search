# LegalDoc_Classifier
Projeto pensado para classificar tipos de documentos dos Tribunais de Justiça brasileiro conforme a padronização CNJ de classificação. Além disso também permite a pesquisa de qualquer texto com a tabela de assuntos na padronização CNJ.

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


### searchType
searchType é um número inteiro passado para as funções *main* com o objetivo de selecionar qual tabela será pesquisada sendo:
- 0 -> tabela CNJ de Documentos
- 1 -> tabela CNJ de Assuntos

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
)

func main() {

	//Search document name on CNJ document table
	documentName := "7b00616 - Recurso de Revista.pdf"

	docClass, err := Classifier.SearchCNJ(documentName, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(docClass)

	//Search subject name on CNJ subject table
	subjectSearch := "plano de saúde"

	subjClass, err := Classifier.SearchCNJ(subjectSearch, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(subjClass)

	//CSV search document example
	raw := "CSV/testDocuments.csv"
	separator := ','
	resultFolder := "Test"

	err = CSV.LegalDocumentCSV(raw, separator, resultFolder, 0)
	if err != nil {
		fmt.Println(err)
	}
}
 ```
Retorno
``` 
{7b00616 - Recurso de Revista.pdf 233 48 Recurso de Revista [{48 3 Recurso} {233 48 Recurso de Revista}]}

{plano de saúde 2364 2581 Plano de Saúde [{2364 2581 Plano de Saúde} {10064 10028 Saúde}]}

Files created

 ```

## Functions

### Main Function:
- SearchCNJ(searchString string, searchType int)  ->  retorna uma *FinalData* necessitantando de um identificador, do nome do arquivo a ser classificado e do tipo *searchType*
- LegalDocumentCSV(rawFilePath string, separator rune, nameResultFolder string, searchType int)  -> retorna dois arquivos .CSV
 necessitando apenas do caminho do arquivo a ser analisado, o separador (';' ',' etc..) de colunas, o nome da pasta em que os resultados devem ser salvos e do tipo *searchType* O arquivo a ser analisado deve ter duas colunas {id, fileName/SearchString}
 
### Define Docment Functions:
- IsAssetsToPledge(docName string)       ->  retorna true para um documento *Pedido de Penhora*
- IsBacen(docName string)                ->  retorna true para um documento *Sisbajud*
- IsCertificate(docName string)          ->  retorna true para um documento *Certidão*
- IsCounterarguments(docName string)     ->  retorna true para um documento *Contrarrazões da Apelação*
- IsDiverseMiddleRequest(docName string) ->  retorna true para um documento *Petição Intermediária*
- IsDiverseRequest(docName string)       ->  retorna true para um documento *Petição (Outras)* 
- IsDocument(docName string)             ->  retorna true para um documento *Documentos Diversos*
- IsExpertClarification(docName string)  ->  retorna true para um documento *Esclarecimento de Perito*
- IsExpertFees(docName string)           ->  retorna true para um documento *Solicitação de Honorários de Perito* 
- IsExpertProof(docName string)          ->  retorna true para um documento *Agendamento de Vistoria - Prova Pericial - Art. 474 do CPC*
- IsExpertsManifestation(docName string) ->  retorna true para um documento *Manifestação do Perito*
- IsHearingRequest(docName string)       ->  retorna true para um documento *Pedido de Designação/Redesignação de Audiência*
- IsInformation(docName string)          ->  retorna true para um documento *Informação* 
- IsInitialRequest(docName string)       ->  retorna true para um documento *Petição Inicial*
- IsInitialRequestEmd(docName string)    ->  retorna true para um documento *Emenda à Inicial*
- IsLetterOfAttorney(docName string)     ->  retorna true para um documento *Instrumento de Procuração*
- IsMemorial(docName string)             ->  retorna true para um documento *Memoriais*
- IsPrepositonLetter(docName string)     ->  retorna true para um documento *Pedido*
- IsProof(docName string)                ->  retorna true para um documento *Elementos de prova*
- IsProofOfResidence(docName string)     ->  retorna true para um documento *Comprovante de Endereço*
- IsQualificationRequest(docName string) ->  retorna true para um documento *Pedido de Habilitação*
- IsReplacementRequest(docName string)   ->  retorna true para um documento *Instrumento de Procuração*
- IsReplica(docName string)              ->  retorna true para um documento *Réplica*
- IsRequirements(docName string)         ->  retorna true para um documento *Apresentação de Quesitos/Indicação de Assistente Técnico*


## Disclaimer
Esse classificador foi testado e apresenta, até o momento, 92% de assertividade, ainda assim, use com cautela.

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
