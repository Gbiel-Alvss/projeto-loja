package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := CarregarConfig()
	iniciarBanco()
	

	go IniciarServidorMCP()

	r := gin.Default()
	r.GET("/health", healthHandler)
	r.GET("/pedidos", listarPedidosHandler)
	r.POST("/pedidos", criarPedidoHandler)
	r.POST("/ia/perguntar", perguntarIAHandler)

	r.Run(":" + cfg.Port)
}
