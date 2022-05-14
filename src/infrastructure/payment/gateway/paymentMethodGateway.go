package service

import "fmt"

type PaymentMethodGateway interface {
	Pay(value int)
}

//---

type CashGateway struct {
}

func NewCashGateway() *CashGateway {
	return &CashGateway{}
}

func (cash *CashGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM DINHEIRO: ", value)
}

//---

type CreditCardGateway struct {
}

func NewCreditCardGateway() *CreditCardGateway {
	return &CreditCardGateway{}
}

func (creditCard *CreditCardGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE CRÉDITO: ", value)
}

//---

type DebitCardGateway struct {
}

func NewDebitCardGateway() *DebitCardGateway {
	return &DebitCardGateway{}

}

func (debitCard *DebitCardGateway) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE DÉBITO: ", value)
}
