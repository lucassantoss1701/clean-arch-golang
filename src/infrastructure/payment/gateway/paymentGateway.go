package service

import (
	models "lucassantoss1701/clean/src/entity/payment/models"
)

type PaymentGateway struct {
	CashGateway       *CashGateway
	CreditCardGateway *CreditCardGateway
	DebitCardGateway  *DebitCardGateway
}

func NewPaymentStrategyGateway(
	cashGateway *CashGateway,
	creditCardGateway *CreditCardGateway,
	debitCardGateway *DebitCardGateway) *PaymentGateway {
	return &PaymentGateway{
		CashGateway:       cashGateway,
		CreditCardGateway: creditCardGateway,
		DebitCardGateway:  debitCardGateway,
	}
}

func (gateway *PaymentGateway) Pay(value int, paymentMethod models.PAYMENT_METHOD) {
	gateway.getPaymentMethod(paymentMethod).Pay(value)
}

func (gateway *PaymentGateway) getPaymentMethod(paymentMethod models.PAYMENT_METHOD) PaymentMethodGateway {
	maps := make(map[models.PAYMENT_METHOD]func() PaymentMethodGateway)

	maps[models.CASH] = func() PaymentMethodGateway {
		return gateway.CashGateway
	}
	maps[models.CREDIT_CARD] = func() PaymentMethodGateway {
		return gateway.CreditCardGateway
	}
	maps[models.DEBIT_CARD] = func() PaymentMethodGateway {
		return gateway.DebitCardGateway
	}
	return maps[paymentMethod]()
}
