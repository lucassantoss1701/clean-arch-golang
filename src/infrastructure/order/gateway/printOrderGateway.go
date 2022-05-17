package gateway

import (
	"fmt"
	models "lucassantoss1701/clean/src/entity/order/models"
)

type TerminalPrintGateway struct {
}

func NewTerminalPrintGateway() *TerminalPrintGateway {
	return &TerminalPrintGateway{}
}

func (terminalPrintGateway *TerminalPrintGateway) PrintCustomerOrderGateway(order *models.Order) {
	fmt.Println()
	fmt.Println("+--CLIENTE----------------------------------+")
	fmt.Println("+ Cupom: ", order.Cupom)
	fmt.Println("+ Data: ", order.CreatedAt)
	fmt.Println("+ Senha: ", order.Code)
	fmt.Println("+ Items:")
	fmt.Println("+")
	for _, item := range order.Items {

		fmt.Println("+--- Id: ", item.Id)
		fmt.Println("+--- Nome: ", item.Name)
		fmt.Println("+--- Quantiade: ", item.Quantity)
		fmt.Println("+--- Preco: ", item.Price)
		fmt.Println("+--- Total: ", item.Total())
		fmt.Println("+")
	}
	fmt.Println("+ Total: ", order.Total())
	fmt.Println("+ Forma de pagamento: ", order.PaymentMethod)
	fmt.Println("+-------------------------------------------+")
	fmt.Println()
}

func (terminalPrintGateway *TerminalPrintGateway) PrintKitchenOrderGateway(order *models.Order) {
	fmt.Println()
	fmt.Println("+--COZINHA----------------------------------+")
	fmt.Println("+ Senha: ", order.Code)
	fmt.Println("+ Items:")
	fmt.Println("+")
	for _, item := range order.Items {

		fmt.Println("+--- Id: ", item.Id)
		fmt.Println("+--- Nome: ", item.Name)
		fmt.Println("+--- Quantiade: ", item.Quantity)
		fmt.Println("+--- Preco: ", item.Price)
		fmt.Println("+--- Total: ", item.Total())
		fmt.Println("+")
	}
	fmt.Println("+-------------------------------------------+")
	fmt.Println()
}
