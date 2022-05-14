package gateway

import "lucassantoss1701/clean/src/entity/payment/models"

type PaymentGateway interface {
	Pay(value int, paymentType *models.PAYMENT_TYPE)
}
