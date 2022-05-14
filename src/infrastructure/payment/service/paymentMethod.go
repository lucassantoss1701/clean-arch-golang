package service

import "fmt"

type PaymentMethod interface {
	Pay(value int)
}

type Cash struct {
}

func NewCash() *Cash {
	return &Cash{}
}

type CreditCard struct {
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

type DebitCard struct {
}

func NewDebitCard() *DebitCard {
	return &DebitCard{}

}

func (cash *Cash) Pay(value int) {
	fmt.Println("PAGAMENTO COM DINHEIRO: ", value)
}

func (creditCard *CreditCard) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE CRÉDITO: ", value)
}

func (debitCard *DebitCard) Pay(value int) {
	fmt.Println("PAGAMENTO COM CARTÃO DE DÉBITO: ", value)
}
