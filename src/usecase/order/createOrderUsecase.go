package order

import (
	itemModel "lucassantoss1701/clean/src/entity/item/models"
	orderGateway "lucassantoss1701/clean/src/entity/order/gateway"
	orderModel "lucassantoss1701/clean/src/entity/order/models"
	paymentGateway "lucassantoss1701/clean/src/entity/payment/gateway"
	paymentModel "lucassantoss1701/clean/src/entity/payment/models"
)

type CreateOrderUseCase struct {
	printOrderGateway orderGateway.PrintOrderGateway
	orderGateway      orderGateway.OrderGateway
	paymentGateway    paymentGateway.PaymentGateway
}

func NewCreateOrderUseCase(printOrderGateway orderGateway.PrintOrderGateway,
	orderGateway orderGateway.OrderGateway,
	paymentGateway paymentGateway.PaymentGateway) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		printOrderGateway: printOrderGateway,
		orderGateway:      orderGateway,
		paymentGateway:    paymentGateway,
	}
}

func (useCase *CreateOrderUseCase) Execute(input *Input) {
	order := orderModel.NewOrder(input.Items, input.PaymentType)

	useCase.paymentGateway.Pay(order.Total(), order.PaymentMethod)

	useCase.printOrderGateway.PrintCustomerOrderGateway(order)

	useCase.printOrderGateway.PrintKitchenOrderGateway(order)

	useCase.orderGateway.Create(order)
}

type Input struct {
	Items       []itemModel.Item
	PaymentType paymentModel.PAYMENT_METHOD
}
