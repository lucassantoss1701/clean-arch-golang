package gateway

import models "lucassantoss1701/clean/src/entity/order/models"

type PrintOrderGateway interface {
	PrintCustomerOrderGateway(order *models.Order)
	PrintKitchenOrderGateway(order *models.Order)
}
