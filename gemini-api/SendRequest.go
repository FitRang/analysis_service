package geminiapi

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func SendRequest(
	ctx context.Context,
	user string,
	dossier string,
	product string,
) (string, error) {
	godotenv.Load(".env")
	APIKEY := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: APIKEY,
	})
	if err != nil {
		log.Fatal(err)
	}

	prompt, err := promptBuilder(user, dossier, product)
	if err != nil {
		return "", err
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text(), nil
}
