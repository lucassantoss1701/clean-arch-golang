package gateway

import models "lucassantoss1701/clean/src/entity/payment/models"

type PaymentGateway interface {
	Pay(value int, paymentMethod models.PAYMENT_METHOD)
}
