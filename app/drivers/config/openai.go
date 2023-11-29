package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func OpenAiClient() *openai.Client {
	godotenv.Load(".env")
	c := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	return c
}
