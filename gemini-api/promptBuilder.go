package geminiapi

import (
	"encoding/json"
	"os"
	"strings"
)

func promptBuilder(
	profile Profile,
	dossier Dossier,
	product string,
) (string, error) {
	template, err := os.ReadFile("prompt.txt")
	if err != nil {
		return "", err
	}
	userStr, err := ToPrettyJSON(profile)
	if err != nil {
		return "", err
	}

	dossierStr, err := ToPrettyJSON(dossier)
	if err != nil {
		return "", err
	}
	replacer := strings.NewReplacer(
		"{user}", userStr,
		"{dossier}", dossierStr,
		"{product}", product,
	)
	return replacer.Replace(string(template)), nil
}

func ToPrettyJSON(v any) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
