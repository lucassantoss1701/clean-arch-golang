package gateway

import (
	"fmt"
	"lucassantoss1701/clean/src/entity/order/models"
)

type OrderDbGateway struct {
}

func (orderDbGateway *OrderDbGateway) Create(order *models.Order) {
	fmt.Println("CRIADO NO BANCO")
}
