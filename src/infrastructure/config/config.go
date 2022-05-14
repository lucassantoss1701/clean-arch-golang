package config

import (
	"lucassantoss1701/clean/src/infrastructure/order/controller"
	orderGateway "lucassantoss1701/clean/src/infrastructure/order/gateway"
	paymentGateway "lucassantoss1701/clean/src/infrastructure/payment/gateway"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"
)

type ApplicationContext struct {
	CreateOrderController *controller.CreateOrderController
}

func NewApplicationContext() *ApplicationContext {
	printTerminalGateway := &orderGateway.PrintTerminalGateway{}
	orderDbGateway := &orderGateway.OrderDbGateway{}
	paymentGateway := &paymentGateway.PaymentGateway{}

	createOrderUseCase := orderUseCase.NewCreateOrderUseCase(printTerminalGateway, orderDbGateway, paymentGateway)

	createOrderController := controller.NewCreateOrderController(createOrderUseCase)

	return &ApplicationContext{
		CreateOrderController: createOrderController,
	}
}
