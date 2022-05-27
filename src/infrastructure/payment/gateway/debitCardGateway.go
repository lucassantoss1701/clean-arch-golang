package gateway

import "fmt"

type DebitCardGateway struct {
}

func NewDebitCardGateway() *DebitCardGateway {
	return &DebitCardGateway{}

}

func (debitCard *DebitCardGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE DÉBITO: ", value)
}
