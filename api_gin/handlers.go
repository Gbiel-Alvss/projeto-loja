package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

func criarPedidoHandler(c *gin.Context) {
	var transacao Pedido

	if err := c.ShouldBindJSON(&transacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recibo, err := ProcessarPedido(transacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, recibo)
}

func listarPedidosHandler(c *gin.Context) {
	idParam := c.Query("id")

	if idParam == "" {
		c.JSON(http.StatusOK, ListarPedidos())
		return
	}

	idNum, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, achou := BuscarPedidoPorID(idNum)
	if !achou {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}
	c.JSON(http.StatusOK, pedido)

}

func perguntarIAHandler(c *gin.Context) {
	var req PerguntaIARequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resposta, err := ExecutarAgente(req.Pergunta)
	if err != nil {
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao consultar IA"})
		return
	}

	fmt.Println(">>> RESPOSTA RECEBIDA DA IA:", resposta)

	c.JSON(http.StatusOK, resposta)
}
