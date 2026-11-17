package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	GeminiAPIKey string
}

func CarregarConfig() Config {
	_ = godotenv.Load()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("ERRO CRÍTICO: Variável GEMINI_API_KEY é obrigatória")
	}

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	return Config{
		Port:         porta,
		GeminiAPIKey: apiKey,
	}
}
