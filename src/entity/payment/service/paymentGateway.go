package service

import (
	paymentGateway "lucassantoss1701/clean/src/entity/payment/gateway"
	paymentModel "lucassantoss1701/clean/src/entity/payment/model"
)

type PaymentService struct {
	CashGateway       paymentGateway.PaymentGateway
	CreditCardGateway paymentGateway.PaymentGateway
	DebitCardGateway  paymentGateway.PaymentGateway
}

func NewPaymentService(
	cashGateway paymentGateway.PaymentGateway,
	creditCardGateway paymentGateway.PaymentGateway,
	debitCardGateway paymentGateway.PaymentGateway) *PaymentService {

	return &PaymentService{
		CashGateway:       cashGateway,
		CreditCardGateway: creditCardGateway,
		DebitCardGateway:  debitCardGateway,
	}
}

func (service *PaymentService) Pay(value int, paymentMethod paymentModel.PAYMENT_METHOD) {
	service.getPaymentMethod(paymentMethod).Pay(value)
}

func (service *PaymentService) getPaymentMethod(paymentMethod paymentModel.PAYMENT_METHOD) paymentGateway.PaymentGateway {
	maps := make(map[paymentModel.PAYMENT_METHOD]func() paymentGateway.PaymentGateway)

	maps[paymentModel.CASH] = func() paymentGateway.PaymentGateway {
		return service.CashGateway
	}
	maps[paymentModel.CREDIT_CARD] = func() paymentGateway.PaymentGateway {
		return service.CreditCardGateway
	}
	maps[paymentModel.DEBIT_CARD] = func() paymentGateway.PaymentGateway {
		return service.DebitCardGateway
	}
	return maps[paymentMethod]()
}
