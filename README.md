# LegalDoc_Classifier
Projeto pensado para classificar tipos de documentos dos Tribunais de Justiça brasileiro conforme a padronização CNJ de classificação. Além disso também permite a pesquisa de qualquer texto com a tabela de assuntos na padronização CNJ.

## Instal
``` go get -u github.com/Darklabel91/LegalDoc_Classifier ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string```, ```FinalData``` ou ```CNJArray```, essas últimas são compostas por:

``` 
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

```
### FinalData
- Id: Um identificador. É recomendado que seja associado um número de processo no formato CNJ para facilitar o cruzamento de dados
- DocName: Nome do arquivo para ser classificado conforme o CNJ
- CNJId: Identificador do Código CNJ específico
- CNJIdUpper: Identificador do Código Pai do CNJ específico
- CNJName: Classificação do CNJ
- CnjReturn: Conjunto de dados com a estrutura ```CNJArray``` com todos os códigos CNJ encontrados

### CNJArray
- DocName: Nome do arquivo para ser classificado conforme o CNJ
- IdItem: Identificador do Código CNJ específico
- IdItemUpper: Identificador do Código Pai do CNJ específico
- Name: Classificação do CNJ


### searchType
searchType é um número inteiro passado para as funções *main* com o objetivo de selecionar qual tabela será pesquisada sendo:
- 0 -> tabela CNJ de Documentos
- 1 0> tabela CNJ de Assuntos

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/LegalDoc_Classifier"
)

func main() {
	// searchType: 0 for Document search and 1 for subject search
	var searchType int 
	
	id := "0"
	fileName := "Petição Inicial.pdf"
	searchType = 0

	test, status := LegalDoc_Classifier.DocClassifier(id, fileName, searchType)

	fmt.Println(status)
	fmt.Println(test.Id, test.SearchName, test.CNJId, test.CNJIdUpper, test.CNJName)
	fmt.Println(test.CnjReturn)

	//READING A CSV WITH FILE NAMES

	rawPath := "/Users/Desktop/tjFiles.csv"
	separator := ','
	resultFolder := "Result"
	searchType = 0

	LegalDoc_Classifier.DocClassifierCSV(rawPath, separator, resultFolder, searchType)
}
 ```
Retorno
``` 
success
0 Petição Inicial.pdf 37 3 Petição
[{Petição Inicial.pdf 37 3 Petição} {Petição Inicial.pdf 202 37 Petição Inicial}]
Files created

 ```

## Functions

### Main Function:
- DocClassifier(identifier string, fileName string, searchType int)  ->  retorna uma *FinalData* necessitantando de um identificador, do nome do arquivo a ser classificado e do tipo *searchType*
- DocClassifierCSV(rawFilePath string, separator rune, nameResultFolder string, searchType int) -> retorna dois arquivos .CSV
 necessitando apenas do caminho do arquivo a ser analisado, o separador (';' ',' etc..) de colunas, o nome da pasta em que os resultados devem ser salvos e do tipo *searchType*
 O arquivo a ser analisado deve ter duas colunas {id, fileName/SearchString}
 
### Suport Functions:
- newName(docName string) -> retorna o nome normalizado como *string*. Para ser efetiva essa função faz uso de outras 24 funções que retornam *bool* para cada tipo de documento mapeado (apontadas abaixo)
- SplitName(docName string) -> retorna uma *string* como o nome normalizado que deve ser pesquisado na tabela CNJ. Para normalizar o nome é utilizada a função acima.
- findBestCNJ(docNameSplit string, dataCNJ []Structs.CNJArray) -> retorna *cnjId*, *cnjIdUppe* e *cnjName*, necessitando do nome do arquivo a ser classificado (já normalizado) e uma *CNJArray*
- ReturnCNJ(docNameSplit string, data []Structs.DocumentCNJ) -> retorna uma *CNJArray*, *BestCNJ*(faz uso da função mencionada acima) e *status*, necessitantando do nome do arquivo a ser classificado e o conjunto de dados do CNJ. 

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
