package geminiapi

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed prompt.txt
var promptTemplate string

func promptBuilder(
	profile string,
	dossier string,
	product string,
) (string, error) {
	template, err := os.ReadFile(promptTemplate)
	if err != nil {
		return "", err
	}

	replacer := strings.NewReplacer(
		"{user}", profile,
		"{dossier}", dossier,
		"{product}", product,
	)
	return replacer.Replace(string(template)), nil
}
