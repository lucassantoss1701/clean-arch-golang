package config

import (
	"database/sql"
	paymentService "lucassantoss1701/clean/src/entity/payment/service"
	orderController "lucassantoss1701/clean/src/infrastructure/order/controller"
	orderGateway "lucassantoss1701/clean/src/infrastructure/order/gateway"
	paymentGateway "lucassantoss1701/clean/src/infrastructure/payment/gateway"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"
)

type ApplicationContext struct {
	CreateOrderController *orderController.CreateOrderController
}

func NewApplicationContext() *ApplicationContext {
	terminalPrintGateway := orderGateway.NewTerminalPrintGateway()
	orderDbGateway := orderGateway.NewOrderDbGateway(getConnection())
	cashGateway := paymentGateway.NewCashGateway()
	creditCardGateway := paymentGateway.NewCreditCardGateway()
	debitCardGateway := paymentGateway.NewDebitCardGateway()
	paymentService := paymentService.NewPaymentService(cashGateway, creditCardGateway, debitCardGateway)

	createOrderUseCase := orderUseCase.NewCreateOrderUseCase(terminalPrintGateway, orderDbGateway, paymentService)

	createOrderController := orderController.NewCreateOrderController(createOrderUseCase)

	return &ApplicationContext{
		CreateOrderController: createOrderController,
	}
}

func getConnection() *sql.DB {
	dbConfig, err := GetConnection()
	if err != nil {
		panic(err)
	}
	return dbConfig.conn
}
