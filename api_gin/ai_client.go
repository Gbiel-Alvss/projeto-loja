package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func EnviarChat(req ChatRequest) (ChatResponde, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return ChatResponde{}, errors.New("Chame API não configurada")
	}

	dadosJSON, err := json.Marshal(req)
	if err != nil {
		return ChatResponde{}, err
	}

	httpReq, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions", bytes.NewBuffer(dadosJSON))
	if err != nil {
		return ChatResponde{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}

	resp, err := client.Do(httpReq)
	if err != nil {
		return ChatResponde{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		corpoErro, _ := io.ReadAll(resp.Body)
		return ChatResponde{}, fmt.Errorf("status %d da Google: %s", resp.StatusCode, string(corpoErro))
	}
	corpoBytes, err := io.ReadAll(resp.Body)

	var chatResp ChatResponde

	if err := json.Unmarshal(corpoBytes, &chatResp); err != nil {
		return ChatResponde{}, err
	}

	if len(chatResp.Choices) == 0 {
		return ChatResponde{}, errors.New("Resposta da IA veio vazia")
	}
	return chatResp, nil

}
