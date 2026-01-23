package geminiapi

import (
	_ "embed"
	"strings"
)

//go:embed prompt.txt
var promptTemplate string

func promptBuilder(
	profile string,
	dossier string,
	product string,
) (string, error) {
	replacer := strings.NewReplacer(
		"{user}", profile,
		"{dossier}", dossier,
		"{product}", product,
	)

	return replacer.Replace(promptTemplate), nil
}
