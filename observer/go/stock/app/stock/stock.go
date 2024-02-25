package stock

import (
	"fmt"
	"math/rand"
	"stock/domain"
	"sync"
	"time"
)

type StockServer struct {
	stockClients []domain.Client
	data         domain.MarketData
}

func (s *StockServer) Register(client domain.Client) {
	s.stockClients = append(s.stockClients, client)
}

func (s *StockServer) notifyPricing() {
	clientCount := len(s.stockClients)

	if clientCount > 0 {
		for i := 0; i < clientCount; i++ {
			client := s.stockClients[i]

			client.Update(s.data)
		}
	}
}

func (s *StockServer) runChangePricing() {
	go func() {
		for {
			time.Sleep(time.Second * 4)

			s.data = domain.MarketData{
				Name:    "Corp",
				Pricing: rand.Int(),
			}
		}
	}()

}

func (s *StockServer) RunStockMarket() {
	var wg sync.WaitGroup

	wg.Add(1)
	s.runChangePricing()

	wg.Add(1)
	go func() {
		for {
			time.Sleep(time.Second * 5)

			fmt.Println("[주식 정보 갱신]")

			s.notifyPricing()
		}
	}()

	wg.Wait()
}

func NewStockServer() domain.Stock {
	return &StockServer{}
}
