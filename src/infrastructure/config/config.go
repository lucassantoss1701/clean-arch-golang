package config

import (
	"database/sql"
	"lucassantoss1701/clean/src/infrastructure/order/controller"
	orderGateway "lucassantoss1701/clean/src/infrastructure/order/gateway"
	paymentGateway "lucassantoss1701/clean/src/infrastructure/payment/gateway"
	orderUseCase "lucassantoss1701/clean/src/usecase/order"
)

type ApplicationContext struct {
	CreateOrderController *controller.CreateOrderController
}

func NewApplicationContext() *ApplicationContext {

	terminalPrintGateway := orderGateway.NewTerminalPrintGateway()
	orderDbGateway := orderGateway.NewOrderDbGateway(getConnection())
	cashGateway := paymentGateway.NewCashGateway()
	creditCardGateway := paymentGateway.NewCreditCardGateway()
	debitCardGateway := paymentGateway.NewDebitCardGateway()
	paymentGateway := paymentGateway.NewPaymentStrategyGateway(cashGateway, creditCardGateway, debitCardGateway)

	createOrderUseCase := orderUseCase.NewCreateOrderUseCase(terminalPrintGateway, orderDbGateway, paymentGateway)

	createOrderController := controller.NewCreateOrderController(createOrderUseCase)

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
