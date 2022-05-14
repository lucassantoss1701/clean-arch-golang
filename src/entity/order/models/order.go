package models

import (
	itemModel "lucassantoss1701/clean/src/entity/item/models"
	paymentModel "lucassantoss1701/clean/src/entity/payment/models"

	"time"
)

type Order struct {
	cupom       int
	items       []itemModel.Item
	createdAt   time.Time
	paymentType paymentModel.PAYMENT_TYPE
	code        int
}

func NewOrder(items []itemModel.Item, paymentType paymentModel.PAYMENT_TYPE) *Order {
	return &Order{
		items:       items,
		createdAt:   time.Now(),
		paymentType: paymentType,
	}
}

func (order *Order) Total() int {
	total := 0
	for _, item := range order.items {
		total = +item.Total()
	}
	return total
}

func (order *Order) GetPaymentType() *paymentModel.PAYMENT_TYPE {
	return &order.paymentType
}
