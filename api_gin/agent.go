package main

import "encoding/json"

func ExecutarAgente(prompt string) (string, error) {

	ferramentas := []Tool{
		{
			Type: "function",
			Function: FunctionDefinition{
				Name:        "obter_faturamento",
				Description: "Consulta o banco de dados e retorna o total faturado em reais e a quantidade de pedidos da loja.",
				Parameters: map[string]any{
					"type":       "object",
					"properties": map[string]any{},
				},
			},
		},
	}
	mensagens := []Message{
		{Role: "user", Content: prompt},
	}

	req := ChatRequest{
		Model:    "gemini-3.5-flash",
		Messages: mensagens,
		Tools:    ferramentas,
	}

	resp, err := EnviarChat(req)
	if err != nil {
		return "", err
	}

	msgIA := resp.Choices[0].Message
	if len(msgIA.ToolCalls) == 0 {
		return msgIA.Content, nil
	}

	chamada := msgIA.ToolCalls[0]

	var relatorioJSON []byte
	if chamada.Function.Name == "obter_faturamento" {
		relatorio := ObterFaturamento()
		relatorioJSON, _ = json.Marshal(relatorio)
	}

	req.Messages = append(req.Messages, msgIA)
	req.Messages = append(req.Messages, Message{
		Role:       "tool",
		ToolCallID: chamada.ID,
		Content:    string(relatorioJSON),
	})

	respFinal, err := EnviarChat(req)
	if err != nil {
		return "", err
	}

	return respFinal.Choices[0].Message.Content, nil
}
