package main

import (
	"stock/app/client"
	"stock/app/stock"
	"stock/domain"
)

func main() {
	clients := []domain.Client{
		client.NewStockClient("Client1"),
		client.NewStockClient("Client2"),
		client.NewStockClient("Client3"),
		client.NewStockClient("Client4"),
		client.NewStockClient("Client5"),
	}

	server := stock.NewStockServer()

	for _, item := range clients {
		server.Register(item)
	}

	server.RunStockMarket()
}
