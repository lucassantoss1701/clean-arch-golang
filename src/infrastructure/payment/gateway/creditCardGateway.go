package gateway

import "fmt"

type CreditCardGateway struct {
}

func NewCreditCardGateway() *CreditCardGateway {
	return &CreditCardGateway{}
}

func (creditCard *CreditCardGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE CRÉDITO: ", value)
}
