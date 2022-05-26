package gateway

type PaymentGateway interface {
	Pay(value int)
}
