package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func IniciarServidorMCP() {
	s := server.NewMCPServer("Loja MCP", "1.0.0")

	ferramentaFaturamento := mcp.NewTool(
		"obter_faturamento",
		mcp.WithDescription("Consulta o banco SQLite e retorna o faturamento total em reais e a quantidade de pedidos da loja."),
	)

	s.AddTool(ferramentaFaturamento, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

		relatorio := ObterFaturamento()

		bytesJSON, _ := json.Marshal(relatorio)

		return mcp.NewToolResultText(string(bytesJSON)), nil
	})

	sse := server.NewSSEServer(s)

	fmt.Println("Servidor MCP (SSE/HTTP) rodando na porta :8081...")
	if err := sse.Start(":8081"); err != nil {
		fmt.Printf("Erro no servidor MCP SSE: %v\n", err)
	}

}
