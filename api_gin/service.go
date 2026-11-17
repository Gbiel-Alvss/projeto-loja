package main

func ProcessarPedido(p Pedido) (ReciboResposta, error) {
	var calculofrete float64
	var total float64

	for _, item := range p.Itens {
		total += item.Preco
	}
	calculofrete = total + 15.00

	SalvarPedido(p)

	recibo := ReciboResposta{
		ID:         p.ID,
		Status:     "CONFIRMADO",
		TotalItems: total,
		Frete:      15.00,
		TotalFinal: calculofrete,
	}

	return recibo, nil
}

func ObterFaturamento() FaturamentoResponse {
	var total float64
	pedidos := ListarPedidos()
	TotalPedidos := len(pedidos)

	for _, p := range pedidos {
		for _, item := range p.Itens {
			total += item.Preco
		}
		total += 15.00
	}

	return FaturamentoResponse{
		TotalPedidos:  TotalPedidos,
		TotalFaturado: total,
	}

}
