package gateway

import models "lucassantoss1701/clean/src/entity/order/models"

type OrderGateway interface {
	Create(order *models.Order)
}
