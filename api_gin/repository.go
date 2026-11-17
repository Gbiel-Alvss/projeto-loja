package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func iniciarBanco() {
	var err error
	db, err = gorm.Open(sqlite.Open("pedidos.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar no banco de dados " + err.Error())
	}

	db.AutoMigrate(&Pedido{}, &Item{})
}

func SalvarPedido(p Pedido) {
	db.Create(&p)
}

func ListarPedidos() []Pedido {
	var pedidos []Pedido
	db.Preload("Itens").Find(&pedidos)
	return pedidos
}

func BuscarPedidoPorID(id int64) (Pedido, bool) {
	var pedido Pedido
	result := db.Preload("Itens").First(&pedido, id)
	if result.Error != nil {
		return Pedido{}, false
	}
	return pedido, true
}
