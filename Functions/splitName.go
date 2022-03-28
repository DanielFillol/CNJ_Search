package Functions

import (
	"github.com/Darklabel91/LegalDoc_Classifier/Functions/defineDocument"
	"strings"
)

func SplitName(docName string) string {
	var returnText string
	text := strings.Split(docName, "- ")

	if len(text) > 1 {
		text1 := strings.Split(text[1], ".")
		if len(text1) > 1 {
			returnText = newName(text1[0])
		} else {
			returnText = newName(text[1])
		}
	} else {
		returnText = newName(docName)
	}

	if defineDocument.IsPage(returnText) == true && defineDocument.IsInitialRequest(returnText) == false {
		returnText = newName(text[0])
	}

	return returnText
}
