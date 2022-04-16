package defineDocument

import (
	"strings"
)

func IsAssetsToPledge(docName string) bool {
	example := []string{"indicação de bens à penhora",
		"penhora indicação de bens",
		"penhora  indicação de bens",
		"outros bens ou direitos",
		"nomeação de bens à penhora",
		"bens",
	}

	for _, total := range example {
		if strings.Contains(strings.ToLower(docName), total) {
			return true
		}
	}

	return false

}
