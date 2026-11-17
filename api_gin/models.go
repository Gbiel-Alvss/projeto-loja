package main

type Item struct {
	ID       uint    `json:"-" gorm:"primaryKey"`
	PedidoID int64   `json:"-"`
	Nome     string  `json:"nome" binding:"required"`
	Preco    float64 `json:"preco" binding:"required"`
}

type Pedido struct {
	ID      int64  `json:"pedido_id" gorm:"primaryKey"`
	Cliente string `json:"cliente" binding:"required"`
	Itens   []Item `json:"items" binding:"required,dive" gorm:"foreignKey:PedidoID"`
}

type ReciboResposta struct {
	ID         int64   `json:"pedido_id"`
	Status     string  `json:"status"`
	TotalItems float64 `json:"total_items"`
	Frete      float64 `json:"frete"`
	TotalFinal float64 `json:"total_final"`
}

type FaturamentoResponse struct {
	TotalPedidos  int     `json:"total_pedidos"`
	TotalFaturado float64 `json:"total_faturado"`
}

type Message struct {
	Role        string     `json:"role"`
	Content     string     `json:"content"`
	ToolCalls   []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID  string     `json:"tool_call_id,omitempty"`
	ExtraConent any        `json:"extra_content,omitempty"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Tools    []Tool    `json:"tools,omitempty"`
}

type Choice struct {
	Message Message `json:"message"`
}

type ChatResponde struct {
	Choices []Choice `json:"choices"`
}

type PerguntaIARequest struct {
	Pergunta string `json:"pergunta" binding:"required"`
}

type FunctionDefinition struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

type Tool struct {
	Type     string             `json:"type"`
	Function FunctionDefinition `json:"function"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCall struct {
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	Function     FunctionCall `json:"function"`
	ExtraContent any          `json:"extra_content,omitempty"`
}
