package service

import (
	"lucassantoss1701/clean/src/entity/payment/models"
	service "lucassantoss1701/clean/src/infrastructure/payment/service"
)

type PaymentGateway struct {
}

func (paymentGateway *PaymentGateway) Pay(value int, paymentType *models.PAYMENT_TYPE) {
	getPaymentMethod(paymentType).Pay(value)
}

func getPaymentMethod(paymentType *models.PAYMENT_TYPE) service.PaymentMethod {
	maps := make(map[models.PAYMENT_TYPE]func() service.PaymentMethod)

	maps[models.CASH] = func() service.PaymentMethod {
		return service.NewCash()
	}
	maps[models.CREDIT_CARD] = func() service.PaymentMethod {
		return service.NewCreditCard()
	}
	maps[models.DEBIT_CARD] = func() service.PaymentMethod {
		return service.NewDebitCard()
	}
	return maps[*paymentType]()
}
