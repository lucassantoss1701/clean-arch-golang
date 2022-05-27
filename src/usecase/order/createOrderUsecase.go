package order

import (
	"fmt"
	itemModel "lucassantoss1701/clean/src/entity/item/model"
	orderGateway "lucassantoss1701/clean/src/entity/order/gateway"
	orderModel "lucassantoss1701/clean/src/entity/order/model"
	paymentModel "lucassantoss1701/clean/src/entity/payment/model"
	paymentService "lucassantoss1701/clean/src/entity/payment/service"
)

type CreateOrderUseCase struct {
	printOrderGateway orderGateway.PrintOrderGateway
	orderGateway      orderGateway.OrderGateway
	paymentService    *paymentService.PaymentService
}

func NewCreateOrderUseCase(
	printOrderGateway orderGateway.PrintOrderGateway,
	orderGateway orderGateway.OrderGateway,
	paymentService *paymentService.PaymentService) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		printOrderGateway: printOrderGateway,
		orderGateway:      orderGateway,
		paymentService:    paymentService,
	}
}

func (useCase *CreateOrderUseCase) Execute(input *Input) {
	fmt.Println("CreateOrderUseCase.Input", input)
	order := orderModel.NewOrder(input.Items, input.PaymentMethod)

	useCase.paymentService.Pay(order.Total(), order.PaymentMethod)

	useCase.printOrderGateway.PrintCustomerOrderGateway(order)

	useCase.printOrderGateway.PrintKitchenOrderGateway(order)

	err := useCase.orderGateway.Create(order)
	if err != nil {
		fmt.Println(err)
	}
}

type Input struct {
	Items         []itemModel.Item
	PaymentMethod paymentModel.PAYMENT_METHOD
}
