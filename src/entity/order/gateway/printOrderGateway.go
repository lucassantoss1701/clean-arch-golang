package gateway

import orderModel "lucassantoss1701/clean/src/entity/order/model"

type PrintOrderGateway interface {
	PrintCustomerOrderGateway(order *orderModel.Order)
	PrintKitchenOrderGateway(order *orderModel.Order)
}
