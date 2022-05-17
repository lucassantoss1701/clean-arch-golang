package gateway

import (
	"fmt"
	"lucassantoss1701/clean/src/entity/order/models"
)

type OrderDbGateway struct {
}

func NewOrderDbGateway() *OrderDbGateway {
	return &OrderDbGateway{}
}

func (orderDbGateway *OrderDbGateway) Create(order *models.Order) {
	fmt.Println("Saving in Database...")
}
