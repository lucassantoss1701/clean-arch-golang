package gateway

import orderModel "lucassantoss1701/clean/src/entity/order/model"

type OrderGateway interface {
	Create(order *orderModel.Order) error
}
