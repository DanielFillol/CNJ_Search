package Classifier

import "strings"

func getMatchWord(docName string, words []string) bool {
	for _, word := range words {
		if strings.Contains(strings.ToLower(docName), word) {
			return true
		}
	}
	return false
}
