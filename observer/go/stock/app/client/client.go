package client

import (
	"fmt"
	"stock/domain"
	"strconv"
)

type StockClient struct {
	name string
}

func (s StockClient) Update(data domain.MarketData) {
	fmt.Print(s.name + " 주식 정보 갱신")
	fmt.Print("주식명 : " + data.Name)
	fmt.Print("주식 가격 : " + strconv.Itoa(data.Pricing) + "원")
	fmt.Println(" ")
}

func NewStockClient(name string) domain.Client {
	return &StockClient{
		name: name,
	}
}
