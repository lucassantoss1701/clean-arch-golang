package gateway

import "fmt"

type CashGateway struct {
}

func NewCashGateway() *CashGateway {
	return &CashGateway{}
}

func (cash *CashGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM DINHEIRO: ", value)
}
