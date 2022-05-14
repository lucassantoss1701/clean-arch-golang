package gateway

import (
	"fmt"
	"lucassantoss1701/clean/src/entity/order/models"
)

type PrintTerminalGateway struct {
}

func (printTerminalGateway *PrintTerminalGateway) PrintCustomerOrderGateway(order *models.Order) {
	fmt.Println("NOTA DO CLIENTE: ", order)
}

func (printTerminalGateway *PrintTerminalGateway) PrintKitchenOrderGateway(order *models.Order) {
	fmt.Println("NOTA DA COZINHA: ", order)
}
